package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Scratch creates a scratch directory in the default location and returns the path.
// Scratch creates its directories with the given name prefix.
func scratch(prefix string) string {
	log.Printf("Creating temporary directory with prefix: %s", prefix)
	name, err := ioutil.TempDir("", prefix)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Created temp directory: %s", name)
	return name
}

// Clone a git repository.
func clone(repo, dest string) error {
	log.Printf("Cloning git repo %q to destination: %v", repo, dest)
	return run("git", "clone", repo, dest)
}

// Build a project.
func build(src, dest string) error {
	log.Printf("Building hugo project %q with output %q", src, dest)
	return run("/home/private/bin/hugo", "--source", src, "--destination", dest)
}

// Run a command in the shell and print the combined output.
func run(prog string, args ...string) error {
	log.Printf("Running command %q with args: %v", prog, args)
	cmd := exec.Command(prog, args...)
	stdoutStderr, err := cmd.CombinedOutput()
	log.Printf("%s", stdoutStderr)
	return err
}

// Remove a directory path recursively. Use with care!
func remove(d string) {
	log.Printf("Removing directory: %s", d)
	err := os.RemoveAll(d)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Removed directory: %s", d)
}

// Fullpath gets the absolute path of a given directory.
func fullpath(d string) string {
	name, err := filepath.Abs(d)
	if err != nil {
		log.Fatalln(err)
	}
	return name
}

// Publish the site in the given repo path to a specified output directory.
func publish(repoPath, destPath string) error {
	var err error

	repo := fullpath(repoPath)

	// defer clean(up(abc))
	tmp := scratch(filepath.Base(repo))
	defer remove(tmp)

	err = clone(repo, tmp)
	if err != nil {
		return err
	}

	err = build(tmp, destPath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	const (
		// rely on cwd being set to "." when post-receive hook invoked
		repoPath = "."

		// NearlyFreeSpeech output directory
		destPath = "/home/public"
	)

	err := publish(repoPath, destPath)
	if err != nil {
		log.Fatalf("A fatal error occurred: %v\n", err)
	}
}
