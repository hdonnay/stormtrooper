// stormtrooper -- bootstrap generator
// Copyright 2015 Henry Donnay
// Licensed under the GNU General Public License v3, see the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var binDeps = []string{
	"zip",
	"go",
}

//go:generate go-bindata -prefix bootstrap/ bootstrap/
//go:generate gofmt -s -w $GOFILE bindata.go

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	tgtOS := os.Getenv("GOOS")
	tgtArch := os.Getenv("GOARCH")
	if tgtOS == "" {
		tgtOS = runtime.GOOS
	}
	if tgtArch == "" {
		tgtArch = runtime.GOARCH
	}

	tmp, err := ioutil.TempDir("", "stormtrooper")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmp)

	bootstrapDir := flag.Arg(0)
	if bootstrapDir == "" {
		bootstrapDir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	bootstrapDir, err = filepath.Abs(bootstrapDir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("using bootstrap dir %q\n", bootstrapDir)

	for _, fn := range AssetNames() {
		if err := copyfile(tmp, fn); err != nil {
			log.Fatal(err)
		}
	}

	for _, d := range binDeps {
		if _, err := exec.LookPath(d); err != nil {
			log.Fatalf("Need external program %q, unable to find in $PATH.\n", d)
		}
	}

	// We need the zip command to fix the executable after the concat, so just
	// exec it now. TODO(hank) Determine how to emulate 'zip -A'.
	cmd := exec.Command("zip", "-r", filepath.Join(tmp, "assets.zip"), ".")
	cmd.Dir = bootstrapDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("go", "build", "-o", "exe")
	cmd.Dir = tmp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fmt.Sprintf("bootstrap-%s_%s", tgtOS, tgtArch))
	if err != nil {
		log.Fatal(err)
	}

	exe, err := os.Open(filepath.Join(tmp, "exe"))
	if err != nil {
		log.Fatal(err)
	}
	defer exe.Close()

	zip, err := os.Open(filepath.Join(tmp, "assets.zip"))
	if err != nil {
		log.Fatal(err)
	}
	defer zip.Close()

	if _, err := io.Copy(f, exe); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(f, zip); err != nil {
		log.Fatal(err)
	}
	if err := f.Chmod(0755); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	if err := exec.Command("zip", "-A", f.Name()).Run(); err != nil {
		log.Fatal(err)
	}

}

// copies the file 'name' out of internal pacakging and into the directory
// 'dest'
func copyfile(dest, name string) error {
	f, err := os.Create(filepath.Join(dest, name))
	if err != nil {
		return err
	}
	defer f.Close()
	dat, err := Asset(name)
	if err != nil {
		return err
	}
	if _, err := f.Write(dat); err != nil {
		return err
	}
	return nil
}
