package org.example.service;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;
import org.example.client.OrderClient;
import org.example.order.Order;
import org.example.db.UserDao;
import org.example.db.User;
import org.example.user.Gender;
import org.example.user.UserRequest;
import org.example.user.UserResponse;
import org.example.user.UserServiceGrpc;

import java.util.List;
import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;

public class UserService extends UserServiceGrpc.UserServiceImplBase {
    private Logger logger = Logger.getLogger(UserService.class.getName());
    private UserDao userDao = new UserDao();

    @Override
    public void getUserDetails(UserRequest request, StreamObserver<UserResponse> responseObserver) {
        User user = userDao.getDetails(request.getUsername());

        UserResponse.Builder userResponseBuiler =
                UserResponse.newBuilder()
                        .setId(user.getId())
                        .setUsername(user.getUsername())
                        .setName(user.getName())
                        .setAge(user.getAge())
                        .setGender(Gender.valueOf(user.getGender()));

        List<Order> orders = getOrders(userResponseBuiler);

        userResponseBuiler.setNoOfOrders(orders.size());

        UserResponse userResponse = userResponseBuiler.build();
        responseObserver.onNext(userResponse);
        responseObserver.onCompleted();
    }

    private List<Order> getOrders(UserResponse.Builder userResponseBuiler) {
        logger.info("Creating a channel and calling the Order Client");
        ManagedChannel channel = ManagedChannelBuilder.forTarget("localhost:50052")
                .usePlaintext().build();
        OrderClient orderClient = new OrderClient(channel);
        List<Order> orders = orderClient.getOrders(userResponseBuiler.getId());

        try {
            channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            logger.log(Level.SEVERE,"Channel did not shut down",e);
        }
        return orders;
    }
}
