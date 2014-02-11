#!/bin/sh

echo "Installing clean project"

rm -rf .git/                                                                      && \
git init                                                                          && \
git clone git@github.com:CruWeb/crubrand.git dev/styles/crubrand                  && \
git clone git@github.com:thejamesdempsey/brackets.css.git dev/styles/brackets.css && \
git add dev/styles/brackets.css                                                   && \
git add dev/styles/crubrand                                                       && \
git add .gitmodules                                                               && \
rm go

echo "Cru project successfully installed"

git status