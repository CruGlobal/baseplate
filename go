#!/bin/sh

echo "Installing clean project"

mv .git/modules .                 && \
rm -rf .git/                      && \
git init                          && \
mv modules .git/                  && \
git add dev/styles/brackets.css   && \
git add dev/styles/crubrand       && \
git add .gitmodules               && \
rm go

echo "Cru project successfully installed"

git status
