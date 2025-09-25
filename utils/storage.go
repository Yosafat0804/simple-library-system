package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"perpustakaan-mini/models"
	"perpustakaan-mini/services"
)

type Data struct {
	Books        []models.Book        `json:"books"`
	Transactions []models.Transaction `json:"transactions"`
}

// Simpan data ke file JSON
func SaveData(filename string, library services.Library) {
	data := Data{
		Books:        library.Books,
		Transactions: library.Transactions,
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
	}
}

// Baca data dari file JSON
func LoadData(filename string) services.Library {
	var library services.Library

	file, err := os.Open(filename)
	if err != nil {
		// Jika file tidak ada, kembalikan library kosong
		return library
	}
	defer file.Close()

	var data Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		return library
	}

	library.Books = data.Books
	library.Transactions = data.Transactions
	return library
}
