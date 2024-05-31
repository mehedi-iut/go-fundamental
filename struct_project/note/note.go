package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string)(Note, error){
	if title == "" || content == ""{
		return Note{}, errors.New("Invalid input")
	}
	
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}


func (note Note) Display(){
	fmt.Printf("Your note title: %v\nYour note Content: %v\n", note.Title, note.Content)
}

func (note Note) Save() error{
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(note)
	if err != nil{
		return err
	}
	
	return os.WriteFile(fileName, jsonData, 0644)
	
}