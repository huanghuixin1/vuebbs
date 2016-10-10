//重置html的fontsize
function setFontSize() {
    document.documentElement.style.fontSize = document.documentElement.clientWidth / 10 + "px";
}
setFontSize();
window.onresize = setFontSize;