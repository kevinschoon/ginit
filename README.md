# ginit 
<img width="400px" src="https://raw.githubusercontent.com/mesanine/ginit/master/assets/ginit.png" alt="fit"/>

[![CircleCI](https://img.shields.io/circleci/project/github/mesanine/ginit.svg)](https://circleci.com/gh/mesanine/ginit) [![Godoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://godoc.org/github.com/mesanine/ginit)

`ginit` is a high(er) level library for initializing Linux. The goal of ginit is to make it possible to launch PID 1 processes in Go and initialize a Linux OS without writing shell scripts or relying on existing `/sbin/init` implementations. Ginit is intended for consumption by [mesanine](https://github.com/mesanine/mesanine) and also [Linuxkit](https://github.com/linuxkit/linuxkit). This library will be majorly refactored several times before becoming stable!
