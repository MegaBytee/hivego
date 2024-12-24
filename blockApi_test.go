package hivego

import (
	"fmt"
	"testing"
)

func TestVirtualOps(t *testing.T) {
	rpc := NewHiveRpc("https://api.hive.blog")
	virtualOps, err := rpc.FetchVirtualOps(88386873, true, false)

	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(virtualOps)
}
