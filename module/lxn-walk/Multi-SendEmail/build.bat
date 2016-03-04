@echo off
rsrc -manifest main.manifest -ico emails.ico -o rsrc.syso
go build -ldflags "-H windowsgui -s -w"
del rsrc.syso