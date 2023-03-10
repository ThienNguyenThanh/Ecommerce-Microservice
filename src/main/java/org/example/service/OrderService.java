package org.example.service;

import com.google.protobuf.util.Timestamps;
import io.grpc.stub.StreamObserver;
import org.example.db.Order;
import org.example.db.OrderDao;
import org.example.order.OrderRequest;
import org.example.order.OrderResponse;
import org.example.order.OrderServiceGrpc;

import java.util.List;
import java.util.logging.Logger;
import java.util.stream.Collectors;

public class OrderService extends OrderServiceGrpc.OrderServiceImplBase {
    private Logger logger = Logger.getLogger(OrderService.class.getName());
    private OrderDao orderDao = new OrderDao();

    @Override
    public void getOrderForUser(OrderRequest request, StreamObserver<OrderResponse> responseObserver) {
        List<Order> orders = orderDao.getOrders(request.getUserId());

        logger.info("Get order from OrderDao and converting to OrderResponse proto object");
        List<org.example.order.Order> orderForUser = orders.stream().map(order -> org.example.order.Order.newBuilder()
                .setUserId(order.getUserId())
                .setOrderId(order.getOrderId())
                .setNoOfItem(order.getNoOfItems())
                .setTotalAmount(order.getTotalAmount())
                .setOrderDate(Timestamps.fromMillis(order.getOrderDate().getTime())).build())
                .collect(Collectors.toList());

        OrderResponse orderResponse = OrderResponse.newBuilder().addAllOrder(orderForUser).build();

        responseObserver.onNext(orderResponse);
        responseObserver.onCompleted();
    }
}
