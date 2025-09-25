package models

type Transaction struct {
	BookID    int
	BookTitle string
	Borrower  string
	Action    string // "Pinjam" atau "Kembali"
}