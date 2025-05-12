async function obtenerTransacciones({ tipo, transaccionId, entidadAfectadaId, fechaInicio }) {
    const params = new URLSearchParams();

    if (tipo) params.append('tipo', tipo);
    if (transaccionId) params.append('transaccion_id', transaccionId);
    if (entidadAfectadaId) params.append('entidad_afectada_id', entidadAfectadaId);
    if (fechaInicio) params.append('fecha_inicio', new Date(fechaInicio).toISOString());

    const url = `http://127.0.0.1:8080/transactions?${params.toString()}`;

    try {
        const respuesta = await fetch(url);
        if (!respuesta.ok) throw new Error("Error en la respuesta del servidor");
        return await respuesta.json();
    } catch (error) {
        console.error("Error al obtener transacciones:", error);
        return [];
    }
}

document.getElementById('filtrar').addEventListener('click', async () => {
    const tipo = document.getElementById('tipo').value;
    const transaccion_id = document.getElementById('transaccion_id').value;
    const entidad_afectada_id = document.getElementById('entidad_afectada_id').value;
    const fecha_inicio = document.getElementById('fecha_inicio').value;

    const transacciones = await obtenerTransacciones({
        tipo,
        transaccionId: transaccion_id,
        entidadAfectadaId: entidad_afectada_id,
        fechaInicio: fecha_inicio
    });

    const tbody = document.querySelector('tbody');
    tbody.innerHTML = '';

    if (!transacciones || transacciones.length === 0) {
        alert("No se encontraron transacciones con los filtros seleccionados.");
        return;
    }

    transacciones.forEach(transaccion => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${transaccion.transaccion_id}</td>
            <td>${transaccion.tipo}</td>
            <td>${new Date(transaccion.fecha).toLocaleString()}</td>
            <td>${transaccion.entidad_afectada_id}</td>
            <td>${transaccion.detalle}</td>
        `;
        tbody.appendChild(tr);
    });
});
