# 📚 Sistem Perpustakaan Sederhana

Aplikasi **console-based** yang dibuat menggunakan bahasa pemrograman **Go**. Program ini memungkinkan pengguna untuk mengelola data buku dengan fitur **CRUD (Create, Read, Update, Delete)**, pencarian buku, hingga riwayat transaksi peminjaman. Semua data tersimpan dalam file **JSON**, sehingga tetap aman meskipun aplikasi ditutup.

---

## ✨ Fitur Utama

- ➕ **Tambah Buku** : menambahkan data buku baru (judul & penulis)  
- 📖 **Lihat Buku** : menampilkan semua buku beserta status peminjaman  
- 📦 **Pinjam Buku** : mencatat buku yang dipinjam oleh pengguna tertentu  
- 🔄 **Kembalikan Buku** : mengubah status buku yang sudah dipinjam  
- 🔍 **Cari Buku** : mencari berdasarkan judul atau penulis  
- 🗑️ **Hapus Buku** : menghapus data buku dari sistem  
- 📜 **Riwayat Transaksi** : melihat daftar peminjaman dan pengembalian  
- 💾 **Penyimpanan JSON** : semua data tersimpan otomatis di `library.json`  

---

## 🗂️ Model Data Buku
```
type Book struct {
    ID         int    `json:"id"`
    Title      string `json:"title"`
    Author     string `json:"author"`
    IsBorrowed bool   `json:"is_borrowed"`
}
Model di atas digunakan untuk menyimpan informasi dasar mengenai sebuah buku.
```
## 🛠️ Cara Menjalankan
```
1. Pastikan Go versi 1.25.1 atau lebih baru sudah terpasang

2. Clone repositori:
git clone https://github.com/username/perpustakaan-mini.git

3. Masuk ke folder proyek:
cd perpustakaan-mini

4. Jalankan aplikasi:
go run main.go
```
## 🖥️ Tampilan Menu Utama
```
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
```
## 📖 Panduan Singkat
```
- Tambah Buku → masukkan judul & penulis, lalu data otomatis tersimpan
- Lihat Buku → tampilkan semua daftar buku dan status pinjam
- Pinjam Buku → masukkan ID buku dan nama peminjam
- Kembalikan Buku → masukkan ID buku yang dikembalikan
- Cari Buku → gunakan kata kunci (judul/penulis) untuk pencarian
- Hapus Buku → hapus data buku dari sistem berdasarkan ID
- Riwayat Transaksi → tampilkan catatan peminjaman & pengembalian
```
## 💾 Penyimpanan Data
```
- Semua data tersimpan di file library.json
- File akan dibuat otomatis saat aplikasi pertama kali dijalankan
- Saat keluar (menu nomor 8), data terbaru akan disimpan kembali
```
## 👨‍💻 Author
```
- Nama: Yosafat Bagas Herlianka
- Mata Kuliah: Pemrograman Jaringan
- Keterangan: Aplikasi ini dibuat sebagai latihan implementasi Go language dengan konsep struct, fungsi, dan modularisasi
```
