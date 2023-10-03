const newProjectForm = document.getElementById('formProject');

const POST = "POST"
const GET = "GET"
const PATCH = "PATCH"
const DELETE = "DELETE"

// new project
function addProject() {
    let name = newProjectForm.elements.namedItem("name").value;
    let abstract = newProjectForm.elements.namedItem("abstract").value;
    let description = newProjectForm.elements.namedItem("description").value;
    let link = newProjectForm.elements.namedItem("link").value;
    const url = "/projects";
    
    const body = {
        name: name,
        abstract: abstract,
        description: description,
        link: link,
    };
    submitProject(body, url, POST);
};

// update project
function updateProject(name) {
    let newName = updateProjectForm.elements.namedItem("name").value;
    let abstract = updateProjectForm.elements.namedItem("abstract").value;
    let description = updateProjectForm.elements.namedItem("description").value;
    let link = updateProjectForm.elements.namedItem("link").value;
    const url = "/projects/" + name;
    
    const body = {
        name: name,
        abstract: abstract,
        description: description,
        link: link,
    };
    submitProject(body, url, POST);
};

// delete project
function deleteProject(name) {
    const url = "/projects/" + name;
    submitProject(null, url, DELETE)
}

function submitProject(body, url, method) {
    if (body != null) {
        body = JSON.stringify(body)
    }
    fetch(url, {
        method: method,
        body: body,
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            window.location.reload()
            return 
        } else {
            return response.text().then(text => {
                throw new Error(`${Object.values(JSON.parse(text))}`)
            })
        }
    })
    .catch(error => {
        alert(`${error}`)
    })
};

function deleteNullFileds(obj) {
    Object.keys(obj).forEach(key => {
        if (obj[key] === null || !obj[key]) {
          delete obj[key];
        }
    });
}