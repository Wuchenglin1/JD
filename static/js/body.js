let items = document.getElementsByClassName("item")
let btn1 = document.getElementById("btn1");
let btn2 = document.getElementById("btn2");
let points = document.getElementsByClassName("point")
let time = 0;
let index = 0;

var Clear = function () {
    for (i = 0; i < items.length; i++) {
        items[i].className = "item";
        points[i].className = "point";
    }
}
var goIndex = function () {
    Clear();
    items[index].className = "item active";
}
btn1.onclick = function goPre() {
    if (index == 0) index = items.length - 1;
    else index--;
    goIndex();
    points[index].className = "point active";
    time = 0;
}
var goNext = function () {
    if (index == items.length - 1) index = 0;
    else index++;
    goIndex();
    points[index].className = "point active";
    time = 0;
}
btn2.onclick = function goNext() {
    if (index == items.length - 1) index = 0;
    else index++;
    goIndex();
    points[index].className = "point active";
    time = 0;
};
setInterval(function () {
    time++;
    if (time == 10) {
        goNext();
        time = 0;
    }
}, 200)