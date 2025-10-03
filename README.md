# 📚 Aplikasi To-Do List Sederhana

Aplikasi console-based yang dibuat menggunakan bahasa pemrograman Go. Program ini memungkinkan pengguna untuk mengelola daftar tugas dengan fitur CRUD (Create, Read, Update, Delete), menandai tugas selesai, hingga menyimpan data secara permanen dalam file JSON sehingga daftar tetap tersimpan meskipun aplikasi ditutup.

---

## ✨ Fitur Utama

➕ Tambah Task : menambahkan tugas baru ke dalam daftar

📖 Lihat Task : menampilkan semua tugas beserta status selesai/belum

✅ Tandai Selesai : mengubah status task menjadi selesai

🗑️ Hapus Task : menghapus task dari daftar (opsional, bisa ditambah)

💾 Penyimpanan JSON : semua data tersimpan otomatis di tasks.json

---

## 🗂️ Model Data Task
```
type Task struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}
Struct di atas digunakan untuk menyimpan informasi dasar mengenai sebuah task.
```
## 🛠️ Cara Menjalankan
```
1. Pastikan Go versi 1.25.1 atau lebih baru sudah terpasang

2. Clone repositori:
git clone https://github.com/username/todo-app.git

3. Masuk ke folder proyek:
cd todo-app

4. Jalankan aplikasi:
go run main.go
```
## 🖥️ Tampilan Menu Utama
```
===============================
   APLIKASI TO-DO LIST
===============================
1. Tambah Task
2. Lihat Semua Task
3. Tandai Task Selesai
4. Keluar
```
## 📖 Panduan Singkat
```
- Tambah Task → masukkan nama tugas, lalu data otomatis tersimpan
- Lihat Task → tampilkan semua daftar tugas dengan status selesai/belum
- Tandai Task Selesai → pilih ID task untuk mengubah status menjadi selesai
- Keluar → otomatis menyimpan semua data ke file tasks.json
```
## 💾 Penyimpanan Data
```
- Semua data tersimpan di file tasks.json
- File akan dibuat otomatis saat aplikasi pertama kali dijalankan
- Saat keluar (menu nomor 4), data terbaru akan disimpan kembali
```
## 👨‍💻 Author
```
- Nama: Yosafat Bagas Herlianka
- Mata Kuliah: Pemrograman Jaringan
- Keterangan: Aplikasi ini dibuat sebagai latihan implementasi Go language dengan konsep struct, fungsi, dan modularisasi
```
