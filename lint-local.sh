#!/bin/bash

# Use this file for linting on your local machine
# markdownlint is an npm package
# yamllint is a Python package
markdownlint '**/*.md' --ignore node_modules --fix --config .github/.markdownlint.yml
yamllint -c .github/.yamllint .
