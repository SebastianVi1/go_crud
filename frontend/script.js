console.log("connected")

const buttons = document.querySelectorAll(".navigation_btn>button")

const buttonObj = {
    isActive: false,
    buttonName: ""
}

const estado = new Proxy(buttonObj, {
    set(target, property, value) {
        target[property] = value;
        renderView(target.buttonName);
    }
});

buttons.forEach(button => {
    button.addEventListener("click", () => {

        buttons.forEach(btn => {
            btn.classList.remove("isActive");
        })
        button.classList.add("isActive");
        renderView(button.textContent);
    });

});


async function renderView(pageName) {
    switch (pageName) {
        case "Agregar":
            var dinamic_content = document.querySelector(".dinamic_content");
            dinamic_content.innerHTML = `
            <section class="inputs_form">
                <form action="">
                    <label for="nombre">Nombre:</label>
                    <input type="text" id="nombre" name="nombre">
                    <label for="edad">Edad:</label>
                    <input type="number" id="edad" name="edad">
                    <label for="carrera">Carrera:</label>
                    <input type="text" id="carrera" name="carrera">
                    <label for="promedio">Promedio:</label>
                    <input type="number" id="promedio" name="promedio">
                    <button type="button" onclick="guardarAlumno()">Agregar</button>
                </form>
            </section>
            `
            break;
        case "Mostrar":
            var dinamic_content = document.querySelector(".dinamic_content");
            var alumnos = await obtenerAlumnos();

            dinamic_content.innerHTML = `
            <table class="table_alumnos">
                <thead>
                    <th>ID</th>
                    <th>Nombre</th>
                    <th>Edad</th>
                    <th>Carrera</th>
                    <th>Promedio</th>
                    <th>Aprobado</th>
                </thead>
                ${alumnos.map((alumno) => `
                    <tr>
                        <td>${alumno.id}</td>
                        <td>${alumno.nombre}</td>
                        <td>${alumno.edad}</td>
                        <td>${alumno.carrera}</td>
                        <td>${alumno.promedio}</td>
                        <td>${alumno.aprobado ? "Si" : "No"}</td>
                    </tr>
                `).join('')}
                </tbody>
            </table>`
            break;
        case "Actualizar":
            break;
    }
}


async function guardarAlumno() {

    var nombre = document.querySelector("#nombre").value;
    var edad = document.querySelector("#edad").value;
    var carrera = document.querySelector("#carrera").value;
    var promedio = document.querySelector("#promedio").value;

    try {
        const response = await fetch(
            'http://localhost:8080/alumnos',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "nombre": nombre,
                    "edad": parseInt(edad),
                    "carrera": carrera,
                    "promedio": parseFloat(promedio)
                })
            });
        const data = await response.json();
        console.log('Alumno Creado', data);
        clearData();

    } catch (error) {
        console.log("Error al crear alumno", error);
    }

}

function clearData() {
    var inputs = document.querySelectorAll("input");
    inputs.forEach(input => {
        input.value = "";
    });
}

async function obtenerAlumnos() {

    try {
        var response = await fetch('http://localhost:8080/alumnos', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        var alumnos = await response.json();
        console.log(alumnos);
        return alumnos;

    } catch (error) {
        console.log("Error:" + error);
    }


}

