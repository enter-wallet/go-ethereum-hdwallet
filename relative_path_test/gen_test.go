package genaddr

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

const pubkey string = "xpub6ByFWEnf1iAeLdt6Wz9yNjacgod9swmiBkASY5bSrgso186DLAoRFYnyR6YcymUEaSyyt8RqUQNpv3FNzhZQHkdh3EiFyEZpgCHVVgBkana"
const mne string = "buddy kidney measure angry good melody cancel erosion exercise two photo girl plastic act marriage field pair ripple canal just number nominee earn document"

func TestGenRelativePub(t *testing.T) {
	//wallet, err := hdwallet.NewFromMnemonic(mne)//if no password, use this
	seed := bip39.NewSeed(mne, "ganPsw12345678")
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'")
	_, relatedPub, err := wallet.DeriveKeyFromPath(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("related pubstring:", relatedPub) //this is to gen the relative Pub for TestEthAddrGen
	//refTO https://iancoleman.io/bip39/ Account Extended Public Key
}
func TestEthAddrGen(t *testing.T) { //Derived Addresses
	f, err := os.OpenFile("contract_addresses.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	start := 0
	for index := start; index < 100+start; index++ {
		addr, err := hdwallet.GenAddrFromPub(pubkey, "0/"+strconv.FormatUint(uint64(index), 10))
		if err != nil {
			t.Error(err)
		}
		f.WriteString(addr)
		f.WriteString(",\n")
	}
}
