
const { parentPort, data } = require("worker_threads");
var fs = require('fs');
var thrift = require('thrift');
var currencySvc = require('./thriftjs/CurrencyService.js');
var common = require('./common.js');

var currencyHandler = {
  GetSupportedCurrencies: function(result) {
    common.getCurrencyData((data => {
      data = Object.keys(data)
      result(data)
    }));
  },
  Convert: function(money, to_curr, result) {
    const call = {
      request: {
        from: money,
        to_code: to_curr,
      }
    }
    var reult = {}
    common.convert(call, (err, data) => {
      const a = err
      result(data)
    });
  }
};

var transports = {
  'buffered': thrift.TBufferedTransport,
  'framed': thrift.TFramedTransport
};

var protocols = {
  'json': thrift.TJSONProtocol,
  'binary': thrift.TBinaryProtocol,
  'compact': thrift.TCompactProtocol
};

var serverOpt = {
  transport: thrift.TBufferedTransport,
  protocol: thrift.TBinaryProtocol,
  key: fs.readFileSync('./cert/cert.pem'),
  cert: fs.readFileSync('./cert/key.pem'),
}


thriftPort = process.env.THRIFT_PORT;

var s = thrift.createServer(currencySvc, currencyHandler, serverOpt).on('error', function(error) { console.log(error); }) //.listen(thriftPort);
s.listen(thriftPort)
console.log("CurrencyService thrift server started on port: " + thriftPort);
