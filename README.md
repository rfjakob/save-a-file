save-a-file benchmarks the many strategies
an editor could use to save a file.

Example output on a SSD:

ext4:
```
$ ./save-a-file ~/tmp/foo
29.195µs per save for strategy: unlink + creat
87.218µs per save for strategy: open O_TRUNC
126.471µs per save for strategy: rename
193.053µs per save for strategy: rename overwrite
6.040365ms per save for strategy: fsync + rename overwrite
```

Btrfs:
```
$ ./save-a-file /mnt/btrfs.mnt/foo
42.11µs per save for strategy: unlink + creat
21.577µs per save for strategy: open O_TRUNC
71.706µs per save for strategy: rename
124.655µs per save for strategy: rename overwrite
12.492762ms per save for strategy: fsync + rename overwrite
```
