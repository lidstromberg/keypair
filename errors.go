package keypair

import "errors"

//errors
var (
	//ErrKeyPairNotExist occurs if the key pair cannot be read
	ErrKeyPairNotExist = errors.New("Keypair could not be created")
)
