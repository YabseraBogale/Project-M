fetch("http://maslovd.ru/public/mp3/Muse/Albums%20(CD)/2009%20-%20The%20Resistance%20(Japanese%20Limited%20Edition)%20-%20(320%20kbps)/")
    .then((response)=>{
        return response.text()
    }).then(data=>{
        let parser=new DOMParser()
        const parsedData=parser.parseFromString(data,"text/html")
        console.log(parsedData)
    }).catch((err)=>{
        console.log(err)
    })