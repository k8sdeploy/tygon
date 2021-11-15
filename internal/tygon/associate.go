package tygon

import (
	"errors"
	"fmt"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/jackc/pgx/v4"
)

func (t *Tygon) getConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(
		t.Context,
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s",
			t.Config.Username,
			t.Config.Password,
			t.Config.RDSAddress,
			t.Config.Database))
	if err != nil {
		return nil, bugLog.Error(err)
	}

	return conn, nil
}

func (t *Tygon) GetSecretAndAccount(hostname string) error {
	a := Account{}

	conn, err := t.getConnection()
	if err != nil {
		return bugLog.Error(err)
	}
	defer func() {
		if err := conn.Close(t.Context); err != nil {
			bugLog.Debugf("failed to close associate: %+v", err)
		}
	}()

	if err := conn.QueryRow(
		t.Context,
		`SELECT name, secret, account_id FROM account WHERE hostname = $1`,
		hostname,
	).Scan(&a.Name, &a.Secret, &a.ID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return bugLog.Error(err)
		}
	}
	t.Account = &a

	return nil
}
