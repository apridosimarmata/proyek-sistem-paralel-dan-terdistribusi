const whatsapp = require('./whatsapp');

const grpc = require("@grpc/grpc-js");
const PROTO_PATH = "notifications.proto";
var protoLoader = require("@grpc/proto-loader");

const options = {
  longs: String,
  oneofs: true,
  keepCase: true,
  enums: String,
  defaults: true,
};

var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);
const notificationsProto = grpc.loadPackageDefinition(packageDefinition);

whatsapp.initializeWhatsappWeb()

const server = new grpc.Server()

server.addService(notificationsProto.NotificationService.service ,{
    sendUserVerificationCode : (call, callback) => {
        whatsapp.sendUserVerificationCode(call.request.whatsappNo, call.request.verificationCode)
        callback();
    },
})

server.bindAsync(
    "127.0.0.1:6005",
    grpc.ServerCredentials.createInsecure(),
    (error, port) => {
      console.log("Server running at http://127.0.0.1:6005");
      server.start();
    }
);