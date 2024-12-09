const ul=document.getElementById("ul")

const frame=document.createDocumentFragment()
for(let i=0;i<50;i++){
    const li=document.createElement("li")
    const input=document.createElement("input")
    li.innerHTML=input.checked
    frame.appendChild(li)
}
ul.appendChild(frame)