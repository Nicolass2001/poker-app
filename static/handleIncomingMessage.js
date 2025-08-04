export function handleIncomingMessage(msg) {
  switch (msg.type) {
    case "gameState":
      loadPlayers(msg.data.players, msg.data.currentPlayer);
      gameStartButton(msg.data.gameState);
      loadCommunityCards(msg.data.communityCards);
      document.querySelector(".status").textContent = msg.data.gameState;
      return;
    case "playerInfo":
      loadPlayerInfo(msg.data.player);
      return;
    case "error":
      alert(msg.data);
      return;
    case "winners":
      loadWinners(msg.data);
    default:
      console.log("Unknown message type:", msg.type);
  }
}

function gameStartButton(gameState) {
  if (gameState === "Waiting for Players") {
    document.getElementById("game-start-container").style.display = "block";
    return;
  }
  document.getElementById("game-start-container").style.display = "none";
}

function loadCommunityCards(cards) {
  const communityCardsDiv = document.querySelector(".community-cards");
  communityCardsDiv.innerHTML = ""; // Clear previous cards
  cards.forEach((card) => {
    communityCardsDiv.appendChild(createCard(card));
  });
}

function loadPlayerInfo(player) {
  const playerInfoDiv = document.querySelector(".player-info");
  playerInfoDiv.querySelector(".player-name").textContent = player.name;
  playerInfoDiv.querySelector(
    ".player-chips"
  ).textContent = `Chips: $${player.stack}`;
  playerInfoDiv.querySelector(
    ".player-bet"
  ).textContent = `Bet: $${player.bet}`;
  const cardsDiv = playerInfoDiv.querySelector(".player-cards");
  cardsDiv.innerHTML = ""; // Clear previous cards
  player.cards.forEach((card) => {
    cardsDiv.appendChild(createCard(card));
  });
}

const suits = {
  Spades: "♠",
  Hearts: "♥",
  Diamonds: "♦",
  Clubs: "♣",
};

function createCard(card) {
  const cardDiv = document.createElement("div");
  if (!card.value || !card.suit) {
    return cardDiv;
  }
  cardDiv.className = "card";
  cardDiv.textContent = `${card.value}${suits[card.suit]}`;
  return cardDiv;
}

function loadPlayers(players, currentPlayer) {
  for (let i = 0; i < 6; i++) {
    const playerDiv = document.getElementById(`player${i + 1}`);
    if (i < players.length) {
      playerDiv.style.display = "block";
      playerDiv.classList.add(players[i].id);
      playerDiv.classList.remove("current-player");
      playerDiv.classList.add(
        players[i].id === currentPlayer.id ? "current-player" : "waiting-player"
      );
      playerDiv.textContent = `${players[i].name} ($${players[i].stack}) Bet: ${players[i].bet}`;
    } else {
      playerDiv.style.display = "none";
    }
  }
}

function loadWinners(winners) {
  for (const winner of winners) {
    const winnerDiv = document.querySelector("." + winner.id);
    winnerDiv.classList.add("winner");

    const winnerCardsDiv = document.createElement("div");
    winnerCardsDiv.className = "cards";
    winnerDiv.appendChild(winnerCardsDiv);
    winner.cards.forEach((card) => {
      winnerCardsDiv.appendChild(createCard(card));
    });
  }
}
