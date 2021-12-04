# System76

Firmware Update for Custom Keymap
```shell
# clone fork of embedded controller code
git clone git@github.com:dlstadther/ec.git

# checkout custom keymap
cd ec
git fetch origin
git checkout ds

# install build dependencies
./scripts/deps.sh

# prepare firmware with custom keymap
make BOARD=system76/lemp9 KEYMAP=dillon

# flash firmware, using bugfix command
# ref: https://github.com/system76/ec/issues/185#issuecomment-846033330
PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig/ make flash_internal


# system76 generic instructions:
# https://github.com/system76/ec/blob/master/doc/keyboard-layout-customization.md
```

