@echo off
windres -o demo-res.syso demo.rc
go build -ldflags="-H windowsgui -s -w"
del demo-res.syso