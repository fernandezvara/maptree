package main

import (
	"encoding/json"
	"fmt"

	"github.com/fernandezvara/maptree"
)

type metadata struct {
	Owner string `json:"owner"`
}

func main() {

	t := maptree.New()

	t.Set("directory1/file1.txt", metadata{Owner: "user1"}, "this is the text for the file1")
	t.Set("directory1/file2.txt", metadata{Owner: "user1"}, "this is the text for the file2")
	t.Set("directory2/file3.txt", metadata{Owner: "user1"}, "this is the text for the file3")
	t.Set("directory2/file4.txt", metadata{Owner: "user1"}, "this is the text for the file4")
	t.Set("readme.md", metadata{Owner: "user3"}, "this is a readme file")

	b, err := json.MarshalIndent(t.Tree(), "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}
