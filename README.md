parse a string into a Map in GO.

# example
```go
package example

import (
  "fmt"
  json "github.com/matehaxor03/holistic_json/json"
)

func main() {

  data, errors := json.Parse("{\"hello\":\"world\"}")

  if errors != nil {
      fmt.Println(fmt.Errorf("%s", errors))
      return 1
  }

  // your code goes here, use the data Map

  return 0 
}
```
