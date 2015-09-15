// stormtrooper -- bootstrap generator
// Copyright 2015 Henry Donnay
// Licensed under the GNU General Public License v3, see the LICENSE file.

// +build !linux

// OS is a generic hook to populate anything unique about the target OS. If
// there's not a more specific version, this is a no-op.
func OS() Env {
	return Env(map[string]string{})
}
