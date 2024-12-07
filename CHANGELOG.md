# CHANGELOG of FORK

## v0.2.0
- gen/peg: more syntax support to obey `https://thrift.apache.org/docs/idl`.
- peg: support constant annotation.
- bugfix/peg: some minor syntax define bug.
- test: more test case for new syntax support.

## v0.1.1
- build: upgrade go to 1.18 (to support keyword `any`).

## v0.1.0
- bugfix/peg: BaseTypeName mismatch by prefix (e.g. binary_data_name will be matched as `binary`).
- gen/peg: re-gen via `pigeon`.

## Before
- fork from `https://github.com/samuel/go-thrift` @ `7b67f98e972fff643decea89f37246f661635cdb`.