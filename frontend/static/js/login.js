window.loginControl = function() {
    var inputText = inputBox.value;
    loginStatus = statusList[statusList.length-1];
    if (loginStatus == "inputUsername") {
        user.name = inputText;
        checkUsernameValidate(user.name);
    } else if (loginStatus == "inputPassword") {
        user.password = inputText;
        userLogin(user.name, user.password);
    }
}

async function userLogin(username, password) {
    if (!passwordRegex.test(password)) {
        createMessage(": Your password format invalid");
    } else {
        var data = {
            "Username": username,
            "Password": password,
        }
        var requestOption = {
            method: "post",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify(data)
        }
        const result = await fetch(userUrl + "login", requestOption);
        const resultData = await result.json();
        if (result.status == 200) {
            var token = resultData.message;
            document.cookie = `token=${token}; expires=${new Date(Date.now() + 24 * 60 * 60 * 1000).toUTCString()}; path=/`;
            loginInit();
        } else {
            createMessage(": " + resultData.message) 
        }
    }
} 

function checkUsernameValidate(username) {
    if (!usernameRegex.test(username)) {
        createMessage(": Your username format invalid");
    } else {
        createMessage(username)
        createMessage(": Input your password");
        inputBox.type = "password";
        inputBox.placeholder = "Input your password";
        statusList.push("inputPassword");
    }
}