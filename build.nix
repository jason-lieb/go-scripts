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
  script,
}:

let
  scripts = {
    hello = buildGoApplication {
      pname = "hello";
      version = "0.1";
      go = pkgs.go_1_23;
      src = ./scripts/hello;
      modules = ./scripts/hello/gomod2nix.toml;
    };

    make-imports-absolute = buildGoApplication {
      pname = "make-imports-absolute";
      version = "0.1";
      go = pkgs.go_1_23;
      src = ./scripts/make-imports-absolute;
      modules = ./scripts/make-imports-absolute/gomod2nix.toml;
    };
  };
in
scripts.${script}
