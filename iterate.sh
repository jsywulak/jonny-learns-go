#!/bin/bash

if go test
then
    echo "tests passed, committing"
    git add . && git commit -m "iterate `date +%Y%m%d%H%M%S`"
    git push -q &
else
    echo "tests failed, reverting"
    git checkout .
fi
