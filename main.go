package main

import (
	"fmt"
	"perpustakaan-mini/models"
	"perpustakaan-mini/utils"
)

func main() {
	const dataFile = "library.json"

	// Load data dari file
	library := utils.LoadData(dataFile)
	idCounter := len(library.Books) + 1

	for {
		fmt.Println("\n================================")
		fmt.Println("SISTEM PERPUSTAKAAN SEDERHANA")
		fmt.Println("================================")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Lihat Buku")
		fmt.Println("3. Pinjam Buku")
		fmt.Println("4. Kembalikan Buku")
		fmt.Println("5. Cari Buku")
		fmt.Println("6. Hapus Buku")
		fmt.Println("7. Riwayat Transaksi")
		fmt.Println("8. Keluar")

		pilihan := utils.InputInt("Pilih menu: ")

		switch pilihan {
		case 1:
			title := utils.InputString("Judul Buku: ")
			author := utils.InputString("Penulis: ")
			book := models.Book{
				ID:         idCounter,
				Title:      title,
				Author:     author,
				IsBorrowed: false,
			}
			library.AddBook(book)
			idCounter++

		case 2:
			library.ListBooks()

		case 3:
			id := utils.InputInt("Masukkan ID Buku yang ingin dipinjam: ")
			nama := utils.InputString("Nama Peminjam: ")
			library.BorrowBook(id, nama)

		case 4:
			id := utils.InputInt("Masukkan ID Buku yang ingin dikembalikan: ")
			nama := utils.InputString("Nama Peminjam: ")
			library.ReturnBook(id, nama)

		case 5:
			keyword := utils.InputString("Masukkan kata kunci judul/penulis: ")
			library.SearchBook(keyword)

		case 6:
			id := utils.InputInt("Masukkan ID Buku yang ingin dihapus: ")
			library.DeleteBook(id)

		case 7:
			library.ShowTransactions()

		case 8:
			// Simpan data ke file sebelum keluar
			utils.SaveData(dataFile, library)
			fmt.Println("Data berhasil disimpan ke", dataFile)
			fmt.Println("Terima kasih telah menggunakan sistem perpustakaan!")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
