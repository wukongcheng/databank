package vote

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
	name   = "my vote"
	hash   = crypto.Keccak256Hash([]byte("my content"))
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)


func TestGetVote(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(1000000000)}})
	transactOpts := bind.NewKeyedTransactor(key)

	_, vote, err := DeployVote(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't bind the vote contract: %v", err)
	}
	contractBackend.Commit()

	var a,b,c [32]byte
	copy(a[:], "a")
	copy(b[:], "b")
	copy(c[:], "c")

	var p = [][32]byte{a, b, c};
	if _, err := vote.NewBallot("TestBallot",p); err != nil {
		t.Fatalf("can't add new node: %v", err)
	}
	contractBackend.Commit()

	_, err = vote.Vote(big.NewInt(0), big.NewInt(1))
	if err != nil {
		t.Fatalf("can't vote: %v", err)
	}

	name, err := vote.GetWinnerName(big.NewInt(0))
	if err != nil {
		t.Fatalf("can't getWinnerName: %v", err)
	}

	if name != b {
		t.Fatalf("Vote error, expected %v, got %v", "b", string(name[:]))
	}
}
