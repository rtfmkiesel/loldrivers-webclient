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
	// Get the button from the DOM
	jsDoc := js.Global().Get("document")
	folderButton := jsDoc.Call("getElementById", "select-folder-button")

	fmt.Println("Loading drivers")
	if err := loldrivers.LoadDrivers(); err != nil {
		fmt.Println(err)
		folderButton.Set("innerHTML", "There was an error fetching the driver data, check the console")
		return
	}

	fmt.Println("Registering lookupHash function")
	js.Global().Set("lookupHash", js.FuncOf(MatchHashWrapper))

	// Enable the button if we get here
	folderButton.Set("innerHTML", "Select folder")
	folderButton.Set("disabled", false)

	// Do not quit
	select {}
}
