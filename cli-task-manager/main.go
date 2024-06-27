package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct{
	ID	int	`json:"id"`
	Description	string	`json:"description"`
	Done	bool	`json:"done"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func main() {
	// var tasks =  []Task{}
	// tasks = append(tasks, Task{ID: 1, Description: "Sameer", Done: true})
	
}

// Add tasks to task.json
func addTask(taskList *TaskList, description string) {
	taskList.Tasks= append(taskList.Tasks, Task{
		ID: len(taskList.Tasks) + 1,
		Description: description,
		Done: false,
	})
}

// Load tasks from task.json file
func loadTask() TaskList {
	data, err := os.ReadFile("task.json")	// The data is []byte i.e., Bytes array
	if err != nil {
		return TaskList{}
	}
	fmt.Println(data)
	var taskList TaskList
	// The Bytes array has to be unmarshalled, meaning from json data to a type data
	json.Unmarshal(data, &taskList)
	return taskList
}

// Save tasks in task.json file
func saveData(taskList *TaskList) {
	data, _ := json.MarshalIndent(taskList, "", " ")
	os.WriteFile("task.json", data, 0644)
}

// List tasks from task.json file
func listTasks(taskList *TaskList) {
    for _, task := range taskList.Tasks {
        status := "[ ]"
        if task.Done {
            status = "[x]"
        }
        fmt.Printf("%d. %s %s\n", task.ID, status, task.Description)
    }
}

// Complete tasks written in task.json file
func completeTask(taskList *TaskList, idStr string) {
	// convert the id string to id integer
	var id int
	fmt.Scanf(idStr, "%d", id)

	for i:= range taskList.Tasks {
		if taskList.Tasks[i].ID == id {
			taskList.Tasks[i].Done = true
			fmt.Println("Task completed: ", taskList.Tasks[i].Description)
			return
		}
	}
	fmt.Println("Task not found")
}
