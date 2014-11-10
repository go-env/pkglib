pkglib
======

Common base for all packages of GET.

Code tries be keep in sync with behaviour of default Go tools so it just
uses their source code. For simplify future merges gocmd-*.go files copypasted
from cmd/go sources with main() removed from them. Pkglib exports some private
names from cmd/go code.

Note: feature from 1.4 for import comments (see http://golang.org/s/go14customimport)
removed yet for build GET on 1.3.