// the selectors
let username = document.querySelector(".un")

let email = document.querySelector(".em")
let password = document.querySelector(".ps")
let send = document.querySelector(".btn")


// handel the send btn to send the data 
send.addEventListener("click", async () => {
    let formated_name = username.value.trim()
    let formated_email = email.value.trim()
    let formated_password = password.value.trim()


    const json_data = {
        name: formated_name,
        email: formated_email,
        password: formated_password,
    }
    await send_user_data(json_data)
    console.log("Data Has Been Sent from the listener to the fetcher")
})
// asyncrounse function to send a POST request to the go backend
async function send_user_data(json) {
    const response = await fetch("http://localhost:3000/authintication/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(json),

    })
    // error handling 
    if (!response.ok) {
        const erro = await response.json();
        console.log(erro)
        return erro
    } else {
        const data = await response.json()
        console.log(data)
        if (data) {
            window.location.href = "/"
        }
        return data
    }
}



