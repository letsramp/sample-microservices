//
// Autogenerated by Thrift Compiler (0.16.0)
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//
"use strict";

var thrift = require('thrift');
var Thrift = thrift.Thrift;
var Q = thrift.Q;
var Int64 = require('node-int64');


var ttypes = require('./demo_types');
//HELPER FUNCTIONS AND STRUCTURES

var CurrencyService_GetSupportedCurrencies_args = function(args) {
};
CurrencyService_GetSupportedCurrencies_args.prototype = {};
CurrencyService_GetSupportedCurrencies_args.prototype.read = function(input) {
  input.readStructBegin();
  while (true) {
    var ret = input.readFieldBegin();
    var ftype = ret.ftype;
    if (ftype == Thrift.Type.STOP) {
      break;
    }
    input.skip(ftype);
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

CurrencyService_GetSupportedCurrencies_args.prototype.write = function(output) {
  output.writeStructBegin('CurrencyService_GetSupportedCurrencies_args');
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

var CurrencyService_GetSupportedCurrencies_result = function(args) {
  this.success = null;
  if (args) {
    if (args.success !== undefined && args.success !== null) {
      this.success = Thrift.copyList(args.success, [null]);
    }
  }
};
CurrencyService_GetSupportedCurrencies_result.prototype = {};
CurrencyService_GetSupportedCurrencies_result.prototype.read = function(input) {
  input.readStructBegin();
  while (true) {
    var ret = input.readFieldBegin();
    var ftype = ret.ftype;
    var fid = ret.fid;
    if (ftype == Thrift.Type.STOP) {
      break;
    }
    switch (fid) {
      case 0:
      if (ftype == Thrift.Type.LIST) {
        this.success = [];
        var _rtmp346 = input.readListBegin();
        var _size45 = _rtmp346.size || 0;
        for (var _i47 = 0; _i47 < _size45; ++_i47) {
          var elem48 = null;
          elem48 = input.readString();
          this.success.push(elem48);
        }
        input.readListEnd();
      } else {
        input.skip(ftype);
      }
      break;
      case 0:
        input.skip(ftype);
        break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

CurrencyService_GetSupportedCurrencies_result.prototype.write = function(output) {
  output.writeStructBegin('CurrencyService_GetSupportedCurrencies_result');
  if (this.success !== null && this.success !== undefined) {
    output.writeFieldBegin('success', Thrift.Type.LIST, 0);
    output.writeListBegin(Thrift.Type.STRING, this.success.length);
    for (var iter49 in this.success) {
      if (this.success.hasOwnProperty(iter49)) {
        iter49 = this.success[iter49];
        output.writeString(iter49);
      }
    }
    output.writeListEnd();
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

var CurrencyService_Convert_args = function(args) {
  this.from_curr = null;
  this.to_curr = null;
  if (args) {
    if (args.from_curr !== undefined && args.from_curr !== null) {
      this.from_curr = new ttypes.Money(args.from_curr);
    }
    if (args.to_curr !== undefined && args.to_curr !== null) {
      this.to_curr = args.to_curr;
    }
  }
};
CurrencyService_Convert_args.prototype = {};
CurrencyService_Convert_args.prototype.read = function(input) {
  input.readStructBegin();
  while (true) {
    var ret = input.readFieldBegin();
    var ftype = ret.ftype;
    var fid = ret.fid;
    if (ftype == Thrift.Type.STOP) {
      break;
    }
    switch (fid) {
      case 1:
      if (ftype == Thrift.Type.STRUCT) {
        this.from_curr = new ttypes.Money();
        this.from_curr.read(input);
      } else {
        input.skip(ftype);
      }
      break;
      case 2:
      if (ftype == Thrift.Type.STRING) {
        this.to_curr = input.readString();
      } else {
        input.skip(ftype);
      }
      break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

CurrencyService_Convert_args.prototype.write = function(output) {
  output.writeStructBegin('CurrencyService_Convert_args');
  if (this.from_curr !== null && this.from_curr !== undefined) {
    output.writeFieldBegin('from_curr', Thrift.Type.STRUCT, 1);
    this.from_curr.write(output);
    output.writeFieldEnd();
  }
  if (this.to_curr !== null && this.to_curr !== undefined) {
    output.writeFieldBegin('to_curr', Thrift.Type.STRING, 2);
    output.writeString(this.to_curr);
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

var CurrencyService_Convert_result = function(args) {
  this.success = null;
  if (args) {
    if (args.success !== undefined && args.success !== null) {
      this.success = new ttypes.Money(args.success);
    }
  }
};
CurrencyService_Convert_result.prototype = {};
CurrencyService_Convert_result.prototype.read = function(input) {
  input.readStructBegin();
  while (true) {
    var ret = input.readFieldBegin();
    var ftype = ret.ftype;
    var fid = ret.fid;
    if (ftype == Thrift.Type.STOP) {
      break;
    }
    switch (fid) {
      case 0:
      if (ftype == Thrift.Type.STRUCT) {
        this.success = new ttypes.Money();
        this.success.read(input);
      } else {
        input.skip(ftype);
      }
      break;
      case 0:
        input.skip(ftype);
        break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

CurrencyService_Convert_result.prototype.write = function(output) {
  output.writeStructBegin('CurrencyService_Convert_result');
  if (this.success !== null && this.success !== undefined) {
    output.writeFieldBegin('success', Thrift.Type.STRUCT, 0);
    this.success.write(output);
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

var CurrencyServiceClient = exports.Client = function(output, pClass) {
  this.output = output;
  this.pClass = pClass;
  this._seqid = 0;
  this._reqs = {};
};
CurrencyServiceClient.prototype = {};
CurrencyServiceClient.prototype.seqid = function() { return this._seqid; };
CurrencyServiceClient.prototype.new_seqid = function() { return this._seqid += 1; };

CurrencyServiceClient.prototype.GetSupportedCurrencies = function(callback) {
  this._seqid = this.new_seqid();
  if (callback === undefined) {
    var _defer = Q.defer();
    this._reqs[this.seqid()] = function(error, result) {
      if (error) {
        _defer.reject(error);
      } else {
        _defer.resolve(result);
      }
    };
    this.send_GetSupportedCurrencies();
    return _defer.promise;
  } else {
    this._reqs[this.seqid()] = callback;
    this.send_GetSupportedCurrencies();
  }
};

CurrencyServiceClient.prototype.send_GetSupportedCurrencies = function() {
  var output = new this.pClass(this.output);
  var args = new CurrencyService_GetSupportedCurrencies_args();
  try {
    output.writeMessageBegin('GetSupportedCurrencies', Thrift.MessageType.CALL, this.seqid());
    args.write(output);
    output.writeMessageEnd();
    return this.output.flush();
  }
  catch (e) {
    delete this._reqs[this.seqid()];
    if (typeof output.reset === 'function') {
      output.reset();
    }
    throw e;
  }
};

CurrencyServiceClient.prototype.recv_GetSupportedCurrencies = function(input,mtype,rseqid) {
  var callback = this._reqs[rseqid] || function() {};
  delete this._reqs[rseqid];
  if (mtype == Thrift.MessageType.EXCEPTION) {
    var x = new Thrift.TApplicationException();
    x.read(input);
    input.readMessageEnd();
    return callback(x);
  }
  var result = new CurrencyService_GetSupportedCurrencies_result();
  result.read(input);
  input.readMessageEnd();

  if (null !== result.success) {
    return callback(null, result.success);
  }
  return callback('GetSupportedCurrencies failed: unknown result');
};

CurrencyServiceClient.prototype.Convert = function(from_curr, to_curr, callback) {
  this._seqid = this.new_seqid();
  if (callback === undefined) {
    var _defer = Q.defer();
    this._reqs[this.seqid()] = function(error, result) {
      if (error) {
        _defer.reject(error);
      } else {
        _defer.resolve(result);
      }
    };
    this.send_Convert(from_curr, to_curr);
    return _defer.promise;
  } else {
    this._reqs[this.seqid()] = callback;
    this.send_Convert(from_curr, to_curr);
  }
};

CurrencyServiceClient.prototype.send_Convert = function(from_curr, to_curr) {
  var output = new this.pClass(this.output);
  var params = {
    from_curr: from_curr,
    to_curr: to_curr
  };
  var args = new CurrencyService_Convert_args(params);
  try {
    output.writeMessageBegin('Convert', Thrift.MessageType.CALL, this.seqid());
    args.write(output);
    output.writeMessageEnd();
    return this.output.flush();
  }
  catch (e) {
    delete this._reqs[this.seqid()];
    if (typeof output.reset === 'function') {
      output.reset();
    }
    throw e;
  }
};

CurrencyServiceClient.prototype.recv_Convert = function(input,mtype,rseqid) {
  var callback = this._reqs[rseqid] || function() {};
  delete this._reqs[rseqid];
  if (mtype == Thrift.MessageType.EXCEPTION) {
    var x = new Thrift.TApplicationException();
    x.read(input);
    input.readMessageEnd();
    return callback(x);
  }
  var result = new CurrencyService_Convert_result();
  result.read(input);
  input.readMessageEnd();

  if (null !== result.success) {
    return callback(null, result.success);
  }
  return callback('Convert failed: unknown result');
};
var CurrencyServiceProcessor = exports.Processor = function(handler) {
  this._handler = handler;
};
CurrencyServiceProcessor.prototype.process = function(input, output) {
  var r = input.readMessageBegin();
  if (this['process_' + r.fname]) {
    return this['process_' + r.fname].call(this, r.rseqid, input, output);
  } else {
    input.skip(Thrift.Type.STRUCT);
    input.readMessageEnd();
    var x = new Thrift.TApplicationException(Thrift.TApplicationExceptionType.UNKNOWN_METHOD, 'Unknown function ' + r.fname);
    output.writeMessageBegin(r.fname, Thrift.MessageType.EXCEPTION, r.rseqid);
    x.write(output);
    output.writeMessageEnd();
    output.flush();
  }
};
CurrencyServiceProcessor.prototype.process_GetSupportedCurrencies = function(seqid, input, output) {
  var args = new CurrencyService_GetSupportedCurrencies_args();
  args.read(input);
  input.readMessageEnd();
  if (this._handler.GetSupportedCurrencies.length === 0) {
    Q.fcall(this._handler.GetSupportedCurrencies.bind(this._handler)
    ).then(function(result) {
      var result_obj = new CurrencyService_GetSupportedCurrencies_result({success: result});
      output.writeMessageBegin("GetSupportedCurrencies", Thrift.MessageType.REPLY, seqid);
      result_obj.write(output);
      output.writeMessageEnd();
      output.flush();
    }).catch(function (err) {
      var result;
      result = new Thrift.TApplicationException(Thrift.TApplicationExceptionType.UNKNOWN, err.message);
      output.writeMessageBegin("GetSupportedCurrencies", Thrift.MessageType.EXCEPTION, seqid);
      result.write(output);
      output.writeMessageEnd();
      output.flush();
    });
  } else {
    this._handler.GetSupportedCurrencies(function (err, result) {
      var result_obj;
      if ((err === null || typeof err === 'undefined')) {
        result_obj = new CurrencyService_GetSupportedCurrencies_result((err !== null || typeof err === 'undefined') ? err : {success: result});
        output.writeMessageBegin("GetSupportedCurrencies", Thrift.MessageType.REPLY, seqid);
      } else {
        result_obj = new Thrift.TApplicationException(Thrift.TApplicationExceptionType.UNKNOWN, err.message);
        output.writeMessageBegin("GetSupportedCurrencies", Thrift.MessageType.EXCEPTION, seqid);
      }
      result_obj.write(output);
      output.writeMessageEnd();
      output.flush();
    });
  }
};
CurrencyServiceProcessor.prototype.process_Convert = function(seqid, input, output) {
  var args = new CurrencyService_Convert_args();
  args.read(input);
  input.readMessageEnd();
  if (this._handler.Convert.length === 2) {
    Q.fcall(this._handler.Convert.bind(this._handler),
      args.from_curr,
      args.to_curr
    ).then(function(result) {
      var result_obj = new CurrencyService_Convert_result({success: result});
      output.writeMessageBegin("Convert", Thrift.MessageType.REPLY, seqid);
      result_obj.write(output);
      output.writeMessageEnd();
      output.flush();
    }).catch(function (err) {
      var result;
      result = new Thrift.TApplicationException(Thrift.TApplicationExceptionType.UNKNOWN, err.message);
      output.writeMessageBegin("Convert", Thrift.MessageType.EXCEPTION, seqid);
      result.write(output);
      output.writeMessageEnd();
      output.flush();
    });
  } else {
    this._handler.Convert(args.from_curr, args.to_curr, function (err, result) {
      var result_obj;
      if ((err === null || typeof err === 'undefined')) {
        result_obj = new CurrencyService_Convert_result((err !== null || typeof err === 'undefined') ? err : {success: result});
        output.writeMessageBegin("Convert", Thrift.MessageType.REPLY, seqid);
      } else {
        result_obj = new Thrift.TApplicationException(Thrift.TApplicationExceptionType.UNKNOWN, err.message);
        output.writeMessageBegin("Convert", Thrift.MessageType.EXCEPTION, seqid);
      }
      result_obj.write(output);
      output.writeMessageEnd();
      output.flush();
    });
  }
};
