package repository

import (
	"database/sql"
	"project-sanber/structs"
)

func GetAllBuku(db *sql.DB) (results []structs.Buku, err error) {
	sql := "SELECT * FROM buku"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var buku = structs.Buku{}

		err = rows.Scan(
			&buku.ID, 
			&buku.Judul_buku, 
			&buku.Author, 
			&buku.Tahun_terbit, 
			&buku.Publisher, 
			&buku.ISBN, 
			&buku.Jumlah_halaman, 
			&buku.Is_dipinjam, 
			&buku.Created_at, 
			&buku.Created_by, 
			&buku.Updated_at, 
			&buku.Updated_by,
		)
		if err != nil {
			return
		}

		results = append(results, buku)
	}

	return
}

func InsertBuku(db *sql.DB, buku structs.Buku) (err error) {

	sql := "INSERT INTO buku (judul_buku, author, tahun_terbit, publisher, isbn, jumlah_halaman, is_dipinjam, created_at, created_by, updated_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	errs := db.QueryRow(
		sql, 

		buku.Judul_buku, 
		buku.Author, 
		buku.Tahun_terbit, 
		buku.Publisher, 
		buku.ISBN, 
		buku.Jumlah_halaman, 
		buku.Is_dipinjam, 
		buku.Created_at, 
		buku.Created_by, 
		buku.Updated_by,
		buku.Updated_at, 
	)

	return errs.Err()
}

func UpdateBuku(db *sql.DB, buku structs.Buku) (err error) {

	sql := "UPDATE buku SET judul_buku = $1, author = $2, tahun_terbit = $3, publisher = $4, isbn = $5, jumlah_halaman = $6, is_dipinjam = $7, updated_at = $8, updated_by = $9 WHERE id = $10"

	errs := db.QueryRow(
		sql, 

		buku.Judul_buku, 
		buku.Author, 
		buku.Tahun_terbit, 
		buku.Publisher, 
		buku.ISBN, 
		buku.Jumlah_halaman, 
		buku.Is_dipinjam, 
		buku.Updated_at, 
		buku.Updated_by, 
		buku.ID,
	)

	return errs.Err()
}

func DeleteBuku(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM buku WHERE id = $1"

	errs := db.QueryRow(sql, id)

	return errs.Err()
}
