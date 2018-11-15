package keypair

import (
	"log"
	"os"

	lbcf "github.com/lidstromberg/config"

	"golang.org/x/net/context"
)

//preflight config checks
func preflight(ctx context.Context, bc lbcf.ConfigSetting) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.Println("Started Keypair preflight..")

	//get the config and apply it to the config
	bc.LoadConfigMap(ctx, preflightConfigLoader())

	if bc.GetConfigValue(ctx, "EnvKpGcpBucket") == "" {
		log.Fatal("Could not parse environment variable EnvKpGcpBucket")
	}

	if bc.GetConfigValue(ctx, "EnvKpPrivateKey") == "" {
		log.Fatal("Could not parse environment variable EnvKpPrivateKey")
	}

	if bc.GetConfigValue(ctx, "EnvKpPublicKey") == "" {
		log.Fatal("Could not parse environment variable EnvKpPublicKey")
	}

	if bc.GetConfigValue(ctx, "EnvKpPrivateKeyCredential") == "" {
		log.Fatal("Could not parse environment variable EnvKpPrivateKeyCredential")
	}

	if bc.GetConfigValue(ctx, "EnvKpType") == "" {
		log.Fatal("Could not parse environment variable EnvKpType")
	}

	log.Println("..Finished Keypair preflight.")
}

//preflightConfigLoader loads the config vars
func preflightConfigLoader() map[string]string {
	cfm := make(map[string]string)

	//EnvKpGcpBucket is the cloud storage bucket to target
	cfm["EnvKpGcpBucket"] = os.Getenv("KP_GCP_BUCKET")
	//EnvKpPrivateKey is the jwt private key
	cfm["EnvKpPrivateKey"] = os.Getenv("KP_PRIKEY")
	//EnvKpPublicKey is the jwt public key
	cfm["EnvKpPublicKey"] = os.Getenv("KP_PUBKEY")
	//EnvKpPrivateKeyCredential credential for the jwt private key
	cfm["EnvKpPrivateKeyCredential"] = os.Getenv("KP_PRIKEYP")
	//EnvKpType is the source location of the keypairs (local/bucket)
	cfm["EnvKpType"] = os.Getenv("KP_TYPE")

	if cfm["EnvKpGcpBucket"] == "" {
		log.Fatal("Could not parse environment variable EnvKpGcpBucket")
	}

	if cfm["EnvKpPrivateKey"] == "" {
		log.Fatal("Could not parse environment variable EnvKpPrivateKey")
	}

	if cfm["EnvKpPublicKey"] == "" {
		log.Fatal("Could not parse environment variable EnvKpPublicKey")
	}

	if cfm["EnvKpPrivateKeyCredential"] == "" {
		log.Fatal("Could not parse environment variable EnvKpPrivateKeyCredential")
	}

	if cfm["EnvKpType"] == "" {
		log.Fatal("Could not parse environment variable EnvKpType")
	}

	return cfm
}
