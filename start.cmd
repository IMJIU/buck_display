set gopath=%cd%;e:\go
echo %gopath%
go install SimpleDb
go install DB
go install main
cd bin
main.exe