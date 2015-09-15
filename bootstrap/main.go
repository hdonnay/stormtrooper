// stormtrooper -- bootstrap generator
// Copyright 2015 Henry Donnay
// Licensed under the GNU General Public License v3, see the LICENSE file.

package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

var (
	verbose = flag.Bool("v", false, "be chatty")

	work string

	internalCmds = map[string]func([]string){
		"wget": wget,
	}
)

const (
	pattern = "[0-9][0-9]_*"
)

type Env map[string]string

func (e Env) Merge(other map[string]string) {
	for n, v := range other {
		e[n] = v
	}
}

// Populate the Env with the process's environment
func (e Env) OS() {
	for _, v := range os.Environ() {
		x := strings.Split(v, "=")
		e[x[0]] = x[1]
	}
}

func (e Env) Slice() []string {
	ret := make([]string, len(e))
	i := 0
	for n, v := range e {
		ret[i] = fmt.Sprintf("%s=%s", n, v)
		i++
	}
	return ret
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	work = unpack()
}

func main() {
	flag.Parse()
	env := Env(make(map[string]string))
	env.OS()
	if *verbose {
		fmt.Fprintf(os.Stderr, "unpacked scripts to %q\n", work)
	}
	defer os.RemoveAll(work)

	env.Merge(cfgParse(filepath.Join(work, "cfg")))
	env.Merge(cfgParse(filepath.Join(work, runtime.GOOS, "cfg")))
	env.Merge(OS())
	for _, p := range []string{
		filepath.Join(work, runtime.GOOS, runtime.GOARCH),
		filepath.Join(work, runtime.GOOS, "bin"),
	} {
		fi, err := os.Stat(p)
		if err == nil && fi.IsDir() {
			env["PATH"] = p + string(filepath.ListSeparator) + env["PATH"]
		}
	}

	if *verbose {
		fmt.Fprintf(os.Stderr, "using env:\n")
		for _, e := range env.Slice() {
			fmt.Fprintf(os.Stderr, "\t%s\n", e)
		}
	}

	if err := rundir(filepath.Join(work, runtime.GOOS), env.Slice()); err != nil {
		log.Fatal(err)
	}
}

// unpacks the pacakged scripts and returns the directory they're in.
func unpack() string {
	tmp, err := ioutil.TempDir("", "bootstrap")
	if err != nil {
		log.Fatal(err)
	}
	// open our exe
	r, err := zip.OpenReader(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fi := f.FileInfo()

		if fi.IsDir() {
			if err := os.Mkdir(filepath.Join(tmp, f.Name), fi.Mode()); err != nil {
				log.Fatal(err)
			}
			continue
		}

		rf, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		o, err := os.Create(filepath.Join(tmp, f.Name))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := io.Copy(o, rf); err != nil {
			log.Fatal(err)
		}
		if err := o.Chmod(fi.Mode()); err != nil {
			log.Fatal(err)
		}
		rf.Close()
		o.Close()
	}

	return tmp
}

func cfgParse(file string) map[string]string {
	e := make(map[string]string)
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return e
	}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()
		if len(t) < 3 || t[0] == '#' {
			continue
		}
		sp := strings.SplitN(t, " ", 3)
		switch sp[0] {
		case "cmd":
			if *verbose {
				fmt.Fprintf(os.Stderr, "running %s(%q)\n", sp[1], strings.Split(sp[2], ","))
			}
			internalCmds[sp[1]](strings.Split(sp[2], ","))
		case "env":
			e[sp[1]] = sp[2]
			if *verbose {
				fmt.Fprintf(os.Stderr, "set var %s=%q\n", sp[1], e[sp[1]])
			}
		default:
			log.Printf("unknown pattern %q, ignoring\n", t)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return e
}

func rundir(dir string, env []string) error {
	files, err := filepath.Glob(filepath.Join(dir, pattern))
	if err != nil {
		return err
	}
	if files == nil {
		fmt.Fprintln(os.Stderr, "no files found to run.")
		return nil
	}
	sort.Strings(files)

	var cmd *exec.Cmd
	for _, f := range files {
		fi, err := os.Stat(f)
		if err != nil {
			return err
		}
		// skip if not executable
		if fi.Mode()&0111 == 0 {
			if *verbose {
				fmt.Fprintf(os.Stderr, "file %s:\tskipping, not executable: %o\n", f, fi.Mode())
			}
			continue
		}

		cmd = exec.Command(f)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Dir = dir
		cmd.Env = env
		if err := cmd.Run(); err != nil {
			return err
		}
		if *verbose {
			fmt.Fprintf(os.Stderr, "file %s:\trun OK\n", f)
		}
	}

	return nil
}
