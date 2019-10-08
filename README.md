Build C libraries first (`gcc` and `cmake` are required):
```
./build.sh
...
-- Installing: /usr/local/lib/libcgomultiversion.so.1.0.0
-- Installing: /usr/local/lib/libcgomultiversion.so.1
-- Installing: /usr/local/lib/libcgomultiversion.so
-- Installing: /usr/local/include/cgomultiversion_v1/lib.h
...
-- Installing: /usr/local/lib/libcgomultiversion.so.2.0.0
-- Installing: /usr/local/lib/libcgomultiversion.so.2
-- Installing: /usr/local/lib/libcgomultiversion.so
-- Installing: /usr/local/include/cgomultiversion_v2/lib.h
```


Try to run tests, and you'll see the error: 

```
go test -count=1 -v        
# github.com/vitalyisaev2/cgo_multiversion_go_lib.test
/usr/local/go/pkg/tool/linux_amd64/link: running gcc failed: exit status 1
/usr/bin/ld: /tmp/go-link-218143790/000005.o: in function `add':
/home/isaev/go/src/github.com/vitalyisaev2/cgo_multiversion_go_lib/v2/lib.c:6: multiple definition of `add'; /tmp/go-link-218143790/000002.o:/home/isaev/go/src/github.com/vitalyisaev2/cgo_multiversion_go_lib/v1/lib.c:5: first defined here
collect2: error: ld returned 1 exit status

```

This code cannot be compiled with two different versions of C library because of name overlapping.
