#!/usr/bin/env bash

RUN_TEST_SOURCE=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)

TEST_RESULTS_DIR=${RUN_TEST_SOURCE}/test-reports
export TEST_RESULTS_DIR
mkdir -p "${TEST_RESULTS_DIR}"

TEST_RESULT_FILE=${TEST_RESULTS_DIR}/test-results.txt
VET_RESULT_FILE=${TEST_RESULTS_DIR}/vet-results.txt
LINT_RESULT_FILE=${TEST_RESULTS_DIR}/lint-results.txt
FMT_RESULT_FILE=${TEST_RESULTS_DIR}/fmt-results.txt
COVER_REPORT_FILE=${TEST_RESULTS_DIR}/coverage.txt
COVER_REPORT_FILE_CC=${RUN_TEST_SOURCE}/c.out

COVER_HTML_REPORT_FILE=${TEST_RESULTS_DIR}/coverage.html


test_and_return_failures()
{
    # Get lists of packages depending on their usage, excluding the vendor.
    if [ -z "$1" ]
    then
        # If a relative package is not provided, default to ./...
        IFS=" " read -r -a PKGS <<< "$(go list ./... | grep -v "vendor" | xargs)"
    else
        IFS=" " read -r -a PKGS <<< "$(go list "$1"/... | grep -v "vendor" | xargs)"
    fi

    TEST_CMD="go test -cover -coverprofile=${COVER_REPORT_FILE} -covermode=atomic -v"

    # If a package list is non-empty, execute a test command for each. The API_PKGS are executed
    # with the "-p 1" flag so they are invoked serially.
    [ "${#PKGS[@]}" -gt 0 ] && ${TEST_CMD}  -p 1 "${PKGS[@]}" >> "${TEST_RESULT_FILE}"

    grep -c "FAIL:" "${TEST_RESULT_FILE}" | xargs
}

vet_and_return_errors()
{
    if [ -z "$1" ] # if a relative package is not provided, default to ./...
    then
        VET_CMD="go vet ./..."
    else
        VET_CMD="go vet $1/..."
    fi

    ${VET_CMD} 2>&1 | \
        grep -v vendor | \
        grep -v "exit status" | \
        tee "${VET_RESULT_FILE}" | \
        wc -l | \
        xargs
}

lint_and_return_errors()
{
    go get -u golang.org/x/lint/golint

    if [ -z "$1" ] # if a relative package is not provided, default to ./...
    then
        LINT_CMD="golint ./..."
    else
        LINT_CMD="golint $1/..."
    fi

    ${LINT_CMD} 2>&1 | \
        grep -v vendor | \
        tee "${LINT_RESULT_FILE}" | \
        wc -l | \
        xargs
}

fmt_and_return_errors()
{
    FMT_CMD="gofmt -s -d ${RUN_TEST_SOURCE}"

    ${FMT_CMD} 2>&1 | \
        awk '/^diff/ {print $NF}' | \
        grep -v vendor            | \
        tee "${FMT_RESULT_FILE}"  | \
        wc -l                     | \
        xargs
}

# Create html file of test coverage
create_coverfile()
{
    if [[ -s "${COVER_REPORT_FILE}" ]]; then
        go tool cover -html="${COVER_REPORT_FILE}" -o "${COVER_HTML_REPORT_FILE}"
        echo "Report: ${COVER_HTML_REPORT_FILE}"
        cp ${COVER_REPORT_FILE} ${COVER_REPORT_FILE_CC}
    fi
}

print_vet_results()
{
    cat "${VET_RESULT_FILE}"
}

print_lint_results()
{
    cat "${LINT_RESULT_FILE}"
}

print_fmt_results()
{
    cat "${FMT_RESULT_FILE}"
}

print_test_results()
{
    cat "${TEST_RESULT_FILE}"
}

cleanup()
{
    rm -rf "${TEST_RESULTS_DIR:?Missing Result DIR}"/*
}


