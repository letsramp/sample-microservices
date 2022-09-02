from main import Option
from thrift.transport.TSSLSocket import TSSLServerSocket
from thrift.transport import TSocket
from thrift.server import TServer
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thriftpy.demo import EmailService


class DummyEmailThriftService(EmailService.Iface):
    def SendOrderConfirmation(self, email, order):
        print(f"A request to send order confirmation email to {email} for order {order} has been received.")
        return

def main(opt, dummy_mode):
    if dummy_mode:
        print(f"starting emailservice with configuration {opt.port}")
        processor = EmailService.Processor(DummyEmailThriftService())

        transport, tfactory, pfactory = None, None, None
        if opt.tls:
            transport = TSSLServerSocket("0.0.0.0", opt.port, keyfile="./cert/key.pem", certfile="./cert/cert.pem")
        else:
            transport = TSocket.TServerSocket(port=opt.port)

        pfactory = TBinaryProtocol.TBinaryProtocolFactory()

        if opt.buffered:
            tfactory = TTransport.TBufferedTransportFactory()
        server = TServer.TSimpleServer(processor, transport, tfactory, pfactory)
        server.serve()
    else:
        raise Exception('non-dummy mode not implemented yet')

def startThrift(opt):
  main(opt,  dummy_mode = True)
