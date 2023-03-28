FROM node:19.8-alpine3.16 as base

FROM base as builder
WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install --production

FROM base as without-grpc-health-probe-bin
WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/node_modules ./node_modules
COPY . .

EXPOSE 3090

ENTRYPOINT [ "node", "server.js" ]

