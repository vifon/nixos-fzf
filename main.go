package main

import "github.com/vifon/nixos-fzf/nix"

func main() {
	nix.RootAttrset().Browse()
}
