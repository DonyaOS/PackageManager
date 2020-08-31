#!/bin/bash

markdownlint '**/*.md' --ignore node_modules --fix
yamllint .
