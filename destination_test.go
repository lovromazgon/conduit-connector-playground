package playground_test

import (
	"context"
	"testing"

	playground "github.com/lovromazgon/conduit-connector-playground"
	"github.com/matryer/is"
)

func TestTeardown_NoOpen(t *testing.T) {
	is := is.New(t)
	con := playground.NewDestination()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}
