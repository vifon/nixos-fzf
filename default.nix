{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoPackage rec {
  pname = "nixos-fzf";
  version = "0.9";
  goPackagePath = "github.com/vifon/nixos-fzf";
  src = builtins.filterSource
    (path: type: type != "directory" || baseNameOf path != ".git")
    ./.;
  buildInputs = with pkgs; [
    fzf less
  ];
  postPatch = with pkgs; ''
    substituteInPlace nix/attr.go \
      --replace '"fzf"' '"${fzf}/bin/fzf"' \
      --replace '"less"' '"${less}/bin/less"'
  '';
}
