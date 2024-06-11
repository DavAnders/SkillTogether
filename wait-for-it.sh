#!/bin/sh

TIMEOUT=15
QUIET=0

echoerr() {
    if [ "$QUIET" -ne 1 ]; then
        echo "$@" 1>&2
    fi
}

wait_for() {
    while :; do
        if nc -z "$HOST" "$PORT"; then
            return 0
        fi
        sleep 1
    done
    return 1
}

usage() {
    exitcode="$1"
    cat << USAGE >&2
Usage:
    $0 host:port [-t timeout] [-- command args]
    -q | --quiet                        Do not output any status messages
    -t TIMEOUT | --timeout=timeout      Timeout in seconds, zero for no timeout
    -- COMMAND ARGS                     Execute command with args after the test finishes
USAGE
    exit "$exitcode"
}

while [ $# -gt 0 ]; do
    case "$1" in
        *:* )
        HOST=$(printf "%s\n" "$1" | cut -d : -f 1)
        PORT=$(printf "%s\n" "$1" | cut -d : -f 2)
        shift 1
        ;;
        -q | --quiet)
        QUIET=1
        shift 1
        ;;
        -t)
        TIMEOUT="$2"
        if [ "$TIMEOUT" = "" ]; then break; fi
        shift 2
        ;;
        --timeout=*)
        TIMEOUT="${1#*=}"
        shift 1
        ;;
        --)
        shift
        break
        ;;
        --help)
        usage 0
        ;;
        *)
        echoerr "Unknown argument: $1"
        usage 1
        ;;
    esac
done

if [ "$HOST" = "" ] || [ "$PORT" = "" ]; then
    echoerr "Error: you need to provide a host and port to test."
    usage 2
fi

wait_for
RESULT=$?
if [ "$RESULT" -ne 0 ]; then
    echoerr "Error: unable to contact $HOST:$PORT"
    exit 1
fi

if [ $# -gt 0 ]; then
    exec "$@"
else
    exit 0
fi
