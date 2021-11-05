package tygon

import (
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

func (t *Tygon) associateAccountWithOrganization(orgName string) error {
	conn, err := t.getConnection()
	if err != nil {
		return bugLog.Error(err)
	}
	defer func() {
		if err := conn.Close(t.Context); err != nil {
			bugLog.Debugf("failed to close associate: %+v", err)
		}
	}()

	_, err = conn.Exec(
		t.Context,
		`SELECT account_id FROM account WHERE organization_name = $1`,
		orgName,
		t.Account.ID)
	if err != nil {
		return bugLog.Error(err)
	}

	return nil
}
