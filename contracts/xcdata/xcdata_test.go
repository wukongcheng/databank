package xcdata

import (
	"testing"

	"github.com/xcareteam/xci/accounts/abi/bind"
	"github.com/xcareteam/xci/accounts/abi/bind/backends"
	"github.com/xcareteam/xci/core"
	"github.com/xcareteam/xci/crypto"
	"math/big"
)

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	name   = "my xcdata"
	hash   = crypto.Keccak256Hash([]byte("my content"))
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

func TestGetXCData(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(1000000000)}})
	transactOpts := bind.NewKeyedTransactor(key)

	//_, err := NewWhiteList(transactOpts, addr, contractBackend)
	//if err != nil {
	//	t.Fatalf("can't bind the whitelist contract: %v", err)
	//}

	_, xcdata, err := DeployXCData(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't bind the whitelist contract: %v", err)
	}
	contractBackend.Commit()

	if _, err := xcdata.CommitData("unit_id1", "0xd7ac1f48c345561d6a273765b84b33454d3a8f5d"); err != nil {
		t.Fatalf("can't commit new data: %v", err)
	}
	contractBackend.Commit()

	len, err := xcdata.GetDataLength("unit_id1")
	if err != nil {
		t.Fatalf("can't get data length: %v", err)
	}

	if len.Int64() != 1 {
		t.Fatalf("GetDataLength error, expected %v, got %v", 1, len)
	}

	_, hash, err := xcdata.GetData("unit_id1", big.NewInt(0))
	if err != nil {
		t.Fatalf("can't get data: %v", err)
	}

	if hash != "0xd7ac1f48c345561d6a273765b84b33454d3a8f5d" {
		t.Fatalf("GetData error, expected %v, got %v", "0xd7ac1f48c345561d6a273765b84b33454d3a8f5d", hash)
	}
}
