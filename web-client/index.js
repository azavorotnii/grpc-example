let service = require("./example.proto3_grpc_web_pb");

let client = new service.CalculatorClient("http://localhost:8081");

global.add = function() {
    let args = new service.ComplexArgs();
    let arg1 = new service.Complex();
    arg1.setReal(parseInt(document.getElementById("arg1real").value));
    arg1.setImag(parseInt(document.getElementById("arg1imag").value));
    let arg2 = new service.Complex();
    arg2.setReal(parseInt(document.getElementById("arg2real").value));
    arg2.setImag(parseInt(document.getElementById("arg2imag").value));
    args.setArgList([arg1, arg2]);
    client.add(args, {}, function (err, result) {
        if (err !== null) {
            alert(err.message);
            return;
        }
        document.getElementById("result_real").value = result.getReal();
        document.getElementById("result_imag").value = result.getImag();
    });
};
