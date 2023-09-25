const POST = "POST"
const GET = "GET"
const PATCH = "PATCH"
const DELETE = "DELETE"

// visibility of project form
function toggleVisibility() {
    var x = document.getElementById("containerFormProject");
    if (x.style.display === "none") {
        x.style.display = "block";
    } else {
        x.style.display = "none";
    }
}

function submitProject(body, url, method) {
    fetch(url, {
        method: method,
        body: JSON.stringify(body),
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => {
        if (response.ok) {
            window.location.reload()
            console.log("Succesfully submited")
        } else {
            alert("Failed")
            console.error("Failed to submit");
        }
    }).catch(error => {
        console.error(error);
    });
};

function deleteProject(name) {
    const url = "admin/projects/" + name
    fetch(url, {
        method: DELETE,
    }).then(response => {
        if (response.ok) {
            window.location.reload()
            console.log("Succesfully deleted")
        } else {
            alert("Failed")
            console.error("Failed to delete");
        }
    }).catch(error => {
        console.error(error);
    });
}

function deleteNullFileds(obj) {
    Object.keys(obj).forEach(key => {
        if (obj[key] === null) {
          delete obj[key];
        }
    });
}

