#!/bin/sh
set -e

if ! which sudo >/dev/null 2>&1; then
	case "$DISTRO_NAME" in
		Debian)
			echo '# asking for root password via su'>&2
			echo su -c 'apt-get install -y -q sudo'
			;;
		*)
			printf "What's a %s?\n", "$DISTRO_NAME"
			exit 1
			;;
	esac
fi

if [ ! -f "/etc/sudoers.d/$USER" ]; then
	if groups | grep -q sudo; then
		cat <<EOF |sudo tee /etc/sudoers.d/$USER
$USER    ALL=(ALL) NOPASSWD: ALL
EOF
	else
		echo '# asking for root password via su'>&2
		echo su -c sh -c "
cat <<EOF >/etc/sudoers.d/$USER
$USER    ALL=(ALL) NOPASSWD: ALL
EOF"
	fi
fi
