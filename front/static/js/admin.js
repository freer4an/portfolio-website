let y = document.getElementById('popup-update')
let x = document.getElementById("containerFormProject");

// visibility of project form
function toggleVisibility() {

    if (x.style.display === "none") {
        x.style.display = "block";
    } else {
        x.style.display = "none";
    }
}

function openUpdateForm(name) {
    const updateForm = document.getElementById('form-updateproject');
    let project = document.getElementById(`id-${name}`);
    let abstract = project.querySelector(".project-abstract").textContent;
    let description = project.querySelector(".project-description").textContent;
    let link = project.querySelector(".project-link").href;
    document.getElementById("old_name").innerHTML = name;
    updateForm.elements.namedItem("name").value = name;
    updateForm.elements.namedItem("abstract").value = abstract;
    updateForm.elements.namedItem("description").value = description;
    updateForm.elements.namedItem("link").value = link;
    y.style.display = "block";
}

function closeUpdateForm() {
    if (y.style.display === "block") {
        y.style.display = "none"
    }
}