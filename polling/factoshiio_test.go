// Copyright (c) of parts are held by the various contributors (see the CLA)
// Licensed under the MIT License. See LICENSE file in the project root for full license information.
package polling_test

import (
	"testing"
)

// Need an api key to run this
func TestFactoshiIORatesPeggedAssets(t *testing.T) {
	ActualDataSourceTest(t, "Factoshiio")
}

// TestFixedOpenExchangeRatesPeggedAssets tests all the crypto assets are found on OpenExchangeRates from fixed
func TestFactoshiIOPeggedAssets(t *testing.T) {
	FixedDataSourceTest(t, "Factoshiio", []byte(factoshiioData))
}

var factoshiioData = []byte(`
{"price":0.005108136,"volume":8338582.0,"quote":"USD","base":"PEG","updated_at":1574891112}
`)
