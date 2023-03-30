import logging
import os
import random
import time
from concurrent import futures

import general_pb2
import general_pb2_grpc
import grpc

logger = logging.getLogger("recommendation-service")

class RecommendationSerivce(general_pb2_grpc.RecommendationServiceServicer):
    def ListRecommendations(self, request, context):
        max_responses = 5
        # fetch list of products from product catalog stub
        cat_response = product_stub.ListProducts(general_pb2.Empty())
        product_ids = [x.id for x in cat_response.products]
        filtered_products = list(set(product_ids)-set(request.product_ids))
        num_products = len(filtered_products)
        num_return = min(max_responses, num_products)
        # sample list of indicies to return
        indices = random.sample(range(num_products), num_return)
        # fetch product ids from indices
        prod_list = [filtered_products[i] for i in indices]
        logger.info("[Recv ListRecommendations] product_ids={}".format(prod_list))
        # build and return response
        response = general_pb2.ListRecommendationsResponse()
        response.product_ids.extend(prod_list)
        return response
    

if __name__ == "__main__":
    logger.info("initializing recommendationservice")

    port = os.environ.get("PORT", "3070")
    product_addr = os.environ.get("PRODUCT_CATALOG_SERVICE_ADDR", "")
    if product_addr == "":
        raise Exception('PRODUCT_CATALOG_SERVICE_ADDR environment variable not set')
    
    channel = grpc.insecure_channel(product_addr)
    product_stub = general_pb2_grpc.ProductCatalogServiceStub(channel)

    # Create recommendation gRPC server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    # Add class to gRPC server
    service = RecommendationSerivce()
    general_pb2_grpc.add_RecommendationServiceServicer_to_server(server, service)

    # start server
    logger.info("listening on port: " + port)
    server.add_insecure_port('[::]:'+port)
    server.start()

    # keep alive
    try:
         while True:
            time.sleep(10000)
    except KeyboardInterrupt:
            server.stop(0)