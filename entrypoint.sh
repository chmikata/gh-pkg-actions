#!/bin/bash
set -e

while getopts "c:o:t:m:p:d:" opt; do
  case "${opt}" in
    c)
      cmd=${OPTARG}
      args=${cmd}
    ;;
    o)
      args=$(echo "${args} -o ${OPTARG}")
    ;;
    t)
      args=$(echo "${args} -t ${OPTARG}")
    ;;
    m)
      args=$(echo "${args} -m ${OPTARG}")
    ;;
    p)
      if [ ${cmd} == "tag" ]; then
        args=$(echo "${args} -p ${OPTARG}")
      fi
    ;;
    d)
      if [ ${cmd} == "tag" ]; then
        args=$(echo "${args} -d ${OPTARG}")
      fi
    ;;
  esac
done

# run command
/app/gh-pkg-cli "${args}"
