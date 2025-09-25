# 📚 Sistem Perpustakaan Sederhana

Aplikasi console berbasis Go untuk mengelola data buku di perpustakaan.  
Program ini menyediakan fitur CRUD (Create, Read, Update, Delete), pencarian buku, serta pencatatan riwayat peminjaman dan pengembalian.  
Data tersimpan dalam format JSON sehingga tetap aman meskipun aplikasi ditutup.

---

## 🚀 Fitur

- **Tambah Buku** : Menambahkan data buku baru (judul & penulis)  
- **Lihat Buku** : Menampilkan daftar semua buku beserta status peminjaman  
- **Pinjam Buku** : Mencatat buku yang dipinjam oleh pengguna tertentu  
- **Kembalikan Buku** : Mengubah status buku yang sudah dipinjam  
- **Cari Buku** : Mencari buku berdasarkan judul atau penulis  
- **Hapus Buku** : Menghapus data buku dari sistem  
- **Riwayat Transaksi** : Melihat daftar peminjaman & pengembalian  
- **Penyimpanan JSON** : Data otomatis tersimpan ke file `library.json`

---

## 📂 Struktur Direktori

perpustakaan-mini/
├── main.go              # Entry point aplikasi
├── go.mod               # Go module file
├── models/
│   └── book.go          # Definisi struct Book & Library
├── services/
│   └── transaction.go   # Logika transaksi peminjaman/pengembalian
├── utils/
│   └── io.go            # Fungsi input & penyimpanan file JSON
└── library.json         # Database sederhana untuk menyimpan data buku

---

## 🗂️ Model Data Buku

```go
type Book struct {
    ID         int    `json:"id"`
    Title      string `json:"title"`
    Author     string `json:"author"`
    IsBorrowed bool   `json:"is_borrowed"`
}
Model di atas digunakan untuk menyimpan informasi dasar sebuah buku, termasuk ID, judul, penulis, dan status peminjaman.

🛠️ Cara Menjalankan
Pastikan Go versi 1.20 atau lebih baru sudah terpasang

Clone repositori:

bash
Copy code
git clone https://github.com/username/perpustakaan-mini.git
Masuk ke folder proyek:

bash
Copy code
cd perpustakaan-mini
Jalankan aplikasi:

bash
Copy code
go run main.go
🖥️ Tampilan Menu Utama
markdown
Copy code
================================
SISTEM PERPUSTAKAAN SEDERHANA
================================
1. Tambah Buku
2. Lihat Buku
3. Pinjam Buku
4. Kembalikan Buku
5. Cari Buku
6. Hapus Buku
7. Riwayat Transaksi
8. Keluar
📖 Panduan Singkat
Tambah Buku → masukkan judul & penulis, data otomatis tersimpan

Lihat Buku → tampilkan semua daftar buku dan status pinjam

Pinjam Buku → masukkan ID buku dan nama peminjam

Kembalikan Buku → masukkan ID buku yang dikembalikan

Cari Buku → gunakan kata kunci (judul/penulis) untuk pencarian

Hapus Buku → hapus data buku dari sistem berdasarkan ID

Riwayat Transaksi → tampilkan catatan peminjaman & pengembalian

💾 Penyimpanan Data
Semua data tersimpan di file library.json

File dibuat otomatis saat aplikasi pertama kali dijalankan

Saat keluar (menu nomor 8), data terbaru akan disimpan kembali

👨‍💻 Author
Nama: Yosafat Bagas Herlianka

Mata Kuliah: Pemrograman Jaringan

Keterangan: Aplikasi ini dibuat sebagai latihan implementasi Go language dengan konsep struct, fungsi, dan modularisasi.

yaml
Copy code
