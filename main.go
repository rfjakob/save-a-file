package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"
	"time"
)

const (
	repeat = 1000
)

func main() {
	flag.Parse()
	filename := flag.Args()[0]
	filenameTmp := filename + ".tmp"

	t0 := time.Now()
	var i int

	t0 = time.Now()

	t0 = time.Now()
	for i = 0; i < repeat; i++ {
		syscall.Unlink(filename)
		fd, err := syscall.Creat(filename, 0600)
		if err != nil {
			log.Panicf("Creat failed: %v", err)
		}
		if _, err = syscall.Write(fd, []byte("hello world")); err != nil {
			log.Panicf("Write failed: %v", err)
		}
		if err = syscall.Close(fd); err != nil {
			log.Panicf("Close failed: %v", err)
		}
	}
	fmt.Printf("%v per save for strategy: unlink + creat\n", time.Since(t0)/repeat)

	t0 = time.Now()
	for i = 0; i < repeat; i++ {
		fd, err := syscall.Open(filename, syscall.O_RDWR|syscall.O_CREAT|syscall.O_TRUNC, 0600)
		if err != nil {
			log.Panicf("Open failed: %v", err)
		}
		if _, err = syscall.Write(fd, []byte("hello world")); err != nil {
			log.Panicf("Write failed: %v", err)
		}
		if err = syscall.Close(fd); err != nil {
			log.Panicf("Close failed: %v", err)
		}
	}
	fmt.Printf("%v per save for strategy: open O_TRUNC\n", time.Since(t0)/repeat)

	for i = 0; i < repeat; i++ {
		syscall.Unlink(filenameTmp)
		fd, err := syscall.Creat(filenameTmp, 0600)
		if err != nil {
			log.Panicf("Creat failed: %v", err)
		}
		if _, err = syscall.Write(fd, []byte("hello world")); err != nil {
			log.Panicf("Write failed: %v", err)
		}
		if err = syscall.Close(fd); err != nil {
			log.Panicf("Close failed: %v", err)
		}
		syscall.Unlink(filename)
		err = syscall.Rename(filenameTmp, filename)
		if err != nil {
			log.Panicf("Rename failed: %v", err)
		}
	}
	fmt.Printf("%v per save for strategy: rename\n", time.Since(t0)/repeat)

	for i = 0; i < repeat; i++ {
		syscall.Unlink(filenameTmp)
		fd, err := syscall.Creat(filenameTmp, 0600)
		if err != nil {
			log.Panicf("Creat failed: %v", err)
		}
		if _, err = syscall.Write(fd, []byte("hello world")); err != nil {
			log.Panicf("Write failed: %v", err)
		}
		if err = syscall.Close(fd); err != nil {
			log.Panicf("Close failed: %v", err)
		}
		err = syscall.Rename(filenameTmp, filename)
		if err != nil {
			log.Panicf("Rename failed: %v", err)
		}
	}
	fmt.Printf("%v per save for strategy: rename overwrite\n", time.Since(t0)/repeat)

	for i = 0; i < repeat; i++ {
		syscall.Unlink(filenameTmp)
		fd, err := syscall.Creat(filenameTmp, 0600)
		if err != nil {
			log.Panicf("Creat failed: %v", err)
		}
		if _, err = syscall.Write(fd, []byte("hello world")); err != nil {
			log.Panicf("Write failed: %v", err)
		}
		if err = syscall.Fsync(fd); err != nil {
			log.Panicf("Fsync failed: %v", err)
		}
		if err = syscall.Close(fd); err != nil {
			log.Panicf("Close failed: %v", err)
		}
		err = syscall.Rename(filenameTmp, filename)
		if err != nil {
			log.Panicf("Rename failed: %v", err)
		}
	}
	fmt.Printf("%v per save for strategy: fsync + rename overwrite\n", time.Since(t0)/repeat)
}
