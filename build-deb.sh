#!/bin/bash
go build flexmanager.go
mv flexmanager rigctl-gui-linux-amd64

mkdir rigctl-gui && mkdir rigctl-gui/DEBIAN
mkdir -p rigctl-gui/usr/local/bin 
mv rigctl-gui-linux-amd64 rigctl-gui/usr/local/bin/
tee rigctl-gui/DEBIAN/control << EOF
Package: rigctl-gui
Version: 1.0
Maintainer: iz4tow
Architecture: amd64
Description: rigtcl GUI
EOF
tee rigctl-gui/DEBIAN/postinst<<EOF
#!/bin/bash
chmod +x /usr/local/bin/rigctl-gui-linux-amd64
EOF
sudo chmod +x rigctl-gui/DEBIAN/postinst
dpkg-deb --build rigctl-gui
sudo rm -rf rigctl-gui


