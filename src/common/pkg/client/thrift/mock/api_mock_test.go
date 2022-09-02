package thrift

import (
	"encoding/json"
	"fmt"
	"testing"

	demo "github.com/pellepedro/gcp-microservices/src/common/idl/thrift/demo"
	"github.com/stretchr/testify/assert"
)

func TestParseProductCatalog(t *testing.T) {
	products := []*demo.Product{}
	err := json.Unmarshal([]byte(jsonProducts), &products)
	assert.Nil(t, err)
	assert.Equal(t, 9, len(products))
	fmt.Println(products[0])
}
