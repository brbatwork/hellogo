$(function() {
  messageHandler = {
    goodbye: function() {
      $("h1").text("Connection has been closed!");
      $("ul").html("");
    },
    welcome: function(data) {
      $("h1").text(data.data);
    },
    default: function(data) {
      $("ul").append("<li><strong>" + data.type + ":</strong>" + data.data + "</li>");
    }
  };

  numberGenerator = function() {
    if (ws.readyState === ws.OPEN) {
      ws.send(JSON.stringify({type: "randomNumber", data: Math.random()}));
      setTimeout(numberGenerator, 1000);
    }
  };

  connect = function() {
    ws = new WebSocket("ws://" + window.location.host + "/ws");
    ws.onopen = function(e) {
      console.log("onopen:", arguments);
      numberGenerator();
    };

    ws.onclose = function(e) {
      messageHandler.goodbye();
    };

    ws.onmessage = function(e) {
      var d = JSON.parse(e.data);
      var f = messageHandler[d.type];
      if (f !== undefined && f !== null) {
        f(d);
      } else {
        messageHandler.default(d);
      }
    };
  };

  connect();
});
