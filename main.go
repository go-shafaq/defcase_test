package defcase_test

import (
	"encoding/json"
	"fmt"
	"github.com/go-shafaq/defcase"
)

var (
	Marshal   = json.Marshal
	Unmarshal = json.Unmarshal
	DefCase   = func(f func(tag string) func(pkgPath, name string) string) {}
)

func SnakeCaseTest() {
	DefCase(defcase.DefCaseFn)
	defcase.DefCase("json", "*", defcase.Snak_case)
	id, firstName := 13, "Jon"
	userJs := fmt.Sprintf(`{"id":%d, "first_name":"%s"}`, id, firstName)
	u := new(User)
	err := Unmarshal([]byte(userJs), u)
	if err != nil {
		panic(err)
	}
	if u.Id != id || u.FirstName != firstName {
		panic(fmt.Sprintf(`%v, id=%d, fn=%s`, u, id, firstName))
	}

	fmt.Println("SUCCESS")
	fmt.Println("DONE")
}

type User struct {
	Id        int
	FirstName string
}
