# Dotfiles Loader
This is a simple go app that grabs your dotfiles from github and installs them

It has a very simple syntax it must follow an example of this is in my [Dotfiles](https://github.com/keloran/dotfiles)

## Options
```
--- Starting Help ---
 Usage dotfiles <command> <flags>

 Commands:
     help          This message
     cli           Install CLI Apps
     gui           Install GUI Apps
     os            Install OS Settings
     dots          Install the dotfiles
     update        Run updaters

 Flags:
     -github-user | github-username <keloran>           Github username
     -github-repository | github-repo <dotfiles>        Github repository
     -force                                             Force installs
     -skip-download                                     Skip Download
     -user <keloran>                                    The username if needs sudo

*** Ending Help ***
```
