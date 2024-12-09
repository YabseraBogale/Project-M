const ul=document.getElementById("ul")

const frame=document.createDocumentFragment()
for(let i=0;i<1000;i++){
    const tr=document.createElement("tr")
    for(let j=0;j<1000;j++){
        const td=document.createElement("td")
        td.innerHTML="<input type='checkbox'>"
        tr.appendChild(td)
    }
    frame.appendChild(tr)
}
ul.appendChild(frame)