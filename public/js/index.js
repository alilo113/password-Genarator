const btn = document.getElementById("btn");
const div = document.getElementById("div");
const characters = ["a", "b", "c", "d", "e", "F", "G", "H", "I", "J","@","#", "!", 1, 2, 3, 4, 5];
const passWordLength = 16;

const charactersStr = characters.join("");
let str = "";
btn.addEventListener("click", () => {
  for(let i = 0; i < passWordLength; i++){
    str += charactersStr.charAt(Math.floor(Math.random() * characters.length));
  }
  div.innerHTML = str;
})