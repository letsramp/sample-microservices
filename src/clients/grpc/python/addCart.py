import asyncio

from grpclib.client import Channel

# generated by protoc
from demo_pb2 import AddItemRequest, CartItem
from demo_grpc import CartServiceStub

async def main():
    async with Channel('cartservice', 7070) as channel:
        cart = CartServiceStub(channel)

        reply = await cart.AddItem(AddItemRequest(user_id='abcde', item=CartItem(product_id="OLJCESPC7Z", quantity=1)))
        print("Successfully added item to cart.")

if __name__ == '__main__':
    asyncio.run(main())