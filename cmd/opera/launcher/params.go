package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://beb3f95429229b31f769e743ee4e341add132f4da59e1bb4e57ba2e8eaf9ecbe8d0f80a24a4d0e3ba67cf5ddfa704679d228681d32cc31a227cf8a451365c508@134.209.98.129:30399",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
