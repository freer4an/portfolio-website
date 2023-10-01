const loginForm = document.getElementById('login-form');

// new project
function login() {
    let username = loginForm.elements.namedItem("username").value;
    let password = loginForm.elements.namedItem("password").value;
    const url = "/login";

    const body = {
        username: username,
        password: password,
    };

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
        }
    });
};