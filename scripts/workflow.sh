#!/bin/bash
set -e

pushd `dirname $0` > /dev/null
SCRIPT_PATH=`pwd -P`
popd > /dev/null
SCRIPT_FILE=`basename $0`

COLOR_INFO='\033[0;36m'
COLOR_NONE='\033[0m'

PROBLEM_DIR="./algorithms/golang"


source ${SCRIPT_PATH}/lib/query_problem.sh

function usage()
{

    echo -e "Usage: ${0} [url]"
    echo -e ""
    echo -e "Example:"
    echo -e ""
    echo -e "   Running workflow for a problem"
    echo -e "   ${0} https://leetcode.com/problems/largest-number/"
    echo -e ""
}

function git_commit(){
    TITLE=$1
    FILE1=$2
    FILE2=$3
    git commit -m "New Problem Solution - \"${TITLE}\""  "${FILE1}" "${FILE2}"
}

if [ $# -lt 1 ] || [[ "${1}" != ${LEETCODE_NEW_URL}* ]] && [[ "${1}" != ${LEETCODE_OLD_URL}* ]]; then
    usage
    exit 255
fi

if [[ "${1}" == ${LEETCODE_OLD_URL}* ]]; then
    LEETCODE_URL=${LEETCODE_OLD_URL}
fi

leetcode_url=$1

get_question_slug ${leetcode_url}
dir_name=`echo ${QUESTION_TITLE_SLUG} | awk -F '-' '{for (i=1; i<=NF; i++) printf("%s", toupper(substr($i,1,1)) substr($i,2)) }'`
dir_name=`echo ${dir_name:0:1} | tr '[A-Z]' '[a-z]'`${dir_name:1}
source_file=`echo ${QUESTION_TITLE_SLUG} | awk -F '-' '{for (i=1; i<=NF; i++) printf("%s", toupper(substr($i,1,1)) substr($i,2)) }'`.go

#mkdir -p ${dir_name}
mkdir -p "${PROBLEM_DIR}"
mkdir -p "${PROBLEM_DIR}/${dir_name}"
echo "Step 1 : Created \"${PROBLEM_DIR}/${dir_name}\" directory!"
cd "${PROBLEM_DIR}/${dir_name}"

if [ -s "${source_file}" ]; then
    file=${source_file}
else
    file=`${SCRIPT_PATH}/comments.sh ${leetcode_url} | grep updated | awk '{print $1}'`
fi
WORKING_DIR=`pwd`
SRC="${dir_name}/${file}"
SRC_FILE="${WORKING_DIR}/${file}"
README_FILE="${SCRIPT_PATH}/../README.md"

echo "Step 2 : Created \"${SRC}\" source file!"

readme=`${SCRIPT_PATH}/readme.sh ${file}`
readme=`echo "${readme}" | head -n 1`

QUESTION_FRONTEND_ID=`echo "${readme}" | awk -F '|' '{print $2}'`
QUESTION_DIFFICULTY=`echo "${readme}" | awk -F '|' '{print $5}'`
QUESTION_TITLE=`echo "${readme}" | awk -F '|' '{print $3}' | sed 's/\[/\]/' |awk -F ']' '{print $2}'`
README_TMP=`mktemp "${README_FILE}.tmp.XXXXXX"`

awk -v row="${readme}" -v problem_id="${QUESTION_FRONTEND_ID}" '
    /^\|[0-9]+\|/ {
        current_id = $0
        sub(/^\|/, "", current_id)
        sub(/\|.*/, "", current_id)

        if (!inserted && current_id + 0 >= problem_id + 0) {
            print row
            inserted = 1
        }
        if (current_id + 0 == problem_id + 0) {
            next
        }
    }
    { print }
    END {
        if (!inserted) {
            print row
        }
    }
' "${README_FILE}" > "${README_TMP}"
cat "${README_TMP}" > "${README_FILE}"
rm -f "${README_TMP}"

PROBLEM_COUNT=`grep -Ec '^\|[0-9]+\|' "${README_FILE}"`
if grep -q '^Problems solved:' "${README_FILE}"; then
    sed -i.bak "s/^Problems solved:.*$/Problems solved: **${PROBLEM_COUNT}**/" "${README_FILE}"
else
    sed -i.bak "/^### LeetCode Algorithm$/a\\
\\
Problems solved: **${PROBLEM_COUNT}**
" "${README_FILE}"
fi
rm -f "${README_FILE}.bak"

echo "Step 3 : Updated the \"README.md\"!"
commit="git commit -m \"New Problem Solution - \\\"${QUESTION_FRONTEND_ID}. ${QUESTION_TITLE}\\\"\""


#git status

commit="${commit} \"${WORKING_DIR}/${file}\" \"${SCRIPT_PATH}/../README.md\""

COMMIT_MESSAGE="New Problem Solution - ${QUESTION_FRONTEND_ID}. ${QUESTION_TITLE}"

echo "${COMMIT_MESSAGE}" | pbcopy

echo "Step 4 : Commit message copied to clipboard!"
echo ""
echo "Commit message:"
echo "  ${COMMIT_MESSAGE}"
echo ""
echo "Run when you're ready:"
echo ""
echo "  git add ${SRC_FILE}"
echo "  git add ${README_FILE}"
echo "  git commit"
echo ""
echo "Done!"