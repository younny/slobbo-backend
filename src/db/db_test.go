package db

import (
	"slobbo/src/types"
	"testing"
)

var (
	testPost = types.Post{
		Title:    "Hello World",
		SubTitle: "Foo",
		Author:   "Jake",
		Category: 0,
		Body:     "ABC",
	}
)

func TestMain(m *testing.M) {

}
