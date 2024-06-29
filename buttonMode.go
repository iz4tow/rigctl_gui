package trash

import (
  "fmt"
  "strings"
)

func main() {
  data := "ACCA\n1234"
  A := strings.SplitN(data, "\n", 2)

  fmt.Println(A[0])
}

