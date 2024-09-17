# Go Scripts

This repo holds Go scripts and acts as a playground for using Go with Nix.

## Commands

To use the dev shell, run `nix develop`.

The main reason to run the dev shell is to generate a new `gomod2nix.toml` whenever a dependency is added to the project. This is done by running `gomod2nix generate` in a dev shell.

To build the projects, run `nix build`.

To run the test, go to the test directory and run `go test`
