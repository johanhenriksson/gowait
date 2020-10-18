#!/bin/bash

if [[ $* == "--build" ]]; then
    ./build.sh
fi

docker run --rm \
    -v $(pwd):/app \
    --network cowait \
    -e COWAIT_GZIP="1" \
    -e COWAIT_TASK="H4sIAOcsjF8C/12P3U4DIRCFX6XhVuvCarUlMX2J3jfYTrpYmNnAsFSbvruwq64x4WJ+zpnzcRUW+8RR6MVVdOAcleqxVbf7hTgEMAzHveEyE61s5VKVt94ppVdPWq0eNvK5Xb/cSamlFMVhvTlBFb9TZ7ADDPYcI2Fzomwsj5Jj3fMyX/xbHC6fddabAFhDMDlXewoTUaXw4Cl8zEs0fowYrx36tHfW2z/m1Ecu4L5qctRNU5CQmxy3TGfA12qjjBBmzjoKlBh+Qwdyyc9tiZkDJqD/sZEpTJ//pmbzUwMOY3n7AmkzE/ZtAQAA" \
    johanhenriksson/gowait
