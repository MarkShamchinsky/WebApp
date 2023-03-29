let button = document.querySelector(".form > button");

if (button) {
    button.onclick = function (e) {
        let inputs = document.querySelectorAll(".form > input");
        let data = {}
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/user/reg");
        xhr.onload = function (e) {
            console.log(e);
        };
        console.log(JSON.stringify(data))
        xhr.send(JSON.stringify(data));
    }
}
