#!/bin/bash
set -e

while getopts "c:o:t:m:p:d:r:" opt; do
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
    r)
      if [ ${cmd} == "tag" ]; then
        args=$(echo "${args} -r ${OPTARG}")
      fi
    ;;
  esac
done

output=$(/app/gh-pkg-cli ${args})
echo "result=${output}" >> "${GITHUB_OUTPUT}"
cat "${GITHUB_OUTPUT}"
