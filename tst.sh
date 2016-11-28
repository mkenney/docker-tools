#!/usr/bin/env bash

function __get_args {
    declare ret_val=
    for var in "$@"; do
        if [ "-" != "${var:0:1}" ]; then
            if [ "" == "$ret_val" ]; then
                ret_val="$var"
            else
                ret_val="$ret_val $var"
            fi
        fi
    done
    echo $ret_val | tr " " "\n" | sort | uniq | tr "\n" " "
}

function __get_opts {
    declare ret_val=
    for var in "$@"; do
        if [ "--" == "${var:0:2}" ]; then
            if [ "" == "$ret_val" ]; then
                ret_val="${var:2}"
            else
                ret_val="$ret_val ${var:2}"
            fi
        fi
    done
    echo $ret_val | tr " " "\n" | sort | uniq | tr "\n" " "
}

function __get_flags {
    declare ret_val=
    for var in "$@"; do
        if [ "-" == "${var:0:1}" ] && [ "-" != "${var:1:1}" ]  && [ "" != "${var:1}" ]; then
            if [ "" == "$ret_val" ]; then
                ret_val="${var:1}"
            else
                ret_val="$ret_val ${var:1}"
            fi
        fi
    done
    echo $ret_val | tr " " "\n" | sort | uniq | tr "\n" " "
}




echo "args: $(__get_args $@)"
echo "opts: $(__get_opts $@)"
echo "flags: $(__get_flags $@)"
exit




asdf=$(tst $@)
for a in $asdf; do
    echo ${[i]}
done

echo "array:"
asdf=(asdf jkl)
echo "$asdf"