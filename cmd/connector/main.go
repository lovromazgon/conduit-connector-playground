package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
	playground "github.com/lovromazgon/conduit-connector-playground"
)

func main() {
	sdk.Serve(playground.Connector)
}
