kakaobot
========

This is a simple chat bot based on Kakao Plus-Friend auto-reply API [1] using
Go language.

Kakao has Yellow ID [2] system, the special ID system for service providers
based on Kakaotalk.  Once you get the ID, you can create a virtual account, aka
Plus Friend, on Kakaotalk world.  Other Kakaotalk users can add your Plus
Friend as thir friend and talk to you via the Plus Friend.  Kakao also provides
API to set the Plus Friend to reply to her/his friends (your customers, maybe)
automatically.  This implementation is using the API.


HOWTO
=====

First, get an Kakaotalk Yellow ID [2] from Kakao (This may consume few days).
Then, set up Golang environment [3].
Finally, enter next commands from your shell:
```
$ go get https://github.com/sjp38/kakaobot
$ kakaobot
```

Second command will start `kakaobot` on your localhost.  It works as an HTTP
server listening on `localhost:8080/kakaobot`.  You can setup your Plus Friend
to be connected with your `kakaobot` instance via Yellow ID control web page
[4].


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

[4] http://yellowid.tistory.com/259
