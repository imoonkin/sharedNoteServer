#!/bin/sh
chmod 777 $(dirname -- "$0")/sharedNoteServer
cd $(dirname -- "$0")
./sharedNoteServer