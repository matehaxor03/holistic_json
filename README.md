parse a string into a Map in GO.

# example
```go
package example

import (
  "fmt"
  "os"
  json "github.com/matehaxor03/holistic_json/json"
)

func main() {

  data, errors := json.Parse("{\"hello\":\"world\"}")

  if errors != nil {
      fmt.Println(fmt.Errorf("%s", errors))
      os.Exit(1)
  }

  // your code goes here, use the data Map
  os.Exit(0)
}
```
