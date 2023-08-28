window.checkUserAlive = async function () {
    token = getCookie("token");
    if (token) {
        requestOption = {
            method: "get",
            headers: {
                "token": token,
            }
        }
        const result = await fetch(userUrl + "useralive", requestOption);
        const resultData = await result.json();
        if (result.status == 200) {
            user.name = resultData.username;
            return true;
        } else {
            return false;
        } 
    }
}

function updateCurrentTime() {
    const currentTimeElement = document.querySelector(".time");
    const now = new Date();
    const hours = formatTimeString(now.getHours());
    const minutes = formatTimeString(now.getMinutes());
    const seconds = formatTimeString(now.getSeconds());
    const formattedTime = `${hours}:${minutes}:${seconds}`;
    currentTimeElement.textContent = formattedTime;
}

function formatTimeString(time) {
    return time < 10 ? "0" + time : time.toString();
}

window.getCookie = function(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    
    if (parts.length === 2) {
        return parts.pop();
    } else {
        return "";
    }

}

window.deleteCookie = function(name) {
    document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}

setInterval(checkUserAlive, 60000);
setInterval(updateCurrentTime, 100);