async function obtenerOfertas({ fechaInicio, montoMin, usuarioId }) {
    const params = new URLSearchParams();

    if (fechaInicio) params.append('fecha_inicio', fechaInicio);
    if (montoMin) params.append('monto_min', montoMin);
    if (usuarioId) params.append('usuario_id', usuarioId);

    const url = `http://127.0.0.1:8080/auctions/11/offers?${params.toString()}`;

    try {
        const respuesta = await fetch(url);
        const data = await respuesta.json();
        return data;
    } catch (error) {
        console.error("Error al obtener ofertas:", error);
        return [];
    }
}

document.querySelector('button').addEventListener('click', async function() {
    const fecha_inicio = document.getElementById('fecha_inicio').value;
    const monto_min = document.getElementById('monto_min').value;
    const usuario_id = document.getElementById('usuario_id').value;

    const ofertas = await obtenerOfertas({
        fechaInicio: fecha_inicio,
        montoMin: monto_min,
        usuarioId: usuario_id
    });

    const tbody = document.querySelector('tbody');
    tbody.innerHTML = '';

    ofertas.forEach(oferta => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${oferta.oferta_id}</td>
            <td>${oferta.subasta_id}</td>
            <td>${oferta.obra}</td>
            <td>${oferta.ofertante}</td>
            <td>${oferta.monto_ofertado}</td>
            <td>${oferta.fecha_oferta}</td>
        `;
        tbody.appendChild(tr);
    });
});
