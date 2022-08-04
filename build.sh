echo "Building for linux"
env GOOS=linux GOARCH=amd64 go build -o pdfstitch
echo "Building for windows"
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o pdfstitch-windows.exe