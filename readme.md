# maptree 

This is a helper library that can be useful if you need to describe data in a tree. 

A directory with files to show or any other data that can be structured in that way (documentation, etc).


*This is not a radix tree and does not pretend it in any way.*

Sample usage:

```
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
```


This is the JSON formatted output: 

(this can be easily walked on your javascript applications, as example)

``` 
{
  "items": {
    "directory1": {
      "items": {
        "file1.txt": {
          "meta": {
            "owner": "user1"
          },
          "data": "this is the text for the file1"
        },
        "file2.txt": {
          "meta": {
            "owner": "user1"
          },
          "data": "this is the text for the file2"
        }
      }
    },
    "directory2": {
      "items": {
        "file3.txt": {
          "meta": {
            "owner": "user1"
          },
          "data": "this is the text for the file3"
        },
        "file4.txt": {
          "meta": {
            "owner": "user1"
          },
          "data": "this is the text for the file4"
        }
      }
    },
    "readme.md": {
      "meta": {
        "owner": "user3"
      },
      "data": "this is a readme file"
    }
  }
}
```

