var repeatPassword = 0;
var repeatPasswordTime = 0;

class User {
    constructor(name = "", email = "", password = "") {
        this.name = name; 
        this.email = email; 
        this.password = password; 
    }
}
var user = new User;

var usernameRegex = /^[a-zA-Z0-9_]{1,10}$/;
var passwordRegex = /^[a-zA-Z0-9_!@#$%^&*]{8,}$/;
var emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;


window.registerControl = function() {
    var inputText = inputBox.value;
    var registerStatus =statusList[statusList.length-1]; 
    if (registerStatus == "createUsername") {
        createUsername(inputText);
    } else if (registerStatus == "createPassword") {
        createPassword(inputText);
    } else if (registerStatus == "verifyPassword") {
        verifyPassword(inputText);
    } else if (registerStatus == "createEmail") {
        createEmail(inputText);
    }
    inputBox.value = '';
    return true;
}

function createUsername(content) {
    createMessage(content)  
    if (!usernameRegex.test(content)) {
        createMessage(": Your mom feel sad because your username invalid");
        return;
    }
    fetch(userUrl + "query/"+content, {method: "GET"})
    .then(res => {
        statusCode = res.status
        return res.json().then(data => {
            return {
                status: res.status,
                data: data
            };
        });
    })
    .then(result => {
        if (result.status != 200) {
            createMessage(": "+ result.data.message);
        }else {
            user.name = content;
            createMessage(": Setting your password and keep secretly");
            inputBox.type = "password";
            inputBox.placeholder = "Input your password"
            statusList.push("createPassword");
        }
    })
    .catch(error => {
        console.log(error);
    })
}

function createPassword(contnet) {
    if (!usernameRegex.test(contnet)) {
        createMessage(": Your mom feel sad because your username invalid");
        return; 
    }
    user.password = contnet 
    createMessage(": Input password again")
    statusList.push("verifyPassword")
}

function verifyPassword(content) {
    if (content != user.password) {
        createMessage(": Incorrect, so sad")
        repeatPasswordTime = repeatPasswordTime+1
        if (repeatPasswordTime >= 3) {
            user.password = ""
            createMessage(": Please input your password")
        }
    }else {
        repeatPassword = 1
        createMessage(": Input your email & verify")
        inputBox.type = "text"
        inputBox.placeholder = "Leave your email"
        statusList.push("createEmail")
    }
}

function createEmail(content) {
    if (!emailRegex.test(content)) {
        createMessage(": Your mom feel sad because your email invalid");
        return; 
    }
    user.email = content
    createMessage(content)
    async function fetchdataorder() {
        try{
            const result = await createUser()
            const resultData = await result.json()
            if (result.status != 200) {
                createMessage(": " + resultData.message)
                user.email = ""
                return
            }
            const mailResult = await sendMail()
            const mailResultData = await mailResult.json()
            if (mailResult.status != 200) {
                createMessage(": " + mailResultData.message)
            } else {
                createMessage("Mail was send, verfiy your mail") 
            }
        }
        catch(error) {
            console.log(error)
        }
    }
    fetchdataorder()
}

async function sendMail(content) {
    var data = {
        "Email": user.email 
    }
    var requestOption = {
        method: "POST",
        Headers: {
            "Content-type": "application/json",
        },
        body: JSON.stringify(data)
    }
    try {
        const result = await fetch(userUrl + "sendmail", requestOption)
        return result;
    } catch {
        console.log(error);
        throw error; 
    }
}

async function createUser() {
    var data = {
        "Email": user.email,
        "username": user.name,
        "password": user.password,
    }
    var requestOption = {
        method: "POST", 
        Headers: {
            "Content-type": "application/json",
        },
        body: JSON.stringify(data)
    }
    try {
        const result = await fetch(userUrl + "create", requestOption);
        return result;
    } catch (error) {
        console.log(error);
        throw error; 
    }
}

