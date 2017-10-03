Cukoo Filter
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/cuckoofilter/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/cuckoofilter?status.svg)](https://godoc.org/github.com/kkdai/cuckoofilter)  [![Build Status](https://travis-ci.org/kkdai/cuckoofilter.svg?branch=master)](https://travis-ci.org/kkdai/cuckoofilter)


[Cuckoo Filter(Hashing)](https://en.wikipedia.org/wiki/Cuckoo_hashing)is a practically Better Than Bloom. Implement base on [paper](http://www.cs.cmu.edu/~binfan/papers/conext14_cuckoofilter.pdf).

Install
---------------
`go get github.com/kkdai/cuckoofilter`


Usage
---------------

```go
    //Create a couting bloom filter expect size 100, false detect rate 0.01
	ck := NewCukooFilter(1000)
	ck.Insert([]byte("abc"))
	ck.Insert([]byte("bcd"))
	ck.Insert([]byte("bed"))
	
	if count := ck.Count(); count != 3 {
		fmt.Println("Couting errer:", count)
	}
```


Inspired
---------------

- [CUCKOO FILTER：设计与实现](https://coolshell.cn/articles/17225.html)
- [seiflotfy/cuckoofilter](https://github.com/seiflotfy/cuckoofilter)



License
---------------

This package is licensed under MIT license. See LICENSE for details.

