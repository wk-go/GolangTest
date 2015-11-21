//调整消息窗口尺寸
function resize_msg_cnt() {
    var body_height = document.body.clientHeight;
    var send_nav_height = $('#send-nav').outerHeight();
    $('#msg-containner').height(body_height - send_nav_height);
    var title_nav_height = $('h1').outerHeight(true);
    var msg_containner_height = $('#msg-containner').innerHeight();
    $('#msg-cnt').height(msg_containner_height - title_nav_height - 2);//别忘了去掉边框的两个像素
    scrollchange_msg_cnt();
}
//调整消息框滚动条
function scrollchange_msg_cnt() {
    $("#msg-cnt").scrollTop($("#msg-cnt")[0].scrollHeight);
}