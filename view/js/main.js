var ai = {
  textBox: null,
  inputBox: null,
  init: function() {
    ai.textBox = document.getElementById("ai");
    ai.textBox.style.height = window.innerHeight - 175 + "px";
    ai.inputBox = document.getElementById("input");
  },
  keydown: function(evt) {
    switch(evt.keyCode) {
      case 13:
        var message = ai.send();
        ai.respond(message);
        break;
    }

  },
  printToScreen: function(message) {
    var textNode = document.createTextNode(message);
    ai.textBox.appendChild(textNode);
    ai.textBox.appendChild(document.createElement("br"));
    ai.inputBox.value = "";
    return message;
  },
  send: function() {
    ai.printToScreen(ai.inputBox.value);
  },
  respond: function(message) {
    ai.printToScreen("Hi!");
  }
};

window.onload = ai.init;
window.onkeydown = ai.keydown;
