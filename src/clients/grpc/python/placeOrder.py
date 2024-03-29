import asyncio

from grpclib.client import Channel

# generated by protoc
from demo_pb2 import PlaceOrderRequest, Address, CreditCardInfo
from demo_grpc import CheckoutServiceStub

async def main():
    async with Channel('checkoutservice', 5050) as channel:
        checkout = CheckoutServiceStub(channel)

        reply = await checkout.PlaceOrder(PlaceOrderRequest(
            user_id="abcde",
            user_currency="USD",
            address=Address(
                street_address="1600 Amp street",
                city="Mountain View",
                state="CA",
                country="USA",
                zip_code=94043),
            email="someone@example.com",
            credit_card=CreditCardInfo(
                credit_card_number="4432-8015-6152-0454",
                credit_card_cvv=672,
                credit_card_expiration_year=2024,
                credit_card_expiration_month=1)))

        print("Successfully placed order.\n")
        print(reply)

if __name__ == '__main__':
    asyncio.run(main())
