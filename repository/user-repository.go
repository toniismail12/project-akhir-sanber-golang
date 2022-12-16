package repository

import (
	"database/sql"
	"project-sanber/structs"
)

func Get_User(db *sql.DB) (results []structs.Users, err error) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var users = structs.Users{}

		err = rows.Scan(
			&users.ID,  
			&users.Username, 
			&users.Password, 
			&users.Email,  
			&users.Wa, 
			&users.Nama,
			&users.Alamat,
			&users.Role,
			&users.Created_by,
			&users.Created_at,
			&users.Updated_by,
			&users.Updated_at,
		)
		if err != nil {
			return
		}

		results = append(results, users)
	}

	return
}

func Insert_user(db *sql.DB, users structs.Users) (err error) {

	sql := "INSERT INTO users (username, password, email, wa, nama, alamat, role, created_by, created_at, updated_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	errs := db.QueryRow(
		sql, 

		users.Username,
		users.Password, 
		users.Email,
		users.Wa,
		users.Nama,
		users.Alamat,
		users.Role,
		users.Created_by, 
		users.Created_at, 
		users.Updated_by,
		users.Updated_at, 
	)

	return errs.Err()
}

func Delete_user(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM users WHERE id = $1"

	errs := db.QueryRow(sql, id)

	return errs.Err()
}

func Detail_user(db *sql.DB, username string) (results []structs.Users, err error) {
	sql := "SELECT * FROM users WHERE username = $1"

	rows, err := db.Query(sql, username)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var users = structs.Users{}

		err = rows.Scan(
			&users.ID,  
			&users.Username, 
			&users.Password, 
			&users.Email,  
			&users.Wa, 
			&users.Nama,
			&users.Alamat,
			&users.Role,
			&users.Created_by,
			&users.Created_at,
			&users.Updated_by,
			&users.Updated_at,
		)
		if err != nil {
			return
		}

		results = append(results, users)
	}

	return
}

func Insert_user_login(db *sql.DB, users structs.User_login) (err error) {

	sql := "INSERT INTO user_login (username, nama, role, token) VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(
		sql, 

		users.Username,
		users.Nama,
		users.Role,
		users.Token,
	)

	return errs.Err()
}

func Info_user_login(db *sql.DB, token string) (results []structs.User_login, err error) {
	sql := "SELECT * FROM user_login WHERE token = $1"

	rows, err := db.Query(sql, token)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var users = structs.User_login{}

		err = rows.Scan(
			&users.ID,
			&users.Username, 
			&users.Nama,
			&users.Role,
			&users.Token,
		)
		if err != nil {
			return
		}

		results = append(results, users)
	}

	return
}
