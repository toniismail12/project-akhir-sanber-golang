package structs

type Buku struct {
	ID             int `json:"id"`
	Judul_buku     string
	Author         string
	Tahun_terbit   int
	Publisher      string
	ISBN           string
	Jumlah_halaman string
	Is_dipinjam    string
	Created_at     string
	Created_by     string
	Updated_at     string
	Updated_by     string
}
