#!/bin/bash

if go test > output.txt
then
    echo "tests passed, committing"
    cat output.txt
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
    cat output.txt
    mv output.txt last_failure.txt
    echo "tests failed, reverting"
    git checkout .
fi
