window.hostUrl = "http://127.0.0.1:8080/"
window.userUrl = hostUrl + "v1/user/";
window.roomUrl = hostUrl + "v1/room/";
window.inputBox = document.querySelector(".input");
window.footer = document.querySelector(".footer")
window.usernameTag = document.querySelector(".username");
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
    ["/home"]: () => {
        homeInit(); 
    },
    ["/logout"]: () => {
        logout();
    }
}

window.usernameRegex = /^[a-zA-Z0-9_]{1,20}$/;
window.passwordRegex = /^[a-zA-Z0-9_!@#$%^&*]+$/;
window.emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

inputBox.addEventListener("keydown", function(event) {
    if (event.key == "Enter") {
        mainControl();
    }
});

window.onload = async function() {
    const result = await checkUserAlive();
    if (result == true) { 
       loginInit(); 
    }
}

function mainControl() {
    var inputText = inputBox.value;
    if (commadHandlers.hasOwnProperty(inputText)) {
        commadHandlers[inputText]();
    } else if (userStatus == "free") {
        homeControl(); 
    } else if (userStatus == "register") {
        registerControl();
    } else if (userStatus == "login") {
        loginControl();
    } else if (userStatus == "online") {
        onlineControl();
    } 
    inputBox.value = "" 
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
}

window.createMessage = function(content) {
    var newMessage = document.createElement("div");
    newMessage.textContent = content; 
    newMessage.className = "message";
    messageList.appendChild(newMessage); 
}

clearMessage = function() {
    while (messageList.childElementCount > 1) {
        messageList.removeChild(messageList.firstChild);
    }
}

window.clearAllMessage = function() {
    while (messageList.childElementCount > 0) {
        messageList.removeChild(messageList.firstChild);
    }
}

window.loginInit = function(){
    userStatus = "online";
    statusList.length = 0;
    inputBox.placeholder = `Welcome ${user.name}`;
    inputBox.type = "text";
    usernameTag.textContent = `User_${user.name}`;
    clearAllMessage();
    createMessage(": Welcome back")
}

window.homeInit = function() {
    statusList.length = 0;
    inputBox.type = "text";
    clearAllMessage();

    if (userStatus == "online") {
        inputBox.placeholder = `Welcome ${user.name}`
        createMessage(": Welcome back") 
    } else {
        userStatus = "free";

        createMessage(": /login or /register, choose one...")
        inputBox.placeholder = "Say Hi"; 
    }
}

function logout() {
    userStatus = "free";
    statusList.length = 0;
    inputBox.type = "text";
    inputBox.placeholder = "Say Hi";
    usernameTag.textContent = `Guest`;
    clearAllMessage();
    createMessage(": /login or /register, choose one...");
    deleteCookie("token");
}
