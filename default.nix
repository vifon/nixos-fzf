{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoPackage rec {
  pname = "nixos-fzf";
  version = "0.9";
  goPackagePath = "gitlab.com/wsiewierski/nixos-fzf";
  src = builtins.filterSource
    (path: type: type != "directory" || baseNameOf path != ".git")
    ./.;
  buildInputs = with pkgs; [
    makeWrapper
    fzf less
  ];
  postInstall = with pkgs; ''
    wrapProgram $out/bin/${pname} \
      --prefix PATH : ${lib.makeBinPath [ fzf less ]}
  '';
}
