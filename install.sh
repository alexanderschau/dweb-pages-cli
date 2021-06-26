#!/bin/sh

export platform="$(uname -s)_$(uname -m)"
export dwebpages=$HOME/.dweb-pages-cli

latestRelease=$(curl -L --silent -s https://api.github.com/repos/alexanderschau/dweb-pages-cli/releases/latest \
    | grep browser_download_url \
    | grep $platform \
    | cut -d : -f 2,3)

export download=true
if command -v dweb-pages &> /dev/null; then
    installedVersion=$(dweb-pages version)

    if echo $latestRelease | grep $installedVersion >/dev/null; then
        exit 0
    fi
fi

echo "Downloading latest dweb-pages version from$latestRelease"
echo $latestRelease | xargs curl -L --output $dwebpages

chmod +x $dwebpages


INSTALL_DIR=$(dirname $0)

bin=$dwebpages #"$INSTALL_DIR/dweb-pages_0.0.1_Linux_x86_64"
binpaths="/usr/local/bin /usr/bin"

# This variable contains a nonzero length string in case the script fails
# because of missing write permissions.
is_write_perm_missing=""

for binpath in $binpaths; do
  if mv "$bin" "$binpath/dweb-pages" ; then
    echo "Moved $bin to $binpath"
    exit 0
  else
    if [ -d "$binpath" ] && [ ! -w "$binpath" ]; then
      is_write_perm_missing=1
    fi
  fi
done

echo "We cannot install $bin in one of the directories $binpaths"

if [ -n "$is_write_perm_missing" ]; then
  echo "It seems that we do not have the necessary write permissions."
  echo "Perhaps try running this script as a privileged user:"
  echo
  echo "    sudo $0"
  echo
fi

exit 1