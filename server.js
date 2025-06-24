const shadowsocks = require('shadowsocks-node');

const server = shadowsocks.createServer({
  password: 'your_secure_password',  // CHANGE this!
  method: 'aes-256-gcm',
  port: 8388
});

server.listen(() => {
  console.log('Shadowsocks server running on port 8388');
});
