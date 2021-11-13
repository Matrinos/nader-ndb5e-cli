package main

import (
	"flag"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestReadProduct(t *testing.T) {

	set := flag.NewFlagSet("flag", 0)
	set.String("address", "/dev/xxx", "test")
	set.Int("slave", 100, "test")

	ctx := cli.NewContext(nil, set, nil)

	readProduct(ctx)

}
