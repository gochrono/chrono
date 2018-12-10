// +build mage

package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Format namespace
type Format mg.Namespace

var ldflags = strings.Join([]string{
	"-X $PACKAGE/commands.commit=$COMMIT_HASH",
	"-X $PACKAGE/commands.date=$BUILD_DATE",
	"-X $PACKAGE/commands.version=$VERSION",
}, " ")

// Default is the target used when mage is ran without specifying a target
var Default = Test
var (
	packageName  = "github.com/gochrono/chrono"
	goexe        = "go"
	pkgPrefixLen = len(packageName)
	pkgs         []string
	pkgsInit     sync.Once
)

func flagEnv() map[string]string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	currentTag, _ := sh.Output("git", "tag", "-l", "--points-at", "HEAD")
	if currentTag == "" {
		longHash, _ := sh.Output("git", "rev-parse", "HEAD")
		currentTag = "SNAPSHOT-" + longHash
	}
	return map[string]string{
		"PACKAGE":     packageName,
		"COMMIT_HASH": hash,
		"BUILD_DATE":  time.Now().Format("2006-01-02T15:04:05Z0700"),
		"VERSION":     currentTag,
	}
}

func packageList() ([]string, error) {
	var err error
	pkgsInit.Do(func() {
		var s string
		s, err = sh.Output(goexe, "list", "./...")
		if err != nil {
			return
		}
		pkgs = strings.Split(s, "\n")
		for i := range pkgs {
			pkgs[i] = "." + pkgs[i][pkgPrefixLen:]
		}
	})
	return pkgs, nil
}

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}
	os.Setenv("GO111MODULE", "on")
}

// Build creates a binary with version information
func Build() error {
	return sh.RunWith(flagEnv(), goexe, "build", "-ldflags", ldflags, packageName)
}

// Test386 runs tests on a x86 architecture
func Test386() error {
	return sh.RunWith(map[string]string{"GOARCH": "386"}, goexe, "test", "./...")
}

// Test runs all tests
func Test() error {
	return sh.Run(goexe, "test", "./...")
}

// MetaLint runs gometalinter
func MetaLint() error {
	pkgs, err := packageList()
	if err != nil {
		return err
	}
	for _, pkg := range pkgs {
		if pkg != "." {
			err = sh.Run("gometalinter", pkg, "--enable", "misspell", "--exclude", "pkg/")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Reformat runs gofmt and overwrites source files with simplified code
func (Format) Reformat() error {
	return sh.Run("gofmt", "-s", "-w", ".")
}

// Check runs gofmt linter
func (Format) Check() error {
	pkgs, err := packageList()
	if err != nil {
		return err
	}
	failed := false
	first := true
	for _, pkg := range pkgs {
		files, err := filepath.Glob(filepath.Join(pkg, "*.go"))
		if err != nil {
			return nil
		}
		for _, f := range files {
			// gofmt doesn't exit with non-zero when it finds unformatted code
			// so we have to explicitly look for output, and if we find any, we
			// should fail this target.
			s, err := sh.Output("gofmt", "-l", f)
			if err != nil {
				fmt.Printf("ERROR: running gofmt on %q: %v\n", f, err)
				failed = true
			}
			if s != "" {
				if first {
					fmt.Println("The following files are not gofmt'ed:")
					first = false
				}
				failed = true
				fmt.Println(s)
			}
		}
	}
	if failed {
		return errors.New("improperly formatted go files")
	}
	return nil
}

// Vet runs go vet on all source files
func Vet() error {
	if err := sh.Run(goexe, "vet", "./..."); err != nil {
		return fmt.Errorf("error running go vet: %v", err)
	}
	fmt.Println("No errors found from go vet")
	return nil
}

// Lint runs golint on all source files
func Lint() error {
	pkgs, err := packageList()
	if err != nil {
		return err
	}
	failed := false
	for _, pkg := range pkgs {
		// We don't actually want to fail this target if we find golint errors,
		// so we don't pass -set_exit_status, but we still print out any failures.
		if _, err := sh.Exec(nil, os.Stderr, nil, "golint", pkg); err != nil {
			fmt.Printf("ERROR: running go lint on %q: %v\n", pkg, err)
			failed = true
		}
	}
	if failed {
		return errors.New("errors running golint")
	}
	return nil
}

// Coverage generates test coverage report
func Coverage() error {
	const (
		coverAll = "coverage-all.out"
		cover    = "coverage.out"
	)
	f, err := os.Create(coverAll)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte("mode: count")); err != nil {
		return err
	}
	pkgs, err := packageList()
	if err != nil {
		return err
	}
	for _, pkg := range pkgs {
		if err := sh.Run(goexe, "test", "-coverprofile="+cover, "-covermode=count", pkg); err != nil {
			return err
		}
		b, err := ioutil.ReadFile(cover)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		idx := bytes.Index(b, []byte{'\n'})
		b = b[idx+1:]
		if _, err := f.Write(b); err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return sh.Run(goexe, "tool", "cover", "-html="+coverAll)
}

// Check runs tests and linters
func Check() {
	mg.Deps(Test386)

	mg.Deps(Format.Check, Vet)
}
