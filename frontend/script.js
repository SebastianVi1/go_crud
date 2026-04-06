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
                    <button type="button" onclick="" class="btn_clear">Limpiar</button>
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
                    <th>Nombre</th>
                    <th>Edad</th>
                    <th>Carrera</th>
                    <th>Promedio</th>
                    <th>Aprobado</th>
                    <th>Acciones</th>
                </thead>
                ${alumnos.map((alumno) => `
                    <tr onclick="obtenerAlumnoPorId(${alumno.id})"style="cursor: pointer;">
                        <td>${alumno.nombre}</td>
                        <td>${alumno.edad}</td>
                        <td>${alumno.carrera}</td>
                        <td>${alumno.promedio}</td>
                        <td>${alumno.aprobado ? "Si" : "No"}</td>
                        <td class="action-buttons">
                            <button onclick="actualizarAlumno(${alumno.id})" class="btn_update">Actualizar</button>
                            <button onclick="eliminarAlumno(${alumno.id})" class="btn_delete ">Eliminar</button>
                        </td>
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
        mostrarNotificacion('Alumno creado correctamente');
        clearData();

    } catch (error) {
        console.log("Error al crear alumno", error);
        mostrarNotificacion('Error al crear el alumno', true);
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

async function obtenerAlumnoPorId(alumnoId) {
    console.log("Mouse/click sobre alumno con ID:", alumnoId);

    try {
        var response = await fetch(`http://localhost:8080/alumnos/${alumnoId}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        var alumno = await response.json();
        console.log("Datos del alumno:", alumno);
        return alumno;

    } catch (err) {
        console.log("Error al obtener alumno por ID:", err);
    }
}


async function eliminarAlumno(alumnoId) {
    const confirmado = await modalConfirmacion('¿Estás seguro de eliminar el alumno?');
    if (confirmado) {
        await borrarAlumno(alumnoId);
        renderView("Mostrar")
    }
}

async function borrarAlumno(alumnoId) {
    try {
        let response = await fetch('http://localhost:8080/alumnos', {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "id": alumnoId
            })
        })
        let data = await response.json();
        console.log("Alumno eliminado", data);
        mostrarNotificacion('Alumno eliminado correctamente');
        renderView("Mostrar");

    } catch (err) {
        console.log("Error al eliminar alumno", err);
        mostrarNotificacion('Error al eliminar alumno', true);
    }
}

async function actualizarAlumno(alumnoId) {
    let alumno = await obtenerAlumnoPorId(alumnoId);
    renderView("Agregar");

    document.querySelector("#nombre").value = alumno.nombre;
    document.querySelector("#edad").value = alumno.edad;
    document.querySelector("#carrera").value = alumno.carrera;
    document.querySelector("#promedio").value = alumno.promedio;

    let btn_submit = document.querySelector(".inputs_form button[onclick='guardarAlumno()']");
    if (btn_submit) {
        btn_submit.textContent = "Actualizar";
        btn_submit.setAttribute("onclick", `modificarAlumno(${alumnoId})`);
    }
}

async function modificarAlumno(alumnoId) {
    var nombre = document.querySelector("#nombre").value;
    var edad = document.querySelector("#edad").value;
    var carrera = document.querySelector("#carrera").value;
    var promedio = document.querySelector("#promedio").value;

    try {
        const response = await fetch('http://localhost:8080/alumnos', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "id": alumnoId,
                "nombre": nombre,
                "edad": parseInt(edad),
                "carrera": carrera,
                "promedio": parseFloat(promedio)
            })
        });
        const data = await response.json();
        console.log('Alumno Modificado', data);
        mostrarNotificacion('Alumno modificado correctamente');
        clearData();
        renderView("Mostrar");

    } catch (error) {
        console.log("Error al modificar alumno", error);
        mostrarNotificacion('Error al modificar el alumno', true);
    }
}

let modalResolve;

function modalConfirmacion(mensaje) {
    return new Promise((resolve) => {
        modalResolve = resolve;
        const modalDiv = document.createElement('div');
        modalDiv.id = 'active_modal';
        modalDiv.innerHTML = `
        <div class="modal">
            <div class="modal_content">
                <p>${mensaje}</p>
                <div class="modal_buttons">
                    <button class="btn_cancelar" onclick="modalCancelar()">Cancelar</button>
                    <button class="btn_aceptar" onclick="modalAceptar()">Aceptar</button>
                </div>
            </div>
        </div>
        `;
        document.body.appendChild(modalDiv);
    });
}

function modalCancelar() {
    const modalDiv = document.getElementById('active_modal');
    if (modalDiv) modalDiv.remove();
    if (modalResolve) modalResolve(false);
}

function modalAceptar() {
    const modalDiv = document.getElementById('active_modal');
    if (modalDiv) modalDiv.remove();
    if (modalResolve) modalResolve(true);
}

function mostrarNotificacion(mensaje, isError = false) {
    let container = document.getElementById("toast_container");
    if (!container) {
        container = document.createElement("div");
        container.id = "toast_container";
        document.body.appendChild(container);
    }

    const toast = document.createElement("div");
    toast.className = `toast ${isError ? 'error' : ''}`;
    toast.textContent = mensaje;

    container.appendChild(toast);


    setTimeout(() => {
        toast.classList.add("visible");
    }, 10);

    setTimeout(() => {
        toast.classList.remove("visible");
        setTimeout(() => {
            toast.remove();
        }, 300);
    }, 3000);
}
