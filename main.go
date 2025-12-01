package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-starting-module/note"
	"example.com/go-starting-module/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.NewTodo(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving the error failed.")
		return
	}

	fmt.Println("Saving the todo succeeded!")

	//above is the todo work and below is the note work

	userNote, err := note.NewNote(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the error failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")

}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	//printf is more versitle than print and allows you to edit the input more
	/*scan is used for only single word inputs
	//because of this it is better to use for longer text input
	//bufio
	fmt.Scanln(&value)
	os.stdin is just the commandline
	*/
	reader := bufio.NewReader(os.Stdin)

	/*this /n tells where the reading of user input would stop
	Also you have to use single quotes here '' otherwise
	this won't work
	*/
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	/*even though above stops reading at line break
	The line break is still there but the command below
	removes that line break.
	on windows a line break is made with /r/n so you have to remove
	both
	*/
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
