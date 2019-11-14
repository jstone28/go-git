package main

import (
	git "gopkg.in/src-d/go-git.v4"
	config "gopkg.in/src-d/go-git.v4/config"
	"log"
	"os"
)

func main() {
	gitops := "vsca"
	gitopsuri := "https://gitlab.com/jstone28/variant-search-coding-assignment.git"

	_, err := git.PlainClone("/tmp/"+gitops, false, &git.CloneOptions{
		URL:      gitopsuri,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Println(err)
	}

	r, err := git.PlainOpen("/tmp/vsca")
	t, err := r.Worktree()


	sub := &config.Submodule{
		Name: "test-name",
		Path: "test-path",
		URL: "git@gitlab.com:jstone28/golang-alpine.git",
	}

	k, err := t.AddSubmodule(sub)

	k.Init()

	f, e := k.Repository()
	if e != nil {
		log.Println(e)
	}

	log.Println(f)
	log.Println(k.Status())

	// needs to be populated
	log.Println(t.Submodules())

}
