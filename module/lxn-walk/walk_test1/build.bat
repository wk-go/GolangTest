@echo off
rsrc -manifest  main.manifest -o rsrc.syso
go build -ldflags "-H windowsgui -s -w"
del rsrc.syso