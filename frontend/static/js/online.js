var roomList = document.querySelector(".room-container");

class Room{
    constructor(roomId = "", roomName = "", ownerName = "", description = "") {
        this.roomId = roomId;
        this.roomName = roomName;
        this.ownerName = ownerName;
        this.description = description;
    }
}

const onlineHadlers = {
    ["/room"]: () => {
        clearAllMessage();
        footer.style.display = "none";
        getRoomlist();
    }
}
function onlineControl() {
    var inputText = inputBox.value;
    if (onlineHadlers.hasOwnProperty(inputText)) {
        onlineHadlers[inputText]();
    } 
}

async function getRoomlist() {
    token = getCookie("token");
    var requestOption = {
        method: "get",
        headers: {
            "token": token,
        }
    }
    const result = await fetch(roomUrl + "getroomlist", requestOption);
    const resultData = await result.json();
    if (result.status != 200) {
        createMessage(`: ${resultData.message}`);
    } else {
        var roomlist = resultData.message;
        for (var roomData of roomlist) {
            var room = new Room(roomData.Room_id, roomData.Name, roomData.Owner_name, roomData.Description);
            createRoom(room)
        }
    }
}


function createRoom(room) {
    var roomCard = document.createElement("div");
    var roomText = document.createElement("div");
    const roomId = document.createElement("h3");
    roomId.textContent = "Room " + room.roomId;
    const roomName = document.createElement("p");
    roomName.textContent = "Name: " + room.roomName;
    const roomOwner = document.createElement("p");
    roomOwner.textContent = "Owner: " + room.ownerName;
    const roomDescription = document.createElement("p");
    roomDescription.textContent = "Description: " + room.description;
    roomCard.className = "room";
    roomText.className = "text-container";
    roomCard.appendChild(roomId);
    roomText.appendChild(roomName);
    roomText.appendChild(roomOwner);
    roomText.appendChild(roomDescription);
    roomCard.appendChild(roomText);
    roomList.appendChild(roomCard);
}

