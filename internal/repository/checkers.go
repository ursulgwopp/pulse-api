package repository

import "context"

func (r *PostgresRepository) CheckLoginExists(login string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)`
	if err := r.db.QueryRowContext(ctx, query, login).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckCountryCodeExists(alpha2 string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM countries WHERE alpha2 = $1)`
	if err := r.db.QueryRowContext(ctx, query, alpha2).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckPhoneExists(phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE phone = $1)`
	if err := r.db.QueryRowContext(ctx, query, phone).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

// func (r *PostgresRepository) CheckUserIdExists(id int) (bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
// 	defer cancel()

// 	var exists bool
// 	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
// 	if err := r.db.QueryRowContext(ctx, query, id).Scan(&exists); err != nil {
// 		return false, err
// 	}

// 	return exists, nil
// }
