package dbop

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	IntializePool()
	euser := GetUserInfo(1)
	fmt.Println(euser)
	defer ClosePool()
}
