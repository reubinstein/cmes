package initializers

import (
	"database/sql"
	"errors"
)

// Policy struct represents a policy in the system
type Policy struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

// GetAllPolicies function returns all policies from the database
func GetAllPolicies() ([]Policy, error) {
	rows, err := db.Query("SELECT id, name, category FROM policies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	policies := []Policy{}
	for rows.Next() {
		var p Policy
		err := rows.Scan(&p.ID, &p.Name, &p.Category)
		if err != nil {
			return nil, err
		}
		policies = append(policies, p)
	}

	if len(policies) == 0 {
		return nil, errors.New("No policies found")
	}

	return policies, nil
}

// GetPolicyById function returns a policy from the database given an id
func GetPolicyById(id int) (Policy, error) {
	var p Policy
	err := db.QueryRow("SELECT id, name, category FROM policies WHERE id = ?", id).Scan(&p.ID, &p.Name, &p.Category)
	if err == sql.ErrNoRows {
		return Policy{}, errors.New("Policy not found")
	} else if err != nil {
		return Policy{}, err
	}
	return p, nil
}

// CreatePolicy function creates a new policy in the database
func CreatePolicy(policy Policy) (int, error) {
	result, err := db.Exec("INSERT INTO policies(name, category) VALUES(?, ?)", policy.Name, policy.Category)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// UpdatePolicy function updates an existing policy in the database
func UpdatePolicy(policy Policy) error {
	result, err := db.Exec("UPDATE policies SET name=?, category=? WHERE id=?", policy.Name, policy.Category, policy.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("No policy found with given id")
	}
	return nil
}

// DeletePolicy function deletes a policy from the database given an id
func DeletePolicy(id int) error {
	result, err := db.Exec("DELETE FROM policies WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("No policy found with given id")
	}
	return nil
}
