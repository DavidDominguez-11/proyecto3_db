async function obtenerVentas({ fechaInicio, metodoPago, paisArtista, estadoEnvio }) {
    const params = new URLSearchParams();

    if (fechaInicio) params.append('fecha_inicio', fechaInicio);
    if (metodoPago) params.append('metodo_pago', metodoPago);
    if (paisArtista) params.append('pais_artista', paisArtista);
    if (estadoEnvio) params.append('estado_envio', estadoEnvio);

    const url = `http://127.0.0.1:8080/sales-report?${params.toString()}`;

    try {
        const respuesta = await fetch(url);
        const data = await respuesta.json();
        return data;
    } catch (error) {
        console.error("Error al obtener reporte de ventas:", error);
        return [];
    }
}

document.querySelector('button').addEventListener('click', async function() {
    const fecha_inicio = document.getElementById('fecha_inicio').value;
    const metodo_pago = document.getElementById('metodo_pago').value;
    const pais_artista = document.getElementById('pais_artista').value;
    const estado_envio = document.getElementById('estado_envio').value;

    const ventas = await obtenerVentas({
        fechaInicio: fecha_inicio,
        metodoPago: metodo_pago,
        paisArtista: pais_artista,
        estadoEnvio: estado_envio
    });

    const tbody = document.querySelector('tbody');
    tbody.innerHTML = '';

    if (!ventas || ventas.length === 0) {
        alert("No se encontraron ventas con los filtros seleccionados.");
        return;
    }


    ventas.forEach(venta => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${venta.venta_id}</td>
            <td>${venta.comprador}</td>
            <td>${venta.obra}</td>
            <td>${venta.pais_artista}</td>
            <td>${venta.monto}</td>
            <td>${venta.metodo_pago}</td>
            <td>${venta.fecha_venta}</td>
            <td>${venta.estado_envio}</td>
        `;
        tbody.appendChild(tr);
    });
});
