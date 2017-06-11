# oi

Status: **Done** (waiting for feedback)

[![Build Status](https://travis-ci.org/mtojek/oi.svg?branch=master)](https://travis-ci.org/mtojek/oi)

## Description

**oi**  is a command-line utility for searching plain-text data sets for next lines defined in patterns file. Its name comes from an abbreviation: *ordered insections*, which explain the way of filtering the given data set.

So far, no regular expressions are supported and the matching line must occur only once.

## Quickstart

Download, install **oi** and filter given data set:
```bash
go get github.com/mtojek/oi
cat data | oi -p patterns
```

## Filter request logs faster than *nix "grep" (11 min to <50 ms)

Default release of *nix grep does not support filtering data sets by matching only one pattern from the given pattern set. It means that if you want to filter a huge log file by providing a list of request IDs, every line of the log file will be compared with all request IDs. It's definetely a set intersection but super-inefficient.

**oi** operates differenly. Provided that a request ID appears only once in a log file and patterns file contains all request IDs to be filtered in the proper order, **the number of comparisons for an every log line is decreased to 1 operation instead of N** (number of request IDs to be filtered).

### Sample usage

The *data* file contains 50000 of request log rows and it's total size is 2MB (2138894 B):

```bash
$ cat data
BA47E06D-00B2-4320-BC20-F3DC52ED242D 1
8F99A036-C084-43EB-95B5-F8540017FA4C 2
215EDC9D-3C49-43C1-8F27-53DE904BF6ED 3
BF96EB67-E3BA-4D90-B62C-A8C86A4D3305 4
27D0BB4E-A52C-4241-BE77-540ED5DDA1DE 5
FF74AD5A-42C5-4DC3-AFA1-FDF3BD6DB14D 6
734FFB36-BEF8-4DAD-A63D-978711715D14 7
E880982D-7A30-4277-B4CE-C1DD88FE1810 8
B2CD083A-C18A-460C-8C8F-D5072CFB0515 9
06428C75-1EB7-406E-801A-4C3796A50BA3 10
...
```

Let's filter 1/3 of requests (16666 rows). The patterns file contains the following request IDs: 

```bash
$ cat patterns
215EDC9D-3C49-43C1-8F27-53DE904BF6ED
FF74AD5A-42C5-4DC3-AFA1-FDF3BD6DB14D
B2CD083A-C18A-460C-8C8F-D5072CFB0515
3A3E702E-442E-4EF8-AD3E-A31CAAFF75D9
61830C85-2054-4673-9020-9F840081059F
5361F961-5FCA-412C-A685-E2C135C90C77
5AAB80AD-5B4A-448E-8878-F63476853527
796B869A-CA5D-4483-913A-99C05EDEF1DB
15E8B340-E082-43C6-B350-E2F549D65763
9A770FF9-213F-497B-8490-0D62C442C4B3
...
```

The default *grep* usage for such data intersection is:

```bash
$ time cat data | grep -f patterns > output
real	11m5.165s
user	10m55.544s
sys	0m1.060s
```
Notice: tested on Macbook Pro 2.6 GHz i5, 8GB DDR3

The operation takes 11 min to finish and the tool performs a plenty of unnecessary comparison. 

Let's try to do the same with **oi**:

```bash
$ time cat data | oi -p patterns > output
real	0m0.048s
user	0m0.014s
sys	0m0.030s
```

The total time decreased <50 ms. Nice, let's start using **oi**!

## Contact

Please feel free to leave any comment or feedback by opening a new issue or contacting me directly via [email](mailto:marcin@tojek.pl). Thank you.

## License

MIT License, see [LICENSE](https://github.com/mtojek/greenwall/blob/master/LICENSE) file.

