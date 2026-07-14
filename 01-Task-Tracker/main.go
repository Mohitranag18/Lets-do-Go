package main
import (
	"fmt"
	"encoding/json"
	"os"
)

type Task struct{
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
}

func saveTasks(tasks []Task) error{
	jsonData, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil{
		return err
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil{
		return err
	}
	return nil
}

func loadTasks() ([]Task, error){
	_, err := os.Stat("tasks.json")
	
	if err != nil{
		if errors.Is(err, os.ErrNotExist){
			return []Task{}, nil
		}
		return nil, err
	}

	data, err := os.ReadFile("tasks.json")

	if err != nil{
		return nil, err
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)

	if err != nil{
		return nil, err
	}

	return tasks, nil
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

	err := saveTasks(task)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("task saved")

	tasks, err := loadTasks()

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(tasks)

}
