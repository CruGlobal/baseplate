#!/bin/sh

echo "Installing clean project"

mv .git/modules .                                                                 && \
rm -rf .git/                                                                      && \
git init                                                                          && \
mv modules .git/                                                                  && \
git clone git@github.com:CruWeb/crubrand.git dev/styles/crubrand                  && \
rm -rf dev/styles/crubrand/.git/                                                  && \
git clone git@github.com:thejamesdempsey/brackets.css.git dev/styles/brackets.css && \
rm -rf dev/styles/brackets.css/.git/                                              && \
git add dev/styles/brackets.css                                                   && \
git add dev/styles/crubrand                                                       && \
git add .gitmodules                                                               && \
rm go

echo "Cru project successfully installed"

git status