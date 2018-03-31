package AES

import (
	"testing"
)

func TestAES(t *testing.T) {

	AESKey := "12345678123456781234567812345678" //256bit

	data := []byte("1234567_1234567_1234567_1234567_1234567_1234567_1234567_1234567_")
	encryptedData,err := Encrypt([]byte(AESKey),data)
	if err != nil {
		t.Fatalf("AES encryption: %v", err)
	}
	originalData,err := Decrypt([]byte(AESKey),encryptedData)
	if string(originalData[:])!=string(data[:]){
		t.Fatalf("AES decrypt result is not match to original string")
	}
}