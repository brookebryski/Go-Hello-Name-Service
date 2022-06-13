const nameInput = document.getElementById("name")
const submitNameButton = document.getElementById("name-submit")

submitNameButton.addEventListener("click", function() {
    let data = {
        Name: nameInput.value
    };
    fetch("/get_name", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            console.log(result)
        });
    }) .catch((error) => {
        console.log(error)
    });
})