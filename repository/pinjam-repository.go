package repository

import (
	"database/sql"
	"project-sanber/structs"
)

func Insert_pinjam(db *sql.DB, borrow structs.Pinjam_buku) (err error) {

	sql := "INSERT INTO pinjam_buku (id_buku, id_user_peminjam, created_by, updated_by, updated_at) VALUES ($1, $2, $3, $4, $5)"

	errs := db.QueryRow(
		sql,

		borrow.Id_buku,
		borrow.Id_user_peminjam,
		borrow.Created_by,
		borrow.Updated_by,
		borrow.Updated_at,
	)

	return errs.Err()
}

func Get_peminjam(db *sql.DB) (results []structs.Peminjam, err error) {
	sql := "SELECT pinjam_buku.id_buku, pinjam_buku.id_user_peminjam, users.nama, buku.judul_buku FROM pinjam_buku JOIN buku ON pinjam_buku.id_buku = buku.id JOIN users ON users.id = pinjam_buku.id_user_peminjam"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var users = structs.Peminjam{}

		err = rows.Scan(
			&users.Id_buku,
			&users.Id_user_peminjam,  
			&users.Nama, 
			&users.Judul_buku, 
		)
		if err != nil {
			return
		}

		results = append(results, users)
	}

	return
}
