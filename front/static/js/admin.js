let y = document.getElementById('popup-update')
let x = document.getElementById("containerFormProject");



// visibility of project form
function toggleVisibility() {

    if (y.style.display === "block") {
        y.style.display = "none";
    }

    if (x.style.display === "none") {
        x.style.display = "block";
    } else {
        x.style.display = "none";
    }
}

function openUpdateForm(name) {
    let project = document.getElementById(`id-${name}`)

    console.log(project.childNodes.forEach())
}