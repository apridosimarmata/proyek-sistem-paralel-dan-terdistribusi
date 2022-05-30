const { Client } = require('whatsapp-web.js');
const qrcode = require('qrcode-terminal');

var client;

module.exports = {
    initializeWhatsappWeb() {
        client = new Client({
            restartOnAuthFail: true,
            puppeteer: {
              headless: true,
              args: [
                '--no-sandbox',
                '--disable-setuid-sandbox',
                '--disable-dev-shm-usage',
                '--disable-accelerated-2d-canvas',
                '--no-first-run',
                '--no-zygote',
                '--single-process',
                '--disable-gpu'
              ],
            }});
          
          client.on('qr', (qr) => {
              qrcode.generate(qr, {small: true});
          });
          
          client.on('ready', () => {
              console.log('Client is ready!');
          });
          
          client.on('message', msg => {
              if (msg.body == '!ping') {
                  msg.reply('pong');
              }
          });
          
          client.initialize();
    },

    sendUserVerificationCode(destination, code){
        client.sendMessage(destination + "@c.us", "*" + code + "* adalah kode verifikasi ReviewJujur Anda.")
    }
}