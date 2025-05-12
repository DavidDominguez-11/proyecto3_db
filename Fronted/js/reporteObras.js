async function obtenerObras({ estado, estilo, pais, precioMin, precioMax }) {
    const params = new URLSearchParams();

    if (estado) params.append('estado', estado);
    if (estilo) params.append('estilo', estilo);
    if (pais) params.append('pais', pais);
    if (precioMin) params.append('precio_min', precioMin);
    if (precioMax) params.append('precio_max', precioMax);

    const url = `http://127.0.0.1:8080/artworks-report?${params.toString()}`;

    try {
        const respuesta = await fetch(url);
        const data = await respuesta.json();
        return data;
    } catch (error) {
        console.error("Error al obtener reporte de obras:", error);
        return [];
    }
}

document.querySelector('button').addEventListener('click', async function() {
    const estado = document.getElementById('estado').value;
    const estilo = document.getElementById('estilo').value;
    const pais = document.getElementById('pais').value;
    const precio_min = document.getElementById('precio_min').value;
    const precio_max = document.getElementById('precio_max').value;

    const obras = await obtenerObras({
        estado: estado,
        estilo: estilo,
        pais: pais,
        precioMin: precio_min,
        precioMax: precio_max
    });

    const tbody = document.querySelector('tbody');
    tbody.innerHTML = '';

    if (!obras || obras.length === 0) {
        alert("No se encontraron obras con los filtros seleccionados.");
        return;
    }

    obras.forEach(obra => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${obra.obra_id}</td>
            <td>${obra.titulo}</td>
            <td>${obra.estado}</td>
            <td>${obra.precio_referencia}</td>
            <td>${obra.año_creacion}</td>
            <td>${obra.estilo_principal}</td>
            <td>${obra.pais_origen}</td>
        `;
        tbody.appendChild(tr);
    });
});
