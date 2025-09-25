package services

import (
	"fmt"
	"perpustakaan-mini/models"
	"strings"
)

type Library struct {
	Books        []models.Book
	Transactions []models.Transaction
}

func (l *Library) AddBook(book models.Book) {
	l.Books = append(l.Books, book)
	fmt.Println("Buku berhasil ditambahkan!")
}

func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("Tidak ada buku di perpustakaan.")
		return
	}

	fmt.Println("\nDaftar Buku:")
	for _, b := range l.Books {
		status := "Tersedia"
		if b.IsBorrowed {
			status = "Dipinjam"
		}
		fmt.Printf("[%d] %s - %s (%s)\n", b.ID, b.Title, b.Author, status)
	}
}

func (l *Library) BorrowBook(id int, borrower string) {
	for i, b := range l.Books {
		if b.ID == id {
			if b.IsBorrowed {
				fmt.Println("Buku sudah dipinjam.")
				return
			}
			l.Books[i].IsBorrowed = true
			fmt.Println("Buku berhasil dipinjam!")
			l.Transactions = append(l.Transactions, models.Transaction{
				BookID:    b.ID,
				BookTitle: b.Title,
				Borrower:  borrower,
				Action:    "Pinjam",
			})
			return
		}
	}
	fmt.Println("Buku tidak ditemukan.")
}

func (l *Library) ReturnBook(id int, borrower string) {
	for i, b := range l.Books {
		if b.ID == id {
			if !b.IsBorrowed {
				fmt.Println("Buku ini belum dipinjam.")
				return
			}
			l.Books[i].IsBorrowed = false
			fmt.Println("Buku berhasil dikembalikan!")
			l.Transactions = append(l.Transactions, models.Transaction{
				BookID:    b.ID,
				BookTitle: b.Title,
				Borrower:  borrower,
				Action:    "Kembali",
			})
			return
		}
	}
	fmt.Println("Buku tidak ditemukan.")
}

func (l *Library) SearchBook(keyword string) {
	fmt.Println("\nHasil Pencarian:")
	found := false
	for _, b := range l.Books {
		if strings.Contains(strings.ToLower(b.Title), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(b.Author), strings.ToLower(keyword)) {
			status := "Tersedia"
			if b.IsBorrowed {
				status = "Dipinjam"
			}
			fmt.Printf("[%d] %s - %s (%s)\n", b.ID, b.Title, b.Author, status)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku yang sesuai.")
	}
}

func (l *Library) DeleteBook(id int) {
	for i, b := range l.Books {
		if b.ID == id {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Printf("Buku '%s' berhasil dihapus!\n", b.Title)
			return
		}
	}
	fmt.Println("Buku tidak ditemukan.")
}

func (l *Library) ShowTransactions() {
	if len(l.Transactions) == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}
	fmt.Println("\nRiwayat Transaksi:")
	for _, t := range l.Transactions {
		fmt.Printf("[%s] %s oleh %s (Buku ID: %d)\n",
			t.Action, t.BookTitle, t.Borrower, t.BookID)
	}
}
