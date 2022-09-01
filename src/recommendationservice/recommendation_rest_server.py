import os
import ssl
import json
import requests
import urllib3
import random
import uvicorn
from fastapi import FastAPI
from typing import List, Optional, Type
from fastapi import FastAPI, Query, Depends
from pydantic import BaseModel
from urllib3.exceptions import InsecureRequestWarning

# Suppress warning for self signed certificate.
requests.packages.urllib3.disable_warnings(category=InsecureRequestWarning)

restPort=os.environ.get('REST_PORT')
productcatalogservice=os.environ.get('PRODUCT_CATALOG_SERVICE_HOST')
print(productcatalogservice)

app=FastAPI()

@app.get("/list-recommendations")
async def listRecommendations(pid: Optional[List[str]] = Query(None)):
    cart_ids = []
    print(pid)
    if pid:
        for p in pid:
            cart_ids.append(p)
    print(cart_ids)
    url = 'http://{server}:{port}/list-products'.format(server=productcatalogservice, port=restPort)
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