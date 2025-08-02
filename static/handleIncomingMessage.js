export function handleIncomingMessage(msg) {
  if (msg.type === "gameState") {
    loadPlayers(msg.data.players);
    if (msg.data.gameState === "Waiting for Players") {
      document.getElementById("game-start-container").style.display = "block";
      return;
    }
    document.getElementById("game-start-container").style.display = "none";
    loadCommunityCards(msg.data.communityCards);
  }
  if (msg.type === "playerInfo") {
    loadPlayerInfo(msg.data.player);
  } else {
    console.log("Unknown message type:", msg.type);
    return;
  }
  document.getElementsByClassName(".status").textContent = msg.data.gameState;
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

function loadPlayers(players) {
  for (let i = 0; i < 6; i++) {
    const playerDiv = document.getElementById(`player${i + 1}`);
    if (i < players.length) {
      playerDiv.textContent = `${players[i].name} ($${players[i].stack}) Bet: ${players[i].bet}`;
    } else {
      playerDiv.textContent = "";
    }
  }
}
