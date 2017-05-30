# pun

Punycode to Ascii command line tool.

```
$ pun -h
Usage of ./pun:
  -d	Decode domain as unicode to ascii
  -f string
    	Defines filename from where to read domain
```

```
$ echo "xn--c1yn36f" | ./pun -d
2017/05/30 14:18:30 Reading from STDIN.
點看

$ echo "xn--c1yn36f" | ./pun -d > out.txt
2017/05/30 14:18:30 Reading from STDIN.

$ cat out.txt
點看

```
