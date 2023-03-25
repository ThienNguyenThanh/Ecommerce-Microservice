const path = require('path');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const pino = require('pino');

const PROTO_PATH = path.join(__dirname, './proto/currency.proto');
const PORT = 3040;

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
 * Helper function that gets currency data from a stored JSON file
 * Uses public data from European Central Bank
 */
function _getCurrencyData (callback) {
    const data = require('./data/currency_conversion.json');
    callback(data);
}

/**
 * Helper function that handles decimal/fractional carrying
 */
function _carry (amount) {
    const fractionSize = Math.pow(10, 9);
    amount.nanos += (amount.units % 1) * fractionSize;
    amount.units = Math.floor(amount.units) + Math.floor(amount.nanos / fractionSize);
    amount.nanos = amount.nanos % fractionSize;
    return amount;
  }

function getSupportedCurrencies(call, callback) {
    logger.info('Getting supported currencies...');
    _getCurrencyData((data) => {
      callback(null, {currency_codes: Object.keys(data)});
    });
}

function convert(call, callback) {
    try {
        _getCurrencyData((data) => {
          const request = call.request;
    
          // Convert: from_currency --> EUR
          const from = request.from;
          const euros = _carry({
            units: from.units / data[from.currency_code],
            nanos: from.nanos / data[from.currency_code]
          });
    
          euros.nanos = Math.round(euros.nanos);
    
          // Convert: EUR --> to_currency
          const result = _carry({
            units: euros.units * data[request.to_code],
            nanos: euros.nanos * data[request.to_code]
          });
    
          result.units = Math.floor(result.units);
          result.nanos = Math.floor(result.nanos);
          result.currency_code = request.to_code;
    
          logger.info(`conversion request successful`);
          callback(null, result);
        });
      } catch (err) {
        logger.error(`conversion request failed: ${err}`);
        callback(err.message);
      }
}

function main() {
    logger.info(`Starting currency server on port ${PORT}`);
    const server = new grpc.Server();
    server.addService(shopProto.CurrencyService.service ,{getSupportedCurrencies, convert})

    server.bindAsync(
        `[::]:${PORT}`,
        grpc.ServerCredentials.createInsecure(),
        function() {
          logger.info(`Currency Service gRPC server started on port ${PORT}`);
          server.start();
        },
       );
}

main();