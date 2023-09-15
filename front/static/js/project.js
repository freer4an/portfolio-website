const addProject = document.getElementById("add-project")
const newProjectForm = document.getElementById('formProject');


// new project
newProjectForm.addEventListener("submit", () =>{
    let name = document.getElementById("inputName").value;
    let abstract = document.getElementById("inputAbstract").value;
    let description = document.getElementById("inputDescription").value;
    let link = document.getElementById("inputLink").value;
    const url = "/admin/projects"
    
    const body = {
        name: name,
        abstract: abstract,
        description: description,
        link: link,
    };
    submitProject(body, url, POST)
})

// update project
// updateProjectForm.addEventListener("submit", () =>{
//     let name = document.getElementById("inputName").value;
//     let abstract = document.getElementById("inputAbstract").value;
//     let description = document.getElementById("inputDescription").value;
//     let link = document.getElementById("inputLink").value;
//     const url = "/admin/projects"
    
//     const body = {
//         name: name,
//         abstract: abstract,
//         description: description,
//         link: link,
//     };
//     submitProject(body, url, POST)
// })

const overlay = document.body;