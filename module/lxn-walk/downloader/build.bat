rem @echo off
rsrc -manifest  main.manifest -o rsrc.syso
rem go build -ldflags "-H windowsgui -s -w"
go build
del rsrc.syso