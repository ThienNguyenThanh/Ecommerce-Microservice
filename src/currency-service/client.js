const path = require('path');
const pino = require('pino');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = path.join(__dirname, './proto/currency.proto');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });

const PORT = 3040;
var shopProto = grpc.loadPackageDefinition(packageDefinition).lofishop;
var client = new shopProto.CurrencyService(`localhost:${PORT}`,
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

const request = {
  from: {
    currency_code: 'CHF',
    units: 300,
    nanos: 0
  },
  to_code: 'EUR'
};

function _moneyToString (m) {
  return `${m.units}.${m.nanos.toString().padStart(9,'0')} ${m.currency_code}`;
}

client.getSupportedCurrencies({}, (err, response) => {
  if (err) {
    logger.error(`Error in getSupportedCurrencies: ${err}`);
  } else {
    logger.info(`Currency codes: ${response.currency_codes}`);
  }
});

client.convert(request, (err, response) => {
  if (err) {
    logger.error(`Error in convert: ${err}`);
  } else {
    logger.info(`Convert: ${_moneyToString(request.from)} to ${_moneyToString(response)}`);
  }
});
