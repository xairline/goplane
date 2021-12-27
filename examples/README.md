# An example project 
> This is tested on MacOS but it can compile lin/win/mac

A minimum setup on MacOS for developing native xplane plugin using go

```
(cwd=examples)
brew install FiloSottile/musl-cross/musl-cross
brew install mingw-w64
make clean 
make -j 3 all
```
3 xpl files will be created under dist folder
