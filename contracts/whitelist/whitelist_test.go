package whitelist

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
	name   = "my name on ENS"
	hash   = crypto.Keccak256Hash([]byte("my content"))
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

func TestGetNewWhiteList(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(1000000000)}})
	transactOpts := bind.NewKeyedTransactor(key)

	//_, err := NewWhiteList(transactOpts, addr, contractBackend)
	//if err != nil {
	//	t.Fatalf("can't bind the whitelist contract: %v", err)
	//}

	_, whitelist, err := DeployWhiteList(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't bind the whitelist contract: %v", err)
	}
	contractBackend.Commit()

	if _, err := whitelist.AddNewNode("enode://xxxxx@ip:port", "{name:silei}"); err != nil {
		t.Fatalf("can't add new node: %v", err)
	}
	contractBackend.Commit()

	did, err := whitelist.GetDID("enode://xxxxx@ip:port")
	if err != nil {
		t.Fatalf("can't get node: %v", err)
	}

	if did != "{name:silei}" {
		t.Fatalf("AddNewNode error, expected %v, got %v", "{name:silei}", did)
	}
}