find "$(dirname $0)" -mindepth 2 -type f -iname "$(basename $0)" | xargs -I % bash %

