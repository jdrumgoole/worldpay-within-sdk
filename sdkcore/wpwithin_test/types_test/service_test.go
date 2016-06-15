package types_test
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"flag"
"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/types"
)

func TestMain(m *testing.M) {

	flag.Parse()

	// Setup here

	// Run the test and store the result code to be returned after teardown
	resultCode := m.Run()

	// Teardown here

	// Return the test result code
	os.Exit(resultCode)
}

func TestAddPrice(t *testing.T) {

	srv, err := types.NewService()

	ok(t, err)

	price, err := types.NewPrice()

	ok(t, err)

	err = srv.AddPrice(*price)

	ok(t, err)
}

func TestRemovePrice(t *testing.T) {

	srv, err := types.NewService()

	ok(t, err)

	price, err := types.NewPrice()

	ok(t, err)

	err = srv.RemovePrice(*price)

	ok(t, err)

	prices := srv.Prices()

	assert.Equal(t, 0, len(prices), "Should be 0 prices")
}

func TestPrices(t *testing.T) {

	srv, err := types.NewService()

	ok(t, err)

	price, err := types.NewPrice()

	ok(t, err)

	err = srv.AddPrice(*price)

	ok(t, err)

	prices := srv.Prices()

	assert.Equal(t, 1, len(prices), "Should be only 1 price")
}

func ok(t *testing.T, err error) {

	assert.Nil(t, err, "Error should be nil")
}