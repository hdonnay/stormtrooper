// stormtrooper -- bootstrap generator
// Copyright 2015 Henry Donnay
// Licensed under the GNU General Public License v3, see the LICENSE file.

// +build linux

package main

import (
	"bytes"
	"os/exec"
	"strings"
)

// OS populates two envinroment variables: DISTRO_NAME and DISTRO_VER. If
// unable to determine this information, they will be set to the empty string
// and 0, respectively.
func OS() Env {
	d := Env(map[string]string{
		"DISTRO_NAME": "",
		"DISTRO_VER":  "0",
	})
	cmd := exec.Command("lsb_release", "-i", "-r", "-s")
	out, err := cmd.Output()
	if err != nil {
		return d
	}
	r := bytes.NewBuffer(out)

	x, err := r.ReadString('\n')
	if err != nil {
		return d
	}
	d["DISTRO_NAME"] = strings.TrimSpace(x)

	x, err = r.ReadString('\n')
	if err != nil {
		return d
	}
	d["DISTRO_VER"] = strings.TrimSpace(x)

	return d
}
