
const networks = document.querySelector('.net')
const devices = document.querySelector('.dev')
const values = document.querySelector('.val')
const states = document.querySelector('.stat')
console.log(networks)

networks.addEventListener("click", () => {
    fetch("http://localhost:8000/network", { mode: 'no-cors'})
        .then(response => {
            return response.json();
        })
        .then(data => {
            console.log(data);
        })
        .catch(err => {
            console.log(err);
        }) 
});

 const ws = new WebSocket("ws:localhost:8000/subscribe")
 ws.addEventListener("open", () => {
     console.log("Connection established")
 })
 ws.addEventListener("message", (data) => {
     let tempData = document.getElementById("temp");
     const jsonData = JSON.parse(data.data)
     console.log(jsonData);
     tempData.innerHTML = jsonData.data
 })