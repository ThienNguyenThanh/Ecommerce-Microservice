package org.example.client;

import io.grpc.Channel;
import org.example.order.Order;
import org.example.order.OrderRequest;
import org.example.order.OrderResponse;
import org.example.order.OrderServiceGrpc;

import java.util.List;
import java.util.logging.Logger;

public class OrderClient {

    private Logger logger = Logger.getLogger(OrderClient.class.getName());
    private OrderServiceGrpc.OrderServiceBlockingStub orderServiceBlockingStub;

    public OrderClient(Channel channel){
        orderServiceBlockingStub = OrderServiceGrpc.newBlockingStub(channel);
    }

    public List<Order> getOrders(int userId){
        logger.info("Order Client calling the Order Service method");
        OrderRequest orderRequest = OrderRequest.newBuilder().setUserId(userId).build();

        OrderResponse orderResponse = orderServiceBlockingStub.getOrderForUser(orderRequest);

        return orderResponse.getOrderList();
    }
}
