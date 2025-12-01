package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// for marshall to work all the items in struct must be capital
/* struct tags are meta data
that can be added to structs below
so that fields can be given any
keta data you want be if you don't
have code that can manage that meta
data then nothing will happen.
*/
type Todo struct {
	Text string `json:"Text"`
}

//this become a method by adding
//(somethinginside) that passes cvalues from the note struct
// the word before the name of the struct is a nickname
// that

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func NewTodo(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("invalid input")
	}

	return &Todo{
		Text: content,
	}, nil
}
