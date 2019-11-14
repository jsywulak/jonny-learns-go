#!/bin/bash

if go test -v
then
    git add . && git commit -m "iterate `date +%Y%m%d%H%M%S`"
    git push -q &
else
    git checkout .
fi
