package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskManager struct {
	Tasks []Task
}

func (tm *TaskManager) AddTask(title string) {
	newTask := Task{
		ID:    len(tm.Tasks) + 1,
		Title: title,
		Done:  false,
	}
	tm.Tasks = append(tm.Tasks, newTask)
	fmt.Println("Task berhasil ditambahkan!")
}

func (tm *TaskManager) ShowTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("Belum ada task.")
		return
	}
	fmt.Println("Daftar Task:")
	for _, task := range tm.Tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
	}
}

func (tm *TaskManager) CompleteTask(id int) {
	for i := range tm.Tasks {
		if tm.Tasks[i].ID == id {
			tm.Tasks[i].Done = true
			fmt.Println("Task selesai!")
			return
		}
	}
	fmt.Println("Task tidak ditemukan.")
}

// Simpan task ke file JSON
func (tm *TaskManager) SaveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // biar rapi
	err = encoder.Encode(tm.Tasks)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
	}
}

// Load task dari file JSON
func (tm *TaskManager) LoadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		// Kalau file belum ada, tidak masalah
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tm.Tasks)
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
	}
}
