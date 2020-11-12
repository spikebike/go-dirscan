go-dirscan is a simple parallel directory scanner, much like a parallel du.  Likely only a significant win on parallel file systems.

To run:
```
$ time go run dirscan.go .. 8
workers = 8
Total Files=1074232 Total Dirs=994209 Total Size=27519870184

real	0m7.512s
user	0m21.561s
sys	0m31.074s
```
