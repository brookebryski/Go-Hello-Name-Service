const nameInput = document.getElementById("name")
const submitNameButton = document.getElementById("name-submit")

submitNameButton.addEventListener("click", function() {
    let data = {
        Name: nameInput.value
    };
    console.log(data)
})