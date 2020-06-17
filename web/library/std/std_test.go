package std_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestReadDir(t *testing.T) {
	s := func(num int) string {
		s := ""
		for i := 0; i < num; i++ {
			s += "\t"
		}
		return s
	}
	listDir("/project/go-web/web", 0, s)
}

func listDir(path string, level int, s func(int) string) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range dir {
		if v.IsDir() {
			fmt.Printf("%s %s %v \n", s(level), v.Mode(), v.Name())
			level++
			listDir(path+"/"+v.Name(), level, s)
			level--
		} else {
			fmt.Printf("%s %v %s %dbyte\n", s(level), v.Mode(), v.Name(), v.Size())
		}
	}
}

type User struct {
	ID       uint64
	Name     string
	Age      uint8
	IsAdmin bool
	Acccount *Account
}

type Account struct {
	Account  string
	Password string
}

func TestPrintf(t *testing.T) {
	u := User{
		ID:   1,
		Name: "jj",
		Age:  12,
		IsAdmin: false,
		Acccount: &Account{
			Account:  "123421",
			Password: "trqew",
		},
	}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)
	fmt.Printf("%T\n", u)
}
