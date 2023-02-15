from statistics import quantiles
from thrift import Thrift
from thrift.transport import THttpClient
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

from thriftpy.demo import CartService
from thriftpy.demo.ttypes import CartItem

def main():
    uri = "thrift-demo.cart-service-port50000.checkout-system.skyramp.test/CartService"
    socket = THttpClient.THttpClient(f'http://{uri}')
    transport = TTransport.TBufferedTransport(socket)
    protocol = TBinaryProtocol.TBinaryProtocol(transport)
    transport.open()
    print("successfully connected to cart-service")
    client = CartService.Client(protocol)
    user_id = "abcde"
    quantity = 1
    cartItem = CartItem(product_id='OLJCESPC7Z', quantity=quantity)
    products = client.AddItem(user_id, cartItem )
    print(f"sucessfully added {quantity} products to cart for user {user_id}")

    user_id = "abcde"
    cart = client.GetCart(user_id)
    print(f"sucessfully called GetCart user {user_id}. result {cart}")

    transport.close()

if __name__ == '__main__':
    try:
        main()
    except Thrift.TException as tx:
        print('%s' % tx)
