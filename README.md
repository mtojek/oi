# oi

Status: **Done** (waiting for feedback)

[![Build Status](https://travis-ci.org/mtojek/oi.svg?branch=master)](https://travis-ci.org/mtojek/oi)

## Description

**oi**  is a command-line utility for searching plain-text data sets for next lines defined in patterns file. Its name comes from an abbreviation: *ordered insections*, which explain the way of filtering the given data set.

So far, no regular expressions are supported and the matching line must occur only one time.

## Quickstart

Download and install **oi**:
```bash
go get github.com/mtojek/oi
```
Filter the given data set:
```bash
cat data | oi -f pattern
```
