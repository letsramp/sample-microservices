package thrift

import "testing"

type TestFeature uint8

const (
	F0 TestFeature = 1 << iota
	ProductCatalog
	Recommendation
	CartService
	Email
	Shipping
	Ad
)

type TestScope struct {
	scope TestFeature
}

func (ts TestScope) testScope(t *testing.T, f TestFeature) {
	if !ts.Has(f) {
		t.Skip("Skip test")
	}
}

func (ts TestScope) Set(f TestFeature) {
	ts.scope = ts.scope | f
}

func (ts TestScope) Clear(f TestFeature) {
	ts.scope = ts.scope &^ f
}

func (ts TestScope) Toggle(f TestFeature) {
	ts.scope = ts.scope ^ f
}

func (ts TestScope) Has(f TestFeature) bool {
	return ts.scope&f != 0
}
