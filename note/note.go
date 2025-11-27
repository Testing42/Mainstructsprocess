package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// for marshall to work all the items in struct must be capital
/* struct tags are meta data
that can be added to structs below
so that fields can be given any
keta data you want be if you don't
have code that can manage that meta
data then nothing will happen.
*/
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

//this become a method by adding
//(somethinginside) that passes cvalues from the note struct
// the word before the name of the struct is a nickname
// that

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func NewNote(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("invalid input")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
