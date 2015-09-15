// stormtrooper -- bootstrap generator
// Copyright 2015 Henry Donnay
// Licensed under the GNU General Public License v3, see the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

func wget(args []string) {
	for i, f := range args {
		if *verbose {
			fmt.Fprintf(os.Stderr, "attempting to fetch %q\n", f)
		}

		u, err := url.Parse(f)
		if err != nil {
			log.Fatal(err)
		}

		fn := path.Base(u.Path)
		if fn == "" {
			fn = "index.html"
		}
		res, err := http.Get(f)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			log.Fatal("got status code %d", res.StatusCode)
		}
		defer res.Body.Close()
		o, err := os.Create(filepath.Join(work, fn))
		if err != nil {
			log.Fatal(err)
		}
		defer o.Close()

		if _, err := io.Copy(o, res.Body); err != nil {
			log.Fatal(err)
		}

		if *verbose {
			fmt.Fprintf(os.Stderr, "file %02d at %s\n", i, fn)
		}
	}
}
