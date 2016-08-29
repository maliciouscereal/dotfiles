package main

import (
	"fmt"
	"github.com/rhysd/dotfiles/dotfiles-command"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app = kingpin.New("dotfiles", "A dotfiles manager")

	clone      = app.Command("clone", "Clone remote repository")
	clone_repo = clone.Arg("repository", "Repository.  Format: 'user', 'user/repo-name', 'git@somewhere.com:repo.git, 'https://somewhere.com/repo.git'").Required().String()
	clone_path = clone.Arg("path", "Path where repository cloned").String()

	link        = app.Command("link", "Put symlinks to setup your configurations")
	link_dryrun = link.Flag("dry", "Show what happens only").Bool()

	list = app.Command("list", "Show a list of symbolic link put by this command")

	clean = app.Command("clean", "Remove all symbolic links put by this command")

	update = app.Command("update", "Update your dotfiles repository")

	version = app.Command("version", "Show version")
)

func unimplemented(cmd string) {
	fmt.Fprintf(os.Stderr, "Command '%s' is not implemented yet!\n", cmd)
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case clone.FullCommand():
		handleError(dotfiles.Clone(*clone_repo, *clone_path))
	case link.FullCommand():
		unimplemented("link")
	case list.FullCommand():
		unimplemented("list")
	case clean.FullCommand():
		unimplemented("clean")
	case update.FullCommand():
		unimplemented("update")
	case version.FullCommand():
		fmt.Println(dotfiles.Version())
	default:
		panic("Internal error: Unreachable! Please report this to https://github.com/rhysd/dotfiles-command/issues")
	}
}
