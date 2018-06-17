# Diff

Just a simple diff implementation in Go. Given 2 files, print a diff that
shows lines added and removed.

## Pre-requisites

This repo uses Bazel as the build system. See this
[link](https://docs.bazel.build/versions/master/getting-started.html) for an
intro to Bazel.

## Run tests

```sh
# from the diff directory
$ bazel test :diff_test
```