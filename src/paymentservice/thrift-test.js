var thrift = require('thrift');
var paymentSvc = require('./thrift-nodejs/PaymentService.js');
var types = require('./thrift-nodejs/demo_types.js');


var amount = new types.Money();
amount.currency_code = "USD";
amount.units = 100;
amount.nanos = 20;

var cc = new types.CreditCardInfo();
cc.credit_card_number = "4432-8015-6152-0454";
cc.credit_card_cvv = 672;
cc.credit_card_expiration_year = 2024;
cc.credit_card_expiration_month= 2;

var connection = thrift.createConnection('localhost', 50000, {
    transport: thrift.TBufferedTransport,
    protocol: thrift.TBinaryProtocol
})
.on ("error", function(error) { console.log(error); } )
.on("connect", function() {
     var client = thrift.createClient(paymentSvc, connection);
     client.Charge( amount, cc, function (error, result) {
        console.log(result);
        connection.end()
    }
     )
})



req = {
	"amount": {
		"currency_code": "USD",
		"units": 100,
		"nanos": 50
	},
	"credit_card": {
		"credit_card_number": "4432-8015-6152-0454",
		"credit_card_cvv": 672,
		"credit_card_expiration_year": 2024,
		"credit_card_expiration_month": 2
	}
}

