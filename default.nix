{ pkgs ? import <nixpkgs> {} }:

with pkgs;
buildGoPackage rec {
  pname = "nixos-fzf";
  version = "0.9";
  src = builtins.filterSource
    (path: type: type != "directory" || baseNameOf path != ".git")
    ./.;
  goPackagePath = "github.com/vifon/nixos-fzf";
}
