const {parentPort, data} = require("worker_threads");
const charge = require('./charge');
var express = require('express');
var app = express();
app.use(express.json())

app.post('/charge', function(req, res){
  res.json( charge(req.body));
})

restPort   = process.env.REST_PORT;

app.listen(restPort, () => {
  console.log(`Rest server started on on port ${restPort}`)
})