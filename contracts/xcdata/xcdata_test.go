package xcdata

import (
	"testing"

	"github.com/wukongcheng/databank/accounts/abi/bind"
	"github.com/wukongcheng/databank/accounts/abi/bind/backends"
	"github.com/wukongcheng/databank/core"
	"github.com/wukongcheng/databank/crypto"
	"math/big"
	"github.com/wukongcheng/databank/crypto/ecies"
	"crypto/rand"
)

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	name   = "my xcdata"
	hash   = crypto.Keccak256Hash([]byte("my content"))
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

func TestPrivateKeyToPublicKey(t *testing.T) {
	privateKeyList := []string{
								"1643625ef3e1f6c1c58202888001f9714146b1eb95982a3f27c62057e2ffc80d",
								"ae999148ea50f4812f4644e21612ddfc833e7e97d28db67bf651d311b3a0a6e1",
								"4afbdd41a599b4ed69a2a8c9c6635cc66d655871efcaca51b3eb4fb9ffe6936e",
								"8c65bcaf2c70977394094630a4edc5ca4c758f926dab0a70beb2af9947449f05",
								"c0ce5fb9564dd2aab673a46f99fe493a5c929668006cf0667ee7f85a464b4683",
								"50e70a10d99cde3d67b9c9a25d51aeb6895b8f2c3a724aa0685020657122ef99",
								"de5ee37cf6c912867bc772be0dd7b36275509e1211525946efe8611d51989d5d",
								"e57ae2116b43f5bf7a845d55596afa72d04f837a7828c62173bcb5bd2bb30c39",
								"178e84c9202d874568f80c021cb0cf2ec3e22d7f92a8004ff82d231c073281ba",
								"e56b00def3fd6cca825684dca9218638908f12f9caeda7b2b0dc3b46ab6894b0",
								"e02a760f6878883cc1b69e475ca5a9e01245f0de7a86a49e0763fb8d505a75d6",
								"51ce3db5aaaa5879c4c907d40143648168a9c71c83968f5550c73c8d8922ab97",
								"7ee2ce33b07ed513840b77d79fb3053c17bfd1442fd28a95eb379ca117fb3ef5",
								}
	for i := 0 ; i < len(privateKeyList); i++ {
		key, _ = crypto.HexToECDSA(privateKeyList[i])
		t.Logf("04%x%x",key.PublicKey.X,key.PublicKey.Y)

	}
}
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

	AESKey := "12345678123456781234567812345678" //256bit

	eciesPrivate := ecies.ImportECDSA(key)

	encryptedAESKeyBytes,err := ecies.Encrypt(rand.Reader, &eciesPrivate.PublicKey, []byte(AESKey), nil, nil)
	if err != nil {
		t.Fatalf("Encrypt with doctor public key error %v", err.Error())
	}

	if _, err := xcdata.CommitData("unit_id1", "0xd7ac1f48c345561d6a273765b84b33454d3a8f5d",encryptedAESKeyBytes); err != nil {
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

	_, hash,AESKeyBytes, err := xcdata.GetData("unit_id1", big.NewInt(0))
	if err != nil {
		t.Fatalf("can't get data: %v", err)
	}

	if hash != "0xd7ac1f48c345561d6a273765b84b33454d3a8f5d" {
		t.Fatalf("GetData error, expected %v, got %v", "0xd7ac1f48c345561d6a273765b84b33454d3a8f5d", hash)
	}

	decryptedAESKey, err :=eciesPrivate.Decrypt(rand.Reader, AESKeyBytes, nil,nil)
	if err != nil {
		t.Fatalf("Decrypt with patient private key error: %+v",err.Error())
	}

	if string(decryptedAESKey[:]) != AESKey {
		t.Fatalf("Decrypt with patient private key error: %+v",err.Error())
	}
}
