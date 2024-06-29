![GoCompileStatus](https://github.com/iz4tow/rigctl_gui/actions/workflows/go.yml/badge.svg)

# Hamlib rigctl GUI
This is a simple GUI for rigctl written in GO for hamlib enthusiast like me.

## Description
This program alongside with Hamlib permit to send command to your radio using CAT protocol, very helpful for hamradio amateurs.

## Install

### Compile from source
To compile the code in Linux (Ubuntu tested):
```
#Install HAMLIB
git clone https://github.com/Hamlib/Hamlib
cd Hamlib
./bootstrap
./configure
make
#make check
sudo make install
cd ..
#INSTALL RIGCTL GUI
git clone https://github.com/iz4tow/rigctl_gui.git
cd rigctl_gui
go mod tidy
go build flexmanager.go
```

To compile the code in Windows:
Install Fyne for Windows following this guide: https://docs.fyne.io/started/
```
git clone https://github.com/iz4tow/rigctl_gui.git
cd rigctl_gui
go init flexmanager
go mod tidy
go build flexmanager.go
```

### Install from releae

#### Linux
Install Hamlib on Linux
```
git clone https://github.com/Hamlib/Hamlib
cd Hamlib
./bootstrap
./configure
make
#make check
sudo make install
cd ..
```
Download rigctl-gui from release and execute it.

#### Windows
Download and install Hamlib for Windows https://github.com/Hamlib/Hamlib/releases
Download rigctl-gui.zip and unzip it on a folder, then just execute it
