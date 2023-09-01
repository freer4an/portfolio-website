const navbar = document.querySelector(".navbar")

const projectsLink = navbar.querySelector('.nav-projects');

projectsLink.addEventListener('click', event => {
    event.preventDefault();
    
    const currentPage = 1;
    const queryParams = {page: currentPage};
    const url = new URL(projectsLink.href);
    Object.keys(queryParams).forEach(key => {
        url.searchParams.append(key, queryParams[key]);
    });

    window.location.href = url.href;
});
