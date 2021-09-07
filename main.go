package main

import "gitlab.com/wsiewierski/nixos-fzf/nix"

func main() {
	nix.RootAttrset().Browse()
}
