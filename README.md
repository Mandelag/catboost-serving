# How to run

1. Install Ubuntu Jammy (using WSL)

2. Install
```

apt install libunistring-dev
apt install libudev-dev
apt update

apt list upgrade

```

3.  (If using WSL) Allocate memory for compiling seastar in WSL 
```
cat <your-windows-root>/.wslconfig
[wsl2]
memory=15GB # Limits VM memory in WSL 2 to 128 GB
# https://clay-atlas.com/us/blog/2021/08/31/windows-en-wsl-2-memory/
```

4. clone seastar, follow instruction to build & install

5. Set aio to this value (or more)
```
fs.aio-max-nr = 176416
```

6. build 

```
 g++ -std=c++20 -Wall -O3 hello.cc $(pkg-config --libs --cflags --static seastar) -o hello
 ./hello
```
