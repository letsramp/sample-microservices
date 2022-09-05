import os
import requests
import random
from fastapi import FastAPI
from typing import List, Optional
from fastapi import FastAPI, Query
from urllib3.exceptions import InsecureRequestWarning


restPort=os.environ.get('REST_PORT')
productcatalogservice=os.environ.get('PRODUCT_CATALOG_SERVICE_HOST')

app=FastAPI()

@app.get("/list-recommendations")
async def listRecommendations(product_id: Optional[List[str]] = Query(None)):
    cart_ids = []
    print(product_id)
    if product_id:
        for p in product_id:
            cart_ids.append(p)
    url = 'http://{server}:{port}/get-products'.format(server=productcatalogservice, port=restPort)
    r = requests.get( url, verify=False)
    if r.status_code == 200:
        product_ids =[]
        products = r.json()
        for product in products:
            product_ids.append(product['id'])
        max_responses = 5
        filtered_products = list(set(product_ids)-set(cart_ids))
        num_products = len(filtered_products)
        num_return = min(max_responses, num_products)
        # sample list of indicies to return
        indices = random.sample(range(num_products), num_return)
        # fetch product ids from indices
        prod_list = [filtered_products[i] for i in indices]
        print(prod_list)
        return prod_list
