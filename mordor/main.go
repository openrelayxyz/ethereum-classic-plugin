package main

import (
	"context"
	"math/big"
	"path/filepath"
	"strings"

	"github.com/openrelayxyz/plugeth-utils/core"
	"github.com/openrelayxyz/plugeth-utils/restricted"
)

var (

	ClassicBootnodes = []string{
		"enode://5e85df7bc6d529647cf9a417162784a89b7ccf2b8e1570fadb6fdf9fa025c8ec2257825d1ec5d7357a6f49898fdfbd9c4c56d22645dbe8b8a6aa67dacbcf3ecc@157.230.152.87:30303", // meowsbits@sfetclabs
		"enode://4539a067ae1f6a7ffac509603ba37baf772fc832880ddc67c53f292b6199fb048267f0311c820bc90bfd39ec663bc6b5256bdf787ec38425c82bde6bc2bcfe3c@127.0.0.1:30303",      // @etccoop-sfo
	}

	ClassicDNSNetwork1 = dnsPrefixETC + "all.mordor.blockd.info"

	dnsPrefixETC string = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@"

	snapDiscoveryURLs []string

	forkBlockIds = []uint64 {301243, 999983, 2520000, 3985893, 5520000, 9957000}                        

	forkTimeIds = []uint64{}
)

type ClassicService struct {
	backend core.Backend
	stack   core.Node
}

var (
	pl      core.PluginLoader
	backend restricted.Backend
	log     core.Logger
	events  core.Feed
)

var (
	httpApiFlagName = "http.api"
	mainnetFlag = "mainnet"
	goerliFlag = "goerli"
	sepoliaFlag = "sepolia"
	holeskyFlag = "holesky"

	networkPanicMsg = "This node is optimized to run the Ethereum Classic Network only, check datadir/plugins/ for a classic.so binary and remove it if this is not the desired behavior"
)

func Initialize(ctx core.Context, loader core.PluginLoader, logger core.Logger) { 
	pl = loader
	events = pl.GetFeed()
	log = logger
	v := ctx.String(httpApiFlagName)
	if v != "" {
		ctx.Set(httpApiFlagName, v+",plugeth")
	} else {
		ctx.Set(httpApiFlagName, "eth,net,web3,plugeth")
		
	}

	switch {
		case ctx.Bool(mainnetFlag):
			panic(networkPanicMsg)
		case ctx.Bool(goerliFlag):
			panic(networkPanicMsg)
		case ctx.Bool(sepoliaFlag):
			panic(networkPanicMsg)
		case ctx.Bool(holeskyFlag):
			panic(networkPanicMsg)
	}


	log.Info("Loaded Mordor testnet plugin")
}

func Is1559(*big.Int) bool {
	return false
}

func Is160(num *big.Int) bool {
	r := num.Cmp(big.NewInt(0))
	return r >= 0
}

func IsShanghai(num *big.Int) bool {
	r := num.Cmp(big.NewInt(9957000))
	return r >= 0
}

func InitializeNode(node core.Node, backend restricted.Backend) {
	db := backend.ChainDb()

	cfg := []byte(`{
		"chainId": 63,
		"networkId": 7,
		"homesteadBlock": 1150000,
		"daoForkBlock": null,
		"daoForkSupport": false,
		"eip150Block": 2500000,
		"eip155Block": 3000000,
		"eip158Block": 8772000,
		"byzantiumBlock": 8772000,
		"constantinopleBlock": 9573000,
		"petersburgBlock": 9573000,
		"istanbulBlock": 10500839,
		"berlinBlock": 13189133,
		"londonBlock": 14525000,
		"ethash": {}
	}`)

	hash := core.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")

	if err := db.Put(append([]byte("ethereum-config-"), hash.Bytes()...), cfg); err != nil {
		log.Error("Error loading Classic config", "err", err)
	}
}

func GetAPIs(stack core.Node, backend core.Backend) []core.API {
	return []core.API{
		{
			Namespace: "plugeth",
			Version:   "1.0",
			Service:   &ClassicService{backend, stack},
			Public:    true,
		},
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   &API{eHashForAPI},
			Public:    true,
		},
	}
}

func ForkIDs([]uint64, []uint64) ([]uint64, []uint64) {
	return forkBlockIds, forkTimeIds
}

func SetDefaultDataDir(path string) string {
	return filepath.Join(path, "mordor")
}

func OpCodeSelect() []int {
	codes := []int{0x48}
	return codes
}

func SetNetworkId() *uint64 {
	var networkId *uint64
	classicNetworkId := uint64(1)
	networkId = &classicNetworkId
	return networkId 
}

func SetBootstrapNodes() []string {
	result := ClassicBootnodes
	return result
}

func SetETHDiscoveryURLs(lightSync bool) []string {

	url := ClassicDNSNetwork1
	if lightSync == true {
		url = strings.ReplaceAll(url, "all", "les")
	}
	result := []string{url}
	snapDiscoveryURLs = result

	return result
}

func SetSnapDiscoveryURLs() []string {
	return snapDiscoveryURLs
}

func (service *ClassicService) Test(ctx context.Context) string {
	return "total classic"
}
