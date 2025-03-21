/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"cryptoSpy/HTTP"
	ui "cryptoSpy/UI"
	"cryptoSpy/cmd"
)

func main() {
	assets := HTTP.GET("https://rest.coinapi.io/v1/assets", nil)
	cmd.Execute()
	ui.RenderList(mapAssets(assets), "Crypto")
}

func mapAssets(jsonStr string) map[string]string {
	type Asset struct {
		AssetID    string  `json:"asset_id"`
		Name       string  `json:"name"`
		IsCrypto   int8    `json:"type_is_crypto"`
		Volume_1hr float32 `json:"volume_1hrs_usd"`
		PriceUsd   float64 `json:"price_usd"`
	}

	// Unmarshal into a slice of Asset
	var assets []Asset
	err := json.Unmarshal([]byte(jsonStr), &assets)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Create a map[string]string to store asset_id -> name
	assetMap := make(map[string]string)
	for _, asset := range assets {
		if asset.IsCrypto == 1 {
      assetMap[asset.AssetID + " | " + asset.Name] = "VOL 1h : " + strconv.FormatFloat(float64(asset.Volume_1hr), 'f', 8, 64)
		}
	}
	return assetMap
}
