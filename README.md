## Description
This program was written to be placed in the windows "Send To" menu. You can then select a group of files, right click, send to, New Folder With Items.exe, and a new folder will be created and the selected files will be moved into it. This replicates a feature in the MacOS finder called `New Folder with Selection`

## Installation 
Build the EXE and install with `go build -ldflags "-s -w -H=windowsgui" -o "%APPDATA%\Microsoft\Windows\SendTo\New Folder With Items.exe" .\NewFolderWithItems.go`
