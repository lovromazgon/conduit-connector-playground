package playground_test

import (
	"context"
	"testing"

	playground "github.com/lovromazgon/conduit-connector-playground"
	"github.com/matryer/is"
)

func TestTeardownSource_NoOpen(t *testing.T) {
	is := is.New(t)
	con := playground.NewSource()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}
