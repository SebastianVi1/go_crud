console.log("connected")

const buttons = document.querySelectorAll(".navigation_btn>button")



buttons.forEach(button => {
    button.addEventListener("click", () => {

        buttons.forEach(btn => {
            btn.classList.remove("isActive");
        })
        button.classList.add("isActive");
        renderView(button.textContent);
    });

});


function renderView(pageName) {
    switch (pageName) {
        case "Agregar":
            document.innerHTML += "<h1>hello</h1>"
            break;
        case "Mostrar":
            break;
        case "Actualizar":
            break;
    }
}
