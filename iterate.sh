#!/bin/bash

if go vet > output.txt 2>&1 && go test >> output.txt 2>&1 
then
    cat output.txt
    echo "tests passed, committing"
    git add .
    if git commit -m "iterate `date +%Y%m%d%H%M%S`"
    then
        # if commit succeeeds, delete output
        rm -f last_failure.txt output.txt
    else
        cat last_failure.txt 2> /dev/null
    fi
    git push -q &
else
    mv output.txt last_failure.txt
    echo "tests failed, reverting"
    git checkout .
fi
