window.userUrl = "http://127.0.0.1:8080/v1/user/"; 
window.inputBox = document.querySelector(".input");
window.messageList = document.querySelector(".message-bar");
window.statusList = [];
window.userStatus = "free";

const commadHandlers = {
    ["/clear"]: () => {
        clearMessage();
    },
    ["Hi"]: () => {
        createMessage(": 你好 我是托尼");
    },
    [""]: () => {
    },
    ["/break"]: () => {
        userStatus = "free";
        statusList.length = 0;
        clearAllMessage();
        createMessage(": /login or /register, choose one...")
    }
}

var usernameRegex = /^[a-zA-Z0-9_]{1,10}$/;
var passwordRegex = /^[a-zA-Z0-9_!@#$%^&*]{8,}$/;
var emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

inputBox.addEventListener("keydown", function(event) {
    if (event.key == "Enter") {
        mainControl();
    }
});

function mainControl() {
    var inputText = inputBox.value;
    if (commadHandlers.hasOwnProperty(inputText)) {
        commadHandlers[inputText]();
        inputBox.value = "";
    } 

    if (userStatus == "free") {
       homeControl(); 
    }else if (userStatus == "register") {
        registerControl();
    } else if (userStatus == "login") {
        loginControl();
    }
}
function homeControl() {
    var inputText = inputBox.value;
    if (inputText == "/register") {
        createMessage(": Tell me your name");
        inputBox.placeholder = "Leave your username";
        userStatus = "register";
        statusList.push("createUsername");
    } else if (inputText == "/login") {
        createMessage(": Input your username");
        inputBox.placeholder = "Input your username";
        userStatus = "login";
        statusList.push("inputUsername");
    }  
    inputBox.value = '';
}

window.createMessage = function(content) {
    var newMessage = document.createElement("div");
    newMessage.textContent = content; 
    messageList.appendChild(newMessage); 
}

clearMessage = function() {
    while (messageList.childElementCount > 1) {
        messageList.removeChild(messageList.firstChild);
    }
}

clearAllMessage = function() {
    while (messageList.childElementCount > 0) {
        messageList.removeChild(messageList.firstChild);
    }
}