package auth

import (
	"fmt"
	"math/rand"
	"strings"

	"main/common/kv"
)

type User struct {
	Id    int
	Unit  int
	Name  string
	Role  string
	Token string
}

func newUser(item kv.Element) *User {
	v := &User{}
	v.Id = int(item.GetProperty("id").GetNumber())
	v.Unit = int(item.GetProperty("unit").GetNumber())
	v.Name = item.GetProperty("fullname").GetString()
	v.Role = item.GetProperty("role").GetString()
	return v
}

func toUser(v interface{}) *User {
	if u, y := v.(*User); y {
		return u
	}
	return nil
}

var users map[string]*User = make(map[string]*User)

func accMix(acc string) string {
	mix := fmt.Sprintf("%d", rand.Intn(100000))
	return acc + strings.Repeat("0", 6-len(mix)) + mix
}
func accTidy(acc string) string {
	l := len(acc) - 6
	return acc[:l]
}
