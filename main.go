package main

import (
  "flag"
  "fmt"
  "os"
  "strings"

  "github.com/keloran/dotfilesLoader/console"
  "github.com/keloran/dotfilesLoader/dots"
)

var USERNAME = ""

func main() {
  command := ""
  if len(os.Args) >= 2 {
    for _, arg := range os.Args {
      if strings.Contains(arg, "-") {
        continue
      }
      command = arg
    }
  }

  var dotfiles, appsGUI, appsCLI, OS, all bool
  help := true

  switch command {
  case "dots":
    dotfiles = true
  case "dotfiles":
    dotfiles = true
  case "gui":
    appsGUI = true
  case "cli":
    appsCLI = true
  case "all":
    all = true
  case "help":
    help = true
  }

  // flags
  github := flag.Bool("github", false, "Grab files from Github")
  force := flag.Bool("force", false, "Force override")
  flag.StringVar(&USERNAME, "user", "", "Username to inject")
  flag.Parse()

  // dots
  dot := dots.Dots{
    Username: USERNAME,
    Force: *force,
    Github: *github,
  }
  if dotfiles {
    err := dot.Install()
    if err != nil {
      console.Error(fmt.Sprintf("dotfiles err: %+v", err))
    }
    return
  }

  // GUI
  if appsGUI {
    return
  }

  // CLI
  if appsCLI {
    return
  }

  // OS
  if OS {
    return
  }

  // All
  if all {
    return
  }

  // Help
  if help {
    return
  }
}
