process.on('SIGTERM', cleanup);
process.on('SIGINT', cleanup);


const { Worker } = require('worker_threads')

process.env['REST_PORT'] = 50000;
process.env['THRIFT_PORT'] = 60000;
process.env['DISABLE_PROFILER'] = "1";
process.env['DISABLE_TRACING'] = "1";
process.env['DISABLE_DEBUGGER'] = "1";

const grpcWorker = new Worker("./server.js");
const restWorker = new Worker("./rest-server.js")
const thriftWorker = new Worker("./thrift-server.js")

function cleanup() {
  restWorker.postMessage('cleanup');
  thriftWorker.postMessage('cleanup');
  process.exit(1);
}
