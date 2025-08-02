import { handleIncomingMessage } from "./handleIncomingMessage.js";

const ws = new WebSocket(
  "ws://" +
    location.host +
    "/ws/" +
    encodeURIComponent(code) +
    "?nickname=" +
    encodeURIComponent(nickname)
);

ws.onmessage = (e) => {
  console.log(e.data);
  const msg = JSON.parse(e.data);
  handleIncomingMessage(msg);
};

// Game start button
const gameStartButton = document.getElementById("game-start");
gameStartButton.onclick = () => {
  sendJSONMessage({ type: "startGame" });
};

// Action buttons
const callButton = document.getElementById("call-button");
callButton.onclick = () => {
  sendAction("call");
};
const checkButton = document.getElementById("check-button");
checkButton.onclick = () => {
  sendAction("check");
};
const foldButton = document.getElementById("fold-button");
foldButton.onclick = () => {
  sendAction("fold");
};
const raiseButton = document.getElementById("raise-button");
const betAmount = document.getElementById("bet-amount");
raiseButton.onclick = () => {
  const amount = betAmount.value;
  if (amount) {
    sendAction("raise");
  } else {
    alert("Please enter a raise amount.");
  }
};

function sendAction(action) {
  if (!betAmount.value) {
    betAmount.value = "0";
  }
  sendJSONMessage({
    type: "action",
    data: {
      action: action,
      amount: parseInt(betAmount.value),
    },
  });
  betAmount.value = "";
}

function sendJSONMessage(message) {
  ws.send(JSON.stringify(message));
}
