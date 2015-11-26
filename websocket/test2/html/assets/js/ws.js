var sock = null;
var wsuri = "ws://127.0.0.1:1234";
var nickname = '';

window.onload = function () {
    console.log("onload");
    set_nickname();
    connect();
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
        var msg_data = JSON.parse(evt.data);
        msg_cnt.innerHTML += '<div class="col-lg-12">' + msg_data.from_user + ":" + msg_data.msg +'</div>';
        if(msg_data['_t']=='set'){
            if(msg_data['nickname'] != undefined && msg_data['nickname'] != ''){
                if(msg_data['status'] == 1) {
                    nickname = msg_data['nickname'];
                }else{
                    set_nickname('昵称用冲突,请重新设置昵称');
                }
            }
        }
        scrollchange_msg_cnt();
    }
}
function set_nickname(msg){
    if(msg == undefined){
        msg = '请设置您的昵称';
    }
    layer.prompt({
        title: msg,
        formType: 0
    }, function(nick,index){
        if(sock == null){
            connect();
        }
        if(nick != ''){
            send_msg('设置昵称','set','',{nickname:nick});
            nickname = nick
        }
        layer.close(index);
    });
}

function send_msg(msg,type,to_user,other_data){
    if(type == undefined){
        type = "normal"
    }
    if(to_user == undefined){
        to_user = ""
    }
    var msg_data = {
        status:1,
        _t:type,
        msg:msg,
        to_user:to_user,
        from_user:nickname
    };
    if( other_data != undefined){
        $.extend(msg_data, other_data);
    }
    sock.send(JSON.stringify(msg_data));
}

//send a msg to server
function send() {
    var msg = $('#message').val();
    if (msg != '') {
        send_msg(msg)
    }
    $('#message').val('');
}