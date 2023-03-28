const path = require('path');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const pino = require('pino');
const redis = require('redis');
const redisClient = redis.createClient({
  
  socket: {
    port: 6379,
    host: "redis-cart"
  }
});
const { callbackify } = require("util");

const PROTO_PATH = path.join(__dirname, './proto/cart.proto');
const PORT = 3090;

const shopProto = _loadProto(PROTO_PATH).lofishop;

const logger = pino({
    name: 'cartservice-server',
    messageKey: 'message',
    formatters: {
      level (logLevelString, logLevelNum) {
        return { severity: logLevelString }
      }
    }
  });

/**
 * Private function for server ONLY
 */
function _loadProto (path) {
    const packageDefinition = protoLoader.loadSync(
        path,
        {
          keepCase: true,
          longs: String,
          enums: String,
          defaults: true,
          oneofs: true
        }
    );
    return grpc.loadPackageDefinition(packageDefinition);
}

/**
 * 
 * @param {AddItemRequest} call 
 * @param {Empty} callback 
 */
function addItem(call, callback) {
  try {
    const request = call.request;

    async function setItem() {
      await redisClient.hSet(request.user_id, request.item.product_id, request.item.quantity)
    }

    var itemsCallback = callbackify(setItem);
    itemsCallback((err) => {
      if (err) throw err;

      logger.info(`Add Item successful`);
      callback(null, {});
    })


  } catch (err) {
    logger.error(`Add Item request failed: ${err}`);
    callback(err.message);
  }
}

/**
 * 
 * @param {GetCartRequest} call 
 * @param {Cart} callback 
 */
function getCart(call, callback) {
  try {
  
    const request = call.request;
    let items = new Array();

    // Get all items in cart
    async function getItem() {
      await redisClient.hGetAll(request.user_id).then( res => {
        for(let productId in res){
          items.push({
            "product_id": productId,
            "quantity": res[productId]
          });
        }
      })
    }

    // Convert Promise Items to callback
    var cartCallback = callbackify(getItem);
    cartCallback((err) => {
      if (err) throw err;
      const result = {
        "user_id": request.user_id,
        "items": items
      }

      logger.info(`Get cart successful`);
      callback(null, result);
    })

  } catch (err) {
    logger.error(`Get cart request failed: ${err}`);
    callback(err.message);
  }
}

/**
 * 
 * @param {EmptyCartRequest} call 
 * @param {Empty} callback 
 */
function emptyCart(call, callback) {
  try {
    let request = call.request;

    async function clearCart() {
      await redisClient.hGetAll(request.user_id).then(res => {
        for(let itemId in res) {
            redisClient.hDel(request.user_id, itemId);
        }
      })
    }

    var cartCallback = callbackify(clearCart);
    cartCallback((err) => {
      if (err) throw err;

      logger.info(`Empty cart successful`);
      callback(null, {});
    })
  } catch (err) {
    logger.error(`Empty cart request failed: ${err}`);
    callback(err.message);
  }
}

/**
 * Starts an RPC server that receives requests for the
 * Cart service at the sample server port
 */
function main() {
    redisClient.connect();
    redisClient.on('error', err => console.log('Redis Client Error', err));
    redisClient.on('connect', () => console.log('Redis connected!'));

    logger.info(`Starting cart server on port ${PORT}`);
    const server = new grpc.Server();
    server.addService(shopProto.CartService.service ,{addItem, getCart, emptyCart})

    server.bindAsync(
        `[::]:${PORT}`,
        grpc.ServerCredentials.createInsecure(),
        function() {
          logger.info(`Cart Service gRPC server started on port ${PORT}`);
          server.start();
        },
       );
}

main();