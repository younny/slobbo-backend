package tests

import (
	"testing"

	"github.com/younny/slobbo-backend/src/types"
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
