const {parentPort, data} = require("worker_threads");
var fs = require('fs');
var http = require('http');
var https = require('https');
var privateKey  = fs.readFileSync('cert/key.pem', 'utf8');
var certificate = fs.readFileSync('cert/cert.pem', 'utf8');
const charge = require('./charge');
var credentials = {key: privateKey, cert: certificate};
var express = require('express');
var app = express();
app.use(express.json())

app.post('/charge', function(req, res){
  res.json( charge(req.body));
})

restPort   = process.env.REST_PORT;
var httpServer = http.createServer(app);
var httpsServer = https.createServer(credentials, app);
console.log("Rest Server running on port: " + restPort);
httpsServer.listen(restPort);