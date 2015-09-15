Stormtrooper is meant to be your first boots on the ground on a new machine.
Given a configuration directory, it will construct a bootstrap binary for a
given GOOS and GOARCH.

Use
---

Running

	GOOS=$OS GOARCH=$ARCH stormtrooper [dir]

will produce a binary named `bootstrap-$GOOS_$GOARCH`, containing the
configuration directory rooted at `dir`, or the current directory if not
provided. The `GOOS` and `GOARCH` variables are optional, and control what
OS+arch the bootstrap binary is built for.

The stormtrooper binary requires the `zip` tool and a `go` toolchain be
installed on the system. The produced binary has no dependencies.

Configuration Directory
-----------------------

A configuration directory structure is arranged like so:

	 /
	[/cfg]
	 /$GOOS/
	[/$GOOS/cfg]
	[/$GOOS/bin/]
	[/$GOOS/$GOARCH/]
	 /$GOOS/[0-9][0-9]_*

The `cfg` files are files in the cfg language (see below) that allow for some
customizing of the environment before scripts are run. They are evaluated from
least specific to most specific.

If `$GOOS/$GOARCH` or `$GOOS/bin` directories exist, they are added to the front
of the path before any scripts are run.

The scripts in the `$GOOS` directory must fit the glob pattern `[0-9][0-9]_*`
and be executable. They are run in lexical order. The scripts work best if they
are idempotent, although this cannot be enforced.

cfg Language
------------

The cfg language is a series of lines looking like:

	<cmd> <param> <arg>

All 3 strings must be separated by a single space. `Arg` may be omitted, and
will be treated as the empty string. The preceding space must be present,
though.  Lines beginning with `#` are ignored.

Current cmds are:

* * *

* **env**:
  * param: Name of environment variable.
  * arg: Value of environment variable.

`Env` allows setting of environment variables. No expansion or interpolation is
done.

Example:

	env BOOTSTRAP 1
	env WITH_SPACES this will be a string with spaces

* * *

* **cmd**:
  * param: Subcommand to run.
  * arg: Comma-separated arguments to the subcommand specified by `param`.

`Cmd` allow for running some helper subcommands before the script execution
starts. If the subcommand fails, the bootstrap process is halted.

Current subcommands are:

* wget

Takes a list of HTTP URLs to fetch. Files will be saved to the same directory
scripts are executed from.

Example:

	cmd wget http://golang.org/dl/go1.5.1.tar.gz
