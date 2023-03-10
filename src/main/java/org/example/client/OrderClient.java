package org.example.client;

import io.grpc.Channel;
import org.example.order.OrderServiceGrpc;

public class OrderClient {

    private OrderServiceGrpc.OrderServiceBlockingStub orderServiceBlockingStub;

    public OrderClient(Channel channel){
        orderServiceBlockingStub = OrderServiceGrpc.newBlockingStub(channel);
    }
}
