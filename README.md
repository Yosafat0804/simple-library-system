📚 Sistem Perpustakaan Sederhana

Aplikasi berbasis console yang dibuat menggunakan bahasa pemrograman Go.
Program ini memungkinkan pengguna untuk mengelola data buku dengan fitur CRUD (Create, Read, Update, Delete), pencarian buku, hingga riwayat transaksi peminjaman.
Semua data tersimpan dalam file JSON, sehingga tetap aman meskipun aplikasi ditutup.

✨ Fitur Utama

Tambah Buku → Menambahkan data buku baru (judul & penulis)

Lihat Buku → Menampilkan semua buku beserta status peminjaman

Pinjam Buku → Mencatat buku yang dipinjam oleh pengguna tertentu

Kembalikan Buku → Mengubah status buku yang sudah dipinjam

Cari Buku → Mencari berdasarkan judul atau penulis

Hapus Buku → Menghapus data buku dari sistem

Riwayat Transaksi → Melihat daftar peminjaman dan pengembalian

Penyimpanan JSON → Semua data tersimpan otomatis di library.json

📂 Struktur Direktori
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

🗂️ Model Data Buku
type Book struct {
    ID         int    `json:"id"`
    Title      string `json:"title"`
    Author     string `json:"author"`
    IsBorrowed bool   `json:"is_borrowed"`
}


Model di atas digunakan untuk menyimpan informasi dasar mengenai sebuah buku.

🛠️ Cara Menjalankan

Pastikan Go versi 1.20 atau lebih baru sudah terpasang

Clone repositori:

git clone https://github.com/username/perpustakaan-mini.git


Masuk ke folder proyek:

cd perpustakaan-mini


Jalankan aplikasi:

go run main.go

🖥️ Tampilan Menu Utama
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

Tambah Buku → Masukkan judul & penulis, data otomatis tersimpan

Lihat Buku → Menampilkan semua daftar buku dan status pinjam

Pinjam Buku → Masukkan ID buku dan nama peminjam

Kembalikan Buku → Masukkan ID buku yang dikembalikan

Cari Buku → Gunakan kata kunci (judul/penulis) untuk pencarian

Hapus Buku → Hapus data buku dari sistem berdasarkan ID

Riwayat Transaksi → Menampilkan catatan peminjaman & pengembalian

💾 Penyimpanan Data

Semua data tersimpan di file library.json

File akan dibuat otomatis saat aplikasi pertama kali dijalankan

Saat keluar (menu nomor 8), data terbaru akan disimpan kembali

👨‍💻 Author

Nama: Yosafat Bagas Herlianka

Mata Kuliah: Pemrograman Jaringan

Keterangan: Aplikasi ini dibuat sebagai latihan implementasi Go language dengan konsep struct, fungsi, dan modularisasi.
