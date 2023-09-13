const addProject = document.getElementById("add-project")

document.querySelectorAll(".add-project-btn").forEach(signIn =>{
    signIn.addEventListener("click", function() {
        toggleContainerVisibility(addProject);
    });
});

const overlay = document.body;

function toggleContainerVisibility(containerId) {
    var container = document.getElementById(containerId);
    if (container.classList.contains("hidden") && container === signInContainer)  {
        container.classList.remove("hidden");
        signUpContainer.classList.add("hidden");
        overlay.style.display ="block"
    } else if (container.classList.contains("hidden") && container === signUpContainer) {
        container.classList.remove("hidden");
        signInContainer.classList.add("hidden");
        overlay.style.display ="block"
    }
}