let div1 = document.getElementById("div1");
// var msg = "Hello, world!!!";
let x = 5, y;
y = x++;
// y = 5, x = 6
let msg;
msg = "<b>Постфиксная форма (y = x++;):<" + "/b><br> y = ";
msg += y + "<br>x = " + x + "<br><br>";
x = 5;
y = ++x;
// y = 6, x = 6
msg += "<b>Префиксная форма (y = ++x;):<" + "/b><br> y = ";
msg += y + "<br>x = " + x;

div1.innerHTML = msg;
