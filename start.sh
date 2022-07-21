#!/bin/bash

go list -f "{{.Name}}__{{.ImportPath}}__{{.Imports}}" /sdc/dev/go-module-projects/collector_appsflyer/src/downloader/... > pkg_list.txtgo
#./app
#rm pkg_list.txt
