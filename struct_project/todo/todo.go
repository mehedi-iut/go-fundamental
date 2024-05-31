package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text   string `json:"text"`
}

func New(text string)(Todo, error){
	if text == ""{
		return Todo{}, errors.New("Invalid input")
	}
	
	return Todo{
		Text:   text,
	}, nil
}


func (todo Todo) Display(){
	fmt.Printf("Your note title: %v\nYour note Content: %v\n", todo.Text)
}

func (todo Todo) Save() error{
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo)
	if err != nil{
		return err
	}
	
	return os.WriteFile(fileName, jsonData, 0644)
	
}