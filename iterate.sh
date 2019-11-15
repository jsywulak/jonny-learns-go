#!/bin/bash

if go test -v > output.txt
then
    echo "tests passed, committing"
    git add .
    if git commit -m "iterate `date +%Y%m%d%H%M%S`"
    then
        rm -f last_failure.txt
    else
        cat last_failure.txt
    fi
    # if commit succeeeds, delete output
    git push -q &
else
    # save output
    mv output.txt last_failure.txt
    echo "tests failed, reverting"
    git checkout .
fi
