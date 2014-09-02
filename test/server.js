var http = require("http");

http.createServer(function (request, response) {
  response.writeHead(200, {'Content-Type': 'text/plain'});
  response.end('Hello World\n');
}).listen(9999, '127.0.0.1');

console.log("Listening on port", 9999);