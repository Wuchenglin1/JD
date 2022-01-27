
let input1 = document.getElementById("input1");
let btn1 = document.getElementById('btn1');
let phone = "";
btn1.onclick = function () {
    ajax({
        type: 'post',
        url: 'http://110.42.165.192:8080/verify/sms/register',
        data: {
            phone: input1.value,
        },
        async: true,
        success: function (data) {
            alert(data);
        }
    })
}

function ajax(opt) {
    let defaultParam = {
        type: 'get',
        url: '#',
        data: {},
        async: true,
        success: function () { },
    }

    for (key in opt) {
        defaultParam[key] = opt[key];
    }
    let paramStr = "";
    for (key in defaultParam.data) {
        paramStr += key + '=' + defaultParam.data[key] + '&';
    }
    paramStr = paramStr.substr(0, paramStr.length - 1);

    let xhr = new XMLHttpRequest();
    if (defaultParam.type == 'get') {
        xhr.open(defaultParam.type, defaultParam.url + '?' + paramStr, defaultParam.async);
    } else {
        xhr.open(defaultParam.type, defaultParam.url, defaultParam.async);
        xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
        xhr.send(paramStr);
    }

    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                defaultParam.success(xhr.responseText);
            }
        }
    }

}
