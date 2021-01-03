When you run 

> go run use.go 

from this folder, you will import the remote go module located at https://github.com/hamelsmu/test_go_module

When you run 

> go get -u

This will upgrade all modules, including the remote one according to the tagged version numbers.  However, to make sure the new verion is visible:see [this SO](https://stackoverflow.com/a/63207539) post otherwhise it takes a really long time for updates in the remote repo to become visible:

> set GOPRIVATE=github.com/myuser
