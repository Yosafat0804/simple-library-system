package main

import (
	"fmt"
	"todo-app/todo"
)

func main() {
	var manager todo.TaskManager

	// Load data dari file saat start
	manager.LoadFromFile("tasks.json")

	var pilihan int
	for {
		fmt.Println("\n=== Aplikasi To-Do List ===")
		fmt.Println("1. Tambah Task")
		fmt.Println("2. Lihat Semua Task")
		fmt.Println("3. Tandai Task Selesai")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var title string
			fmt.Print("Masukkan nama task: ")
			fmt.Scan(&title)
			manager.AddTask(title)
		case 2:
			manager.ShowTasks()
		case 3:
			var id int
			fmt.Print("Masukkan ID task: ")
			fmt.Scan(&id)
			manager.CompleteTask(id)
		case 4:
			// Simpan data sebelum keluar
			manager.SaveToFile("tasks.json")
			fmt.Println("Data disimpan ke tasks.json")
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
