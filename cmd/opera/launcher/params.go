package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://beb3f95429229b31f769e743ee4e341add132f4da59e1bb4e57ba2e8eaf9ecbe8d0f80a24a4d0e3ba67cf5ddfa704679d228681d32cc31a227cf8a451365c508@134.209.98.129:30399", //digital ocean singapore
		"enode://93efbfd654a890000cbba98a97aa60818807a9d2c104c852a3a3bff751878765d476131e496ef4753987ac48f7078cfe1eaf89864ba11661df6020e1487f9181@34.126.88.133:30399", //google asia-southeast1-b
		"enode://0121dcf0a7ac82b607fd54b274976e3503209b6f6cbb2f370bc281c3f79e7106b0d4273bfbe0465e113855814e50179b93d7a4a236055c5e671d3bd987c11709@34.107.57.65:30399", //google europe-west-a
		"enode://d6efe29b88223cd5599fa0d1f65e44faaf1259e3465c00f0e153bba792874b8a1d7d356baef1bdf1cd2f94e8a8c08bb8ed5994a9c4c116e605508edc23271ac9@34.87.197.240:30399", //google australia-southeast1-b
		"enode://e6eb6f81050b212bcc80b408555969e36825bb15ad12be6e86c9930c8f9709c7f91651c04ea449c342c3f1d87e0104c6ed26b93726cc7ad533ffee308bb045da@34.150.87.60:30399", //google asia-east2-a
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
