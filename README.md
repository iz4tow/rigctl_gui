![GoCompileStatus](https://github.com/iz4tow/rigctl_gui/actions/workflows/go.yml/badge.svg)
![golang-lint](https://github.com/iz4tow/rigctl_gui/actions/workflows/golang-lint.yml/badge.svg)


# Hamlib rigctl GUI
This is a simple GUI (Graphical User Interface) for rigctl written in GO for hamlib enthusiast like me.

## Description
This program alongside with Hamlib permit to send command to your radio using CAT protocol, very helpful for hamradio amateurs.

## Compile from source
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
go get fyne.io/fyne/v2
go install fyne.io/fyne/v2/cmd/fyne@latest
go mod tidy
go build flexmanager.go
```

To compile the code in Windows:
Install Fyne for Windows following this guide: https://docs.fyne.io/started/

- Install MSYS2 with MingW-w64 [link](https://www.msys2.org/)
- Install TDM-GCC [link](https://jmeubank.github.io/tdm-gcc/download/)
- Istall cygwin [link](https://www.cygwin.com/) - Please install git as package!
- Launch "MSYS2 MinGW 64-bit"
- Execute the following commands (if asked for install options be sure to choose “all”):
```
pacman -Syu
pacman -S git mingw-w64-x86_64-toolchain
git clone https://github.com/iz4tow/rigctl_gui.git
cd rigctl_gui
"c:/Program Files/Go/bin/go.exe" get fyne.io/fyne/v2
"c:/Program Files/Go/bin/go.exe" install fyne.io/fyne/v2/cmd/fyne@latest
"c:/Program Files/Go/bin/go.exe" mod tidy
"c:/Program Files/Go/bin/go.exe" build flexmanager.go
```
Pay Attention, if you are on a Virtualbox environment with a Windows guest you must add opengl32.dll to your project folder! You can find it in release file for windows alongside rigctl-gui.exe.

## Install from release

### Linux
If you are using Ubuntu you can look at the next paragraph for .deb installation instructions.
Install Hamlib on Linux then download rigctl-gui from release.
```
git clone https://github.com/Hamlib/Hamlib
cd Hamlib
./bootstrap
./configure
make
#make check
sudo make install
cd ..

#Download rigctl-gui from release with wget as follow or download it manually
wget https://github.com/iz4tow/rigctl_gui/releases/download/v1/rigctl-gui-linux-amd64
sudo chmod +x rigctl-gui-Linux-amd64

#to execute it
./rigctl-gui-Linux-amd64
```

### Ubuntu
Simply download .deb package from relase page the install with apt.
The package will install last release of Hamlib (and all the required packages) then rigctl-gui.
```
wget https://github.com/iz4tow/rigctl_gui/releases/download/v1/rigctl-gui.deb
sudo apt update
sudo apt install -y ./rigctl-gui.deb
```

### Windows
Download and install Hamlib for Windows https://github.com/Hamlib/Hamlib/releases.
Add Hamlib bin folder to $PATH.
Download rigctl-gui.zip and unzip it on a folder, then just execute it.
