async function obtenerObrasPorCategoria({ categoriaId, precioMax, estado, artistaId }) {
    const params = new URLSearchParams();

    if (categoriaId) params.append('categoria_id', categoriaId);
    if (precioMax) params.append('precio_max', precioMax);
    if (estado) params.append('estado', estado);
    if (artistaId) params.append('artista_id', artistaId);

    const url = `http://127.0.0.1:8080/category-artworks?${params.toString()}`;

    try {
        const respuesta = await fetch(url);
        const data = await respuesta.json();
        return data;
    } catch (error) {
        console.error("Error al obtener obras por categoría:", error);
        return [];
    }
}

document.querySelector('button').addEventListener('click', async function() {
    const categoria_id = document.getElementById('categoria_id').value;
    const precio_max = document.getElementById('precio_max').value;
    const estado = document.getElementById('estado').value;
    const artista_id = document.getElementById('artista_id').value;

    if (!categoria_id) {
    alert("Por favor ingresa el ID de la categoria.");
    return;
    }

    const obras = await obtenerObrasPorCategoria({
        categoriaId: categoria_id,
        precioMax: precio_max,
        estado: estado,
        artistaId: artista_id
    });

    const tbody = document.querySelector('tbody');
    tbody.innerHTML = '';

    obras.forEach(obra => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${obra.categoria_id}</td>
            <td>${obra.categoria}</td>
            <td>${obra.obra_id}</td>
            <td>${obra.titulo}</td>
            <td>${obra.precio_referencia}</td>
            <td>${obra.estado}</td>
            <td>${obra.artista_id}</td>
            <td>${obra.total_ventas}</td>
        `;
        tbody.appendChild(tr);
    });
});
