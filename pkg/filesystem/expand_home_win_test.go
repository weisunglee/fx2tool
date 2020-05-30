// +build windows

package filesystem

import (
	"fmt"
	"testing"
)

func TestExpandHome(t *testing.T) {
	home := ExpandHome("~")

	if len(home) > 0 {
		fmt.Println("home: ", home)
		t.Log("home: ", home)
	} else {
		t.Fatal("get home dir failed")
	}
}
