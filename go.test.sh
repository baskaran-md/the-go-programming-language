#!/usr/bin/env bash
##
# @brief Test & Validations for all Go Packages.
##

TEST_SOURCE=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)

# shellcheck source=/dev/null
source "${TEST_SOURCE}"/test_lib.sh || exit 1

SRC=github.com/the-go-programming-language

echo ""
echo "Cleanup old reports..."
cleanup

echo ""
echo "Executing tests now..."
TEST_FAILURES=$(test_and_return_failures $SRC)

print_test_results

echo ""
echo "Generating code coverage now..."
create_coverfile

echo ""
echo "Running go vet now..."
VET_ERRORS=$(vet_and_return_errors $SRC)

print_vet_results

echo ""
echo "Running go lint now..."
LINT_ERRORS=$(lint_and_return_errors $SRC)

print_lint_results

echo ""
echo "Running go fmt now..."
FMT_ERRORS=$(fmt_and_return_errors $SRC)

print_fmt_results

echo ""
echo "========================="
echo -e "Test Failures:\\t${TEST_FAILURES:-1} (0=PASS)"
echo -e "Go Vet Errors:\\t${VET_ERRORS:-1} (0=PASS)"
echo -e "Go Lint Errors:\\t${LINT_ERRORS:-1} (0=PASS)"
echo -e "Go Fmt Errors:\\t${FMT_ERRORS:-1} (0=PASS)"
echo "========================="
echo ""

if [[ ${TEST_FAILURES:-1} -ne 0 ]] || \
   [[ ${VET_ERRORS:-1}    -ne 0 ]] || \
   [[ ${LINT_ERRORS:-1}   -ne 0 ]] || \
   [[ ${FMT_ERRORS:-1}    -ne 0 ]] ;
then
    exit 1
fi

exit 0
