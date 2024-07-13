# https://devenv.sh/basics/
{ pkgs, lib, config, inputs, ... }:

{
  env.GREET = "htmlformat - HTML formatter";

  packages = with pkgs; [];

  scripts.hello.exec = "echo $GREET";

  enterShell = ''
    hello
  '';

  enterTest = ''
    go test
  '';

  languages.go.enable = true;
}

