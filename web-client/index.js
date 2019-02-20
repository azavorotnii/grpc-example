let service = require("./example.proto3_grpc_web_pb");

let client = new service.CalculatorClient("http://localhost:8081");

let args = new service.ComplexArgs();
let arg1 = new service.Complex();
arg1.setReal(10.0);
let arg2 = new service.Complex();
arg2.setReal(13.0);
args.setArgList([arg1, arg2]);
client.add(args, {}, function (err, result) {
    console.log(err, result.getReal(), result.getImag());
});
console.log(client);