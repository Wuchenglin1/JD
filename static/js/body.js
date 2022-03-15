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

let wrapper = document.getElementsByClassName("wrapper")[0];
let slider = document.getElementsByClassName("slider")[0];
let left = 0;
let time1 = 0
wrapper.innerHTML += wrapper.innerHTML;
function autoPlay() {
    left -= 2;
    if (left === -(150 * 7 + 10 * 7)) {
        left = 0;
    }
    wrapper.style.left = left + 'px';

}
function timerPlay() {
    time1 = setInterval(function () {
        autoPlay();
    }, 20)
}
timerPlay();
slider.onmouseover = function () {
    clearInterval(time1);
}
slider.onmouseout = function () {
    timerPlay();
}
//侧边工具栏
let backToTop = document.getElementById('top');

backToTop.onclick = function () {

}