var sock = null;
var wsuri = "ws://127.0.0.1:1234";

window.onload = function () {

    console.log("onload");
    connect()
};

function connect() {
    if (sock != null && sock instanceof WebSocket) {
        sock.close();
    }
    sock = new WebSocket(wsuri);

    sock.onopen = function () {
        console.log("connected to " + wsuri);
        $('#send-msg').show();
        $('#connect-srv').hide();
    };

    sock.onclose = function (evt) {
        console.log("connection closed (" + evt.code + ")");
        $('#send-msg').hide();
        $('#connect-srv').show();
    };

    sock.onerror = function (evt) {
        console.log('Error occured: ' + evt.data);
        $('#send-msg').hide();
        $('#connect-srv').show();
    };

    //receive msg from server
    sock.onmessage = function (evt) {
        console.log("message received: " + evt.data);
        var msg_cnt = document.getElementById('msg-cnt');
        msg_cnt.innerHTML += '<div class="col-lg-12">' + evt.data + '</div>';
        scrollchange_msg_cnt();
    }
}

//send a msg to server
function send() {
    var msg = $('#message').val();
    if (msg != '') {
        sock.send(msg);
    }
    $('#message').val('');
}