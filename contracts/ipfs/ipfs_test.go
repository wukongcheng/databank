package ipfs

import (
	"math/big"
	"github.com/xcareteam/xci/crypto"
	"testing"
	"github.com/xcareteam/xci/accounts/abi/bind/backends"
	"github.com/xcareteam/xci/core"
	"github.com/xcareteam/xci/accounts/abi/bind"
)

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

func TestIPFS(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(100000000000000)}})
	transactOpts := bind.NewKeyedTransactor(key)
	callOpts := &bind.CallOpts{
		From:	addr,
	}

	t.Logf("account: %v",addr.Hex())
	result, ipfs, err := DeployIPFS(transactOpts, callOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't bind the ipfs contract: %v", err)
	}
	t.Logf("contract address %s",result.String())

	contractBackend.Commit()

	owner,err := ipfs.GetOwner()
	if err != nil {
		t.Fatalf("can't bind the ipfs contract owner: %v", err)
	}

	if owner.Hex() != addr.Hex() {
		t.Logf("msg send %s",owner.Hex())
		t.Fatalf("Contract owner is wrong")
	}


	if _, err := ipfs.AddNewIpfsUrl("2018-03-15", "hello","/ipfs/QmXgBq2xJKMqVo8jZdziyudNmnbiwjbpAycy5RbfDBoJRM"); err != nil {
		t.Fatalf("can't add new ipfs url: %v", err)
	}
	contractBackend.Commit()

	if _, err := ipfs.AddNewIpfsUrl("2018-03-15", "blockchain","/ipfs/QmSQ3eNodmzvnBp9suPWFgZ4EhcgDr5dfZXQFnjAR3yZsT"); err != nil {
		t.Fatalf("can't add new ipfs url: %v", err)
	}
	contractBackend.Commit()

	did, err := ipfs.GetIpfsUrl("2018-03-15", "hello")
	if err != nil {
		t.Fatalf("can't get ipfs url: %v", err)
	}

	if did != "/ipfs/QmXgBq2xJKMqVo8jZdziyudNmnbiwjbpAycy5RbfDBoJRM" {
		t.Fatalf("AddNewIpfsUrl error, expected %v, got %v", "/ipfs/QmXgBq2xJKMqVo8jZdziyudNmnbiwjbpAycy5RbfDBoJRM", did)
	}

	fileNameList, err := ipfs.GetFileListByDate("2018-03-15")
	t.Logf("date: 2018-03-15, fileNameList: %s",fileNameList)

	url, err := ipfs.GetIpfsUrl("2018-03-15", "blockchain")
	if err != nil {
		t.Fatalf("can't get ipfs url: %v", err)
	}
	t.Logf("date: 2018-03-15, fileName: blockchain, ipfs url: %s",url)
}