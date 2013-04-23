Broken is a simple toolkit for testing web applications. It has several features:

Usage: `broken port`
Help: `broken [-h] [--help] [help]`

Example: Receive a 403 status code after 10 seconds:
http://localhost:[port]/?status=403&timeout=10000

Status codes:
Broken can respond with a status code of your choice. For example, to get a 404, request:
http://localhost:[port]/?status=404
Requests for status codes other than 3 digits cause an error.

Sleep:
To get a response after x milliseconds, request:
http://localhost:[port]/?sleep=x

Body:
To insert a body (b) in the response, request:
http://localhost:[port]/?body=b

Uptime:
To get broken's uptime in milliseconds, request:
http://localhost:5000/up

Error handling:
On errors, broken returns immediately with a status code of 500 and an error message.
