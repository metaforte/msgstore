# msgstore

Dependencies: Gorilla Mux
~~~~~~~~~~~~~~~~~~~~~~~~~
go get github.com/gorilla/mux

Note: 
1. The max allowed length for message is 4096. Added this constraint for disallowing too big strings

Test
----

curl http://localhost:8080/messages -d "Life is beautiful"

{"id":1}

curl http://localhost:8080/messages/1 

Life is beautiful


go test results
--------------
C:\go-projects\digitalanimal>go test github.com/pikeview/msgstore
ok      github.com/pikeview/msgstore    0.264s

C:\go-projects\digitalanimal>go test github.com/pikeview/msgstore -v
=== RUN   TestStoreMsg
2016/06/09 19:42:15 Received message 1 => Hello World
--- PASS: TestStoreMsg (0.00s)
=== RUN   TestStoreTooLongMsg
error occured%!(EXTRA string=internal error>> PostMsgHandler:[Max content length 4096] err[http: request body too large])--- PASS: TestStoreTooLongMsg (0.02s)
=== RUN   TestGetMsgHandler
2016/06/09 19:42:15 key 1
2016/06/09 19:42:15 Key found [1]=>[Hello World]
--- PASS: TestGetMsgHandler (0.00s)
=== RUN   TestGetNonExistentMsgHandler
2016/06/09 19:42:15 key 31
2016/06/09 19:42:15 Key Not found 31
--- PASS: TestGetNonExistentMsgHandler (0.00s)
PASS
ok      github.com/pikeview/msgstore    0.387s

C:\go-projects\digitalanimal>
