go test -v || git checkout .
git add . && git commit -m "iterate `date +%Y%m%d%H%M%S`"
git push -q &
