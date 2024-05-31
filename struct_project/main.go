package main

import (
	"bufio"
	"example.com/note/note"
	"example.com/note/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface{
	Save() error
}

type outputable interface{
	saver
	Display()
}

func main(){
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	
	todo, err := todo.New(todoText)
	
	if err != nil{
		fmt.Println(err)
		return
	}
	
	userNote, err := note.New(title, content)
	
	if err != nil{
		fmt.Println("Invalid input")
		return
	}
	
	err = outputData(todo)
	if err != nil{
		return
	}
	
	err = outputData(userNote)
	
	if err != nil{
		return
	}
	
}


func outputData(data outputable) error {
	data.Display()
	return data.Save()
}
func saveData(data saver) error{
	err := data.Save()
	
	if err != nil{
		fmt.Println("Saving failed.")
		return err
	}
	
	fmt.Println("Saving succeeded!")
	return nil
}

func getNoteData()(string, string){
	title := getUserInput("Note title: ")
	
	content := getUserInput("Note Content: ")
	
	return title, content
}

func getUserInput(prompt string)(string){
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil{
		return ""
	}
	
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}