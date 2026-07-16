parse a string into a Map in GO.

# example
```go
package example

import (
  "fmt"
  json "github.com/matehaxor03/holistic_json/json"
)

func main() {

  json, json_errors := json.Parse("{\"hello\":\"world\"}")

  if json_errors != nil {
      fmt.Println(fmt.Errorf("%s", json_errors))
      return 1
  }

  // your code goes here, use the json Map

  return 0 
}
```
