#!/bin/bash

datadir=".training-data"
tempdir=".temp-dir"
tempfile="tempfile.zip"
# repo="https://github.com/Gabriel-Ivarsson/code2vec-demo/archive/refs/heads/main.zip"
repo=$1

if [[ $repo != *.zip ]]; then
  echo "Please provide a download link to the repo as a .zip file"
  exit 1
fi

echo "Downloading" "$repo" "and extracting .go files, please wait..."
curl -L -o "$tempfile" "$repo" # curl zip to tempfile
mkdir -p "$datadir" # create data_dir directory if it doesnt exits.
mkdir -p "$tempdir" # create temp_dir directory if it doesnt exits.
unzip "$tempfile" -d "$datadir" "*.go" -x "*_test.go" # unzip tempfile to data_dir

# clean up
rm -f "$tempfile" # delete tempfile
rm -rf "$tempdir" # delete tempdir
