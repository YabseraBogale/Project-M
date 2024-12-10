const ul=document.getElementById("ul")

const frame=document.createDocumentFragment()
for(let i=0;i<100000;i++){
    const li=document.createElement("li")
    li.innerHTML="<input type='checkbox'>"
    frame.appendChild(li)
}
ul.appendChild(frame)