{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [ (import "${fetchTree gomod2nix.locked}/overlay.nix") ];
    }
  ),
  buildGoApplication ? pkgs.buildGoApplication,
}:

buildGoApplication {
  pname = "hello";
  version = "0.1";
  src = ./scripts/hello;
  modules = ./scripts/hello/gomod2nix.toml;
}
