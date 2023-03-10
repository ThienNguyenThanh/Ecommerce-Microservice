package org.example.service;

import io.grpc.stub.StreamObserver;
import org.example.db.UserDao;
import org.example.db.User;
import org.example.user.Gender;
import org.example.user.UserRequest;
import org.example.user.UserResponse;
import org.example.user.UserServiceGrpc;

public class UserService extends UserServiceGrpc.UserServiceImplBase {
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

        UserResponse userResponse = userResponseBuiler.build();

        responseObserver.onNext(userResponse);
        responseObserver.onCompleted();
    }
}
