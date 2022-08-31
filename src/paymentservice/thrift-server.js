const {parentPort, data} = require("worker_threads");
var fs = require('fs');
var thrift = require('thrift');
var paymentSvc = require('./thrift-nodejs/PaymentService.js');
var charge = require('./charge');

var paymentHandler = {
    Charge: function (amount, creditCard, result) {
        console.log("Received charge")
        request = { "amount" : amount, "credit_card": creditCard, };
        id = charge(request)
        result(id, null);
    }
};

var serverOpt = {
    transport: thrift.TBufferedTransport,
    protocol: thrift.TBinaryProtocol,
    key: fs.readFileSync('./cert/cert.pem'),
    cert: fs.readFileSync('./cert/key.pem'),
 }

 thriftPort   = 50000 //process.env.THRIFT_PORT;


 var s = thrift.createServer(paymentSvc, paymentHandler, serverOpt)
 .on('error', function(error) { console.log(error); });
 s.listen(thriftPort);
 console.log("Thrift Server running on port: " + thriftPort);