let lis = document.getElementsByClassName("item");
let items = document.getElementsByClassName("items");
let index = 0;
let div11 = document.getElementsByClassName("div11")[0];
let i = 0;
let j = 0;

for (i = 0; i < lis.length; i++) {
    lis[i].addEventListener('click', function () {
        index = this.getAttribute('data-index');
        for (j = 0; j < lis.length; j++) {
            items[j].className = 'items';
            lis[j].className = 'item';
        }
        lis[index].className = 'item active';
        items[index].className = 'items show';
    })
}



//加入购物车
let buybtn = document.getElementsByClassName("car-plus")[0];
let deletebtn = document.getElementsByClassName("delete")[0];
let plusbtn = document.getElementsByClassName("plus")[0];
let sizes = document.getElementsByClassName("size");
let size = sizes[0].innerText;
let number = document.getElementsByClassName("number")[0];
let plus = 0;


let k = 0;

for (k; k < sizes.length; k++) {
    sizes[k].addEventListener('click', function () {
        size = this.innerText;
        for (j = 0; j < sizes.length; j++) {
            sizes[j].className = "size";
        }
        this.className = "size active";
    })
}

plus = document.getElementsByClassName("number")[0].innerText;
plusbtn.addEventListener('click', function () {
    plus++;
    number.innerText = plus;
})
deletebtn.addEventListener('click', function () {
    if (plus == 1) {
        alert("At least one!");
    }
    else {
        plus--;
        number.innerText = plus;
    }
})

let data = {
    account: plus,
    size: size,
    price
}
let paramStr = '';
for (key in data) {
    paramStr += key + '=' + data[key] + '&';
}
paramStr = paramStr.substr(0, paramStr.length - 1);
buybtn.addEventListener('click', function () {
    fetch('', {
        method: 'post',
        body: paramStr,
    })
})