// example
package main

import (
	"github.com/go-shafaq/defcase_test"
	"github.com/go-shafaq/encoding-json"
)

func main() {
	defcase_test.Marshal = json.Marshal
	defcase_test.Unmarshal = json.Unmarshal
	defcase_test.DefCase = json.DefCase
	defcase_test.SnakeCaseTest()
}
