package evmdbapi

const (
  LATEST string = "latest"
  RPCServer string = "https://ropsten.infura.io/"
)

type RPCProvider struct {
  address string
  timeout int32
  secure  bool
}

var (
  MyProvider = RPCProvider{
    address: "ropsten.infura.io",
    timeout: 10,
    secure: true,
  }
)
