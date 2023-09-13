const POST = "POST"
const GET = "GET"
const PATCH = "PATCH"
const DELETE = "DELETE"


let addProjectBtn = document.getElementById("addProjectBtn")

// new project
addProjectBtn.addEventListener("click", () =>{
    let name = document.getElementById("project-name").value;
    let title = document.getElementById("project-title").value;
    let abstract = document.getElementById("project-abstract").value;
    let description = document.getElementById("project-description").value;
    let link = document.getElementById("project-link").value;
    const url = "/projects"
    
    const body = {
        abstract: abstract,
        title: title,
        description: description,
        link: link,
    };

    deleteNullFileds(body)
    console.log(body)
    // todo: remove stop
    stop
    submitProject(body, url, POST)
})

function submitProject(body, url, method) {
    fetch(url, {
        method: method,
        body: JSON.stringify(body),
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            window.location.reload()
            console.log("Succesfully submited")
        } else {
            alert("Failed")
            console.error("Failed to submit");
        }
    })
    .catch(error => {
        console.error(error);
    });
}

function deleteProject(name) {
    const url = "admin/projects/" + name
    fetch(url, {
        method: DELETE,
    })
    .then(response => {
        if (response.ok) {
            window.location.reload()
            console.log("Succesfully deleted")
        } else {
            alert("Failed")
            console.error("Failed to delete");
        }
    });
}

function deleteNullFileds(obj) {
    Object.keys(obj).forEach(key => {
        if (obj[key] === null) {
          delete obj[key];
        }
    });
}

