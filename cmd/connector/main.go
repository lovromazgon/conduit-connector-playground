package main

import (
	playground "github.com/lovromazgon/conduit-connector-playground"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	sdk.Serve(playground.Connector)
}
