package http

import (
	"errors"
	"fmt"
	"testing"
)

func TestAssertDesc(t *testing.T) {
	{
		str := AssertDesc("123")
		fmt.Println(str)
	}
	{
		str := AssertDesc(errors.New("ddd"))
		fmt.Println(str)
	}
	{
		str := AssertDesc(nil)
		fmt.Println(str)
	}
}
