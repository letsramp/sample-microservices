import os
import ssl
import json
import urllib3
import random
import uvicorn
from fastapi import FastAPI
from typing import List, Optional, Type
from fastapi import FastAPI, Query, Depends
from pydantic import BaseModel
restPort=os.environ.get('REST_PORT')

app=FastAPI()

@app.post("/send-confirmation")
async def sendConfirmation():
    print("Sending Email confirmation")

def startRest(opt):
    if opt.tls:
        uvicorn.run(
            "email_rest_server:app",
            host="0.0.0.0",
            port=int(opt.port),
            ssl_keyfile="./cert/key.pem",
            ssl_certfile="./cert/cert.pem")
    else:
        uvicorn.run(
            "email_rest_server:app",
            host="0.0.0.0",
            port=int(opt.port))