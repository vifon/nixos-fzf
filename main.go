package main

import (
	"log"
	"os"

	"gitlab.com/wsiewierski/nixos-fzf/nix"
)

func main() {
	if len(os.Args) == 2 {
		nix.AttrsetStartingFrom(os.Args[1]).Browse()
	} else if len(os.Args) == 1 {
		nix.RootAttrset().Browse()
	} else {
		log.Fatal("Invalid arguments.")
	}
}
