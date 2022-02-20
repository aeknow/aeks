GOARM=7 GOARCH=arm64 GOOS=linux go build -v -o ../aeks.arm7
GOOS=freebsd GOARCH=amd64 go build -v -o ../aeks.obsd
GOARCH=amd64 GOOS=linux go build -v -o ../aeks.linux
GOOS=netbsd GOARCH=amd64 go build -v -o ../aeks.nbsd
GOOS=openbsd GOARCH=amd64 go build -v -o ../aeks.obsd
GOARCH=amd64 GOOS=windows go build -v -o ../aeks.exe
GOARCH=amd64 GOOS=darwin go build -v -o ../aeks.darwin
