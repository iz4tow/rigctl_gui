#!/bin/bash
go build flexmanager.go
mv flexmanager rigctl-gui-linux-amd64

mkdir rigctl-gui && mkdir rigctl-gui/DEBIAN
mkdir -p rigctl-gui/usr/local/bin 
mv rigctl-gui-linux-amd64 rigctl-gui/usr/local/bin/
cp flexmanager.jpg rigctl-gui/usr/local/bin/
tee rigctl-gui/DEBIAN/control << EOF
Package: rigctl-gui
Version: 1.0
Maintainer: iz4tow
Architecture: amd64
Description: rigtcl GUI
Depends: git,make,automake,libtool,g++,cmake
EOF
tee rigctl-gui/DEBIAN/postinst<<EOF
#!/bin/bash
chmod +x /usr/local/bin/rigctl-gui-linux-amd64
git clone https://github.com/Hamlib/Hamlib
cd Hamlib
./bootstrap
./configure
make
#make check
sudo make install
cd ..
tee usr/share/applications/Flexmanager.desktop<<FOF
[Desktop Entry]
Name=FlexManager
Comment=HF Comunication System
Exec='/usr/local/bin/rigctl-gui-linux-amd64'
Icon="/usr/local/bin/flexmanager.jpg"
Terminal=false
Type=Application
Categories=Utility;
FOF
EOF
sudo chmod +x rigctl-gui/DEBIAN/postinst
dpkg-deb --build rigctl-gui
sudo rm -rf rigctl-gui
