// wrap in function to avoid polluting global
// namespace & prevent pub access
(function () {
    // websocket object
    var ws = new WebSocket("ws://" + window.location.host + "/ws"),
        // DOM element objects
        chat = document.getElementById('chatpane'),
        alias = document.getElementById('name'),
        text = document.getElementById('msg');

    // websocket message listener
    ws.onmessage = function (msg) {
        // append message + timestamp to chat pane
        // new Date().toLocaleString() 
        chat.innerText += + " " + msg.data + "\n";
    };

    // message box keylistener
    text.onkeydown = function (e) {
        // if the key pressed was enter and the message is valid
        if (e.keyCode === 13 && text.value != "") {
            // send alias + message to websocket handler
            ws.send("<" + (alias.value || 'guest') + ">: " + text.value);
            // reset message box
            text.value = "";
        }
    };
})();