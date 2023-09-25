const loginForm = document.getElementById('login-form');

// new project
loginForm.addEventListener("submit", () => {
    let username = loginForm.elements.namedItem("username").value;
    let password = loginForm.elements.namedItem("password").value;
    const url = "/admin/login";

    const body = {
        username: username,
        password: password,
    };

    debugger

    fetch(url, {
        method: "POST",
        body: JSON.stringify(body),
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => {
        if (response.ok) {
            window.location.href = "/admin"
        } else {
            alert("failed")
        }});
});