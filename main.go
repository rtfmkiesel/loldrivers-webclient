//go:build js

package main

import (
	"encoding/json"
	"fmt"
	"loldrivers-webclient/internal/loldrivers"
	"syscall/js"
)

// Wrapper for loldrivers.MatchHash
func MatchHashWrapper(this js.Value, p []js.Value) interface{} {
	hash := p[0].String()
	if hash == "" {
		return map[string]interface{}{
			"driver": nil,
			"ok":     false,
		}
	}

	// Fetch driver
	ok, driver := loldrivers.MatchHash(hash)
	if !ok {
		// No driver
		return map[string]interface{}{
			"driver": nil,
			"ok":     false,
		}
	}

	// To JSON
	driverJson, err := json.Marshal(driver)
	if err != nil {
		return map[string]interface{}{
			"driver": nil,
			"ok":     false,
		}
	}

	return map[string]interface{}{
		"driver": string(driverJson),
		"ok":     true,
	}
}

func main() {
	fmt.Println("Loading drivers")
	if err := loldrivers.LoadDrivers(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Registering lookupHash")
	js.Global().Set("lookupHash", js.FuncOf(MatchHashWrapper))

	// Do not quit
	select {}
}
