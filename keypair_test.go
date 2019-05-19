package keypair

import (
	"context"
	"testing"

	lbcf "github.com/lidstromberg/config"
)

func Test_KpCreate(t *testing.T) {
	ctx := context.Background()

	//create a config map
	bc := lbcf.NewConfig(ctx)

	//create a keypair
	_, err := NewKeyPair(ctx, bc)
	if err != nil {
		t.Fatal(err)
	}
}
func Test_KpEncrypt(t *testing.T) {
	ctx := context.Background()

	//create a config map
	bc := lbcf.NewConfig(ctx)

	//create a keypair
	kp, err := NewKeyPair(ctx, bc)
	if err != nil {
		t.Fatal(err)
	}

	teststring := "this is a test"

	result, err := kp.EncryptBytes(ctx, []byte(teststring))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("encrypted teststring: %s", result)
}
func Test_KpEncryptDecrypt(t *testing.T) {
	ctx := context.Background()

	//create a config map
	bc := lbcf.NewConfig(ctx)

	//create a keypair
	kp, err := NewKeyPair(ctx, bc)
	if err != nil {
		t.Fatal(err)
	}

	teststring := "this is a test"

	result1, err := kp.EncryptBytes(ctx, []byte(teststring))
	if err != nil {
		t.Fatal(err)
	}

	result2, err := kp.DecryptString(ctx, result1)
	if err != nil {
		t.Fatal(err)
	}

	if result2 != teststring {
		t.Fatalf("decrypted string did not match original test string")
	}
}
