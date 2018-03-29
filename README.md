# svn-info-xml-to-ps1

[![Build Status](https://travis-ci.org/tanelpuhu/svn-info-xml-to-ps1.svg?branch=master)](https://travis-ci.org/tanelpuhu/svn-info-xml-to-ps1)
[![Go Report Card](https://goreportcard.com/badge/github.com/tanelpuhu/svn-info-xml-to-ps1)](https://goreportcard.com/report/github.com/tanelpuhu/svn-info-xml-to-ps1)

Using this to get Repo, trunk/branch/tag info to PS1 at https://github.com/tanelpuhu/dotfiles/blob/master/bash/exports.symlink and seems to be bit faster then Python :)

Usage:

	svn info --xml | svn-info-xml-to-ps1
