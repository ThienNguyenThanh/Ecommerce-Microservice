const path = require('path');
const pino = require('pino');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = path.join(__dirname, './proto/cart.proto');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });

const PORT = 3090;
var shopProto = grpc.loadPackageDefinition(packageDefinition).lofishop;
var client = new shopProto.CartService(`localhost:${PORT}`,
                                       grpc.credentials.createInsecure());

const logger = pino({
  name: 'currencyservice-client',
  messageKey: 'message',
  formatters: {
    level (logLevelString, logLevelNum) {
      return { severity: logLevelString }
    }
  }
});


function main() {

  // Test addItem function
  client.addItem({
    user_id:"thien123",
    item: {
      product_id: 'book4',
      quantity: 1
    }}, err => logger.info(err))

  // Test getCart function
  client.getCart({user_id:"thien123"}, (err, result) => {
      if(err) {
          logger.info(err)
      }else {
          logger.info(`Found cart of thien123: ${JSON.stringify(result.items)}`)
      }
  })

  // Test emptyCart function
  client.emptyCart({user_id:"thien123"}, err => logger.info(err))
}

main()