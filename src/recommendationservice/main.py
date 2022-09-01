import os
from time import sleep
from threading import Thread
import recommendation_server as grpc
import recommendation_thrift_server as thrift
import uvicorn

os.environ["DISABLE_TRACING"] = "1"
os.environ["DISABLE_PROFILER"] = "1"

# keep environment variables PRODUCT_CATALOG_SERVICE_ADDR for upstream grcp service
productCatalogService = os.environ.get('PRODUCT_CATALOG_SERVICE_ADDR', "product-catalog-service:8080")
os.environ["PRODUCT_CATALOG_SERVICE_ADDR"] = productCatalogService

productCatalogServiceHost = productCatalogService.split(":")[0]
os.environ["PRODUCT_CATALOG_SERVICE_HOST"] = productCatalogServiceHost
os.environ["REST_PORT"] = "60000"

thriftPort = 50000
restPort = 60000
productCatalogServiceThriftPort = 50000

def runGrpc():
    print("starting recommendation-service grcp endpoint")
    grpc.startGrpc()

def runRest():
    print("starting recommendation-service rest endpoint")
    uvicorn.run(
    "recommendation_rest_server:app",
    host="0.0.0.0",
    port=int(restPort))

def runThrift():
    print("starting recommendation-service thrift endpoint")
    thrift.startThrift(thriftPort, productCatalogServiceHost, productCatalogServiceThriftPort)

if __name__ == "__main__":
    grpcThread = Thread(target=runGrpc)
    restThread = Thread(target=runRest)
    thriftThread = Thread(target=runThrift)
    grpcThread.start()
    restThread.start()
    thriftThread.start()
    grpcThread.join()
    restThread.join()
    thriftThread.join()
