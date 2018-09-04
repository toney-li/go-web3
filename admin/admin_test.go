package admin

import (
	"testing"
	"github.com/toney-li/go-web3/providers"
	"fmt"
)

func TestAdmin_DataDir(t *testing.T) {

	p := providers.NewHTTPProvider("127.0.0.1:8547", 10, false)
	admin := NewAdmin(p)
	fmt.Println(admin.DataDir())
}
