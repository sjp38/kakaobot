kakaobot
========

This is a simple chat bot Golang implementation based on Kakao Plus-Friend
auto-reply API [1].


HOWTO
=====

First, get an Kakaotalk Yellow ID [2] from Kakao (This may consume few days)
and enable auto-reply function of it.  Then, set up Golang environment [3].
Finally, enter next commands from your shell:
```
$ go get https://github.com/sjp38/kakaobot
$ kakaobot
```

Second command will start `kakaobot` on your localhost.  It works as an HTTP
server listening on `localhost:8080/kakaobot`.  You can connect it with your
Yellow ID by setting the app URL field from the Yellow ID setup web page.


Copyright
=========

GPL v3


Author
======

`SeongJae Park <sj38.park@gmail.com>`


References
==========

[1] https://github.com/plusfriend/auto_reply
[2] https://yellowid.kakao.com/
[3] https://golang.org/doc/install
