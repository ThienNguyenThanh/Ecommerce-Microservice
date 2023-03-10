package org.example.server;

import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import org.example.service.UserService;

public class UserServer {

    private static final Logger logger = Logger.getLogger(UserServer.class.getName());
    private Server server;

    public void startServer() {
        int port = 50051;
        try {
            server = ServerBuilder.forPort(port)
                    .addService(new UserService())
                    .build()
                    .start();
            logger.info("Server start on port 50051");

            Runtime.getRuntime().addShutdownHook(new Thread() {
               @Override
                public void run(){
                   logger.info("Clean server shutdown in case JVM was shutdown");
                   try{
                       UserServer.this.stopServer();
                   }catch (InterruptedException ex){
                       logger.log(Level.SEVERE, "Server shutdown interrupted", ex);
                   }
               }
            });
        } catch (IOException e) {
            logger.log(Level.SEVERE, "Server did not start", e);
        }
    }

    public void stopServer() throws InterruptedException {
        if(server!=null){
            server.shutdown().awaitTermination(30, TimeUnit.SECONDS);
        }
    }

    public void blockUntilShutDown() throws InterruptedException {
        if(server!=null){
            server.awaitTermination();
        }
    }

    public static void main(String[] args) throws InterruptedException{
        UserServer userServer = new UserServer();
        userServer.startServer();
        userServer.blockUntilShutDown();
    }
}
