package postgres

import (
	"fund-management-information-system/internal_types"
	"github.com/jmoiron/sqlx"
)

type ManagerRepositoryPostgres struct {
	db *sqlx.DB
}

func NewManagerPostgres(db *sqlx.DB) *ManagerRepositoryPostgres {
	return &ManagerRepositoryPostgres{db: db}
}

// устарело
/*func (r *ManagerRepositoryPostgres) Delete(managerId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Id=$1", managersTable)
	_, err := r.db.Exec(query, managerId)
	return err
}*/
func (r *ManagerRepositoryPostgres) DeleteById(managerId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	deleteWorkQuery := `
	DELETE FROM manager_work_info
	WHERE id=$1
`
	deletePersonQuery := `
	DELETE FROM persons
	WHERE id=$1
`
	deleteAccountQuery := `
	DELETE FROM accounts
	WHERE id=$1
`
	deleteManagerQuery := `
	DELETE FROM managers
	WHERE id=$1
	
`
	_, err = r.db.Exec(deleteManagerQuery, managerId)
	_, err = r.db.Exec(deletePersonQuery, managerId)
	_, err = r.db.Exec(deleteAccountQuery, managerId)
	_, err = r.db.Exec(deleteWorkQuery, managerId)

	return tx.Commit()
}

func (r *ManagerRepositoryPostgres) GetById(managerId int) (internal_types.Manager, error) {
	var manager internal_types.Manager

	query := `
	SELECT p.email, p.phone, p.address,
       mwi.bank_account, mwi.capital_managment, mwi.profit_percent_day
	FROM managers mng
	LEFT JOIN persons p ON mng.id = p.id
    LEFT JOIN manager_work_info mwi  ON mng.id = mwi.id
	WHERE mng.id=$1`

	if err := r.db.Get(&manager, query, managerId); err != nil {
		return manager, err
	}
	return manager, nil
}
