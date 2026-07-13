package main
import (
	"fmt"
	"encoding/json"
	"errors"
	"os"
)

type Task struct{
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
}

func saveTasks(tasks []Task){
	jsonData, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil{
		fmt.Print(err)
		return
	}

	text := []byte(jsonData)
	os.WriteFile("tasks.json", text, 0644)
}

func loadTasks(){
	_, err := os.Stat("tasks.json")
	
	if errors.Is(err, os.ErrNotExist){
		fmt.Println("file doesn't exits")
		return
	}

	data, err := os.ReadFile("tasks.json")

	if err != nil{
		fmt.Println(err)
		return
	}

	var task2 []Task

	err2 := json.Unmarshal(data, &task2)

	if err2 != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(task2)
}


func main(){
	task := []Task{
		{
			ID: 1,
			Description: "Learn Go.",
			Status: "todo",
		},
		{
			ID: 2,
			Description: "Learn CLI",
			Status: "done",
		},
	}

	saveTasks(task)

	fmt.Println("task saved")

	loadTasks()

}
