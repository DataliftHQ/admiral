#!/bin/sh
set -e
rm -rf manpages
mkdir manpages
go run main.go man | gzip -c -9 >manpages/admiral.1.gz
