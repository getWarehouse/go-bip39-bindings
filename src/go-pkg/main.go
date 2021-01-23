package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func NewCrypto() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
				return "Invalid no of arguments passed"
		}

		entropy, _ := bip39.NewEntropy(args[0].Int())
		mnemonic, _ := bip39.NewMnemonic(entropy)

		seed := bip39.NewSeed(mnemonic, args[1].String())

		masterKey, _ := bip32.NewMasterKey(seed)
		publicKey := masterKey.PublicKey()

		data := map[string]interface{}{
			"entropy": entropy, 
			"mnemonic": mnemonic,
			"seed": seed,
			"masterKey": masterKey,
			"publicKey": publicKey,
		}    
		final := removeNils(data)

		out, _ := json.Marshal(final)
		return js.ValueOf(string(out))
	})
	return jsonFunc
}

func removeNils(initialMap map[string]interface{}) map[string]interface{} {
    withoutNils := map[string]interface{}{}
    for key, value := range initialMap {
        _, ok := value.(map[string]interface{})
        if ok {
            value = removeNils(value.(map[string]interface{}))
            withoutNils[key] = value
            continue
        }
        if value != nil {
            withoutNils[key] = value
        }
    }
    return withoutNils
}

func main() {  
	js.Global().Set("NewCrypto", NewCrypto())
	select {}
}