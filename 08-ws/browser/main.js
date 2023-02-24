
const ws = new WebSocket("wss://192.168.17.129:9000/ws");

//const ws = new WebSocket("ws://192.168.124.13:8001/ws");

ws.onerror = (e)=>{
    console.log(e)
}
ws.onclose = (event)=> {
    if (event.wasClean) {
        console.info('Connection close was clean');
    }
    else {
        console.error('Connection suddenly close');
    }
    console.info('close code : ' + event.code + ' reason: ' + event.reason);
};
ws.onopen = (e)=>{
    console.log("websocket is open!")
};
ws.onmessage = (e)=>{
    console.log("onmessage:")
    let msg = JSON.parse(e.data)
    console.log(msg)
}
const send = (target,eventType,dataObj)=>{
    ws.send(JSON.stringify({
        target:target,
        event:eventType,
        data:JSON.stringify(dataObj)
    }))
}

const test = ()=>{
    ws.send(JSON.stringify({
        type:"client-test",
        data:"hello,i am client"
    }))
}

