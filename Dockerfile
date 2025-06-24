FROM node:16-alpine

WORKDIR /app

RUN apk add --no-cache git && \
    git clone https://github.com/Jigsaw-Code/outline-server.git && \
    cd outline-server && \
    npm install && \
    npm run build

WORKDIR /app/outline-server/build

EXPOSE 8080

CMD ["node", "src/server_manager/server_manager.js"]
