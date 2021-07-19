package sensitive

import (
	"fmt"
	"log"
	"testing"
)

func TestNewTrie(t *testing.T) {
	tree := NewTrie()

	err := tree.LoadDict("words/dict.txt")

	if err != nil {
		log.Fatal(err)
	}

	tree.add("hello")
	tree.BuildFailureLinks()
	fmt.Println(tree.Replace("sb，我的名字叫hello", '*'))
	fmt.Println(tree.Filter("你好，我的名字叫hello"))
	fmt.Println(tree.FindAll("sb，我的名字叫hello"))
}
