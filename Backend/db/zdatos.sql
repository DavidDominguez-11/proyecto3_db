-- 1. Usuarios (30 registros)
INSERT INTO Usuario (nombre, correo, tipo_usuario, fecha_registro)
SELECT 
    'Usuario' || num,
    'usuario' || num || '@galeria.com',
    CASE 
        WHEN num <= 10 THEN 'vendedor'
        WHEN num <= 25 THEN 'comprador'
        ELSE 'admin'
    END,
    CURRENT_DATE - (FLOOR(RANDOM() * 365) || ' days')::INTERVAL
FROM generate_series(1, 30) as num;

-- 2. Perfiles de Artistas (10 registros, solo vendedores)
INSERT INTO PerfilArtista (usuario_id, biografia, pais_origen, estilo_principal)
SELECT 
    usuario_id,
    'Artista destacado en ' || 
    CASE (RANDOM() * 4)::INT
        WHEN 0 THEN 'pintura al óleo'
        WHEN 1 THEN 'escultura moderna'
        WHEN 2 THEN 'fotografía conceptual'
        ELSE 'arte digital'
    END,
    CASE (RANDOM() * 4)::INT
        WHEN 0 THEN 'México'
        WHEN 1 THEN 'España'
        WHEN 2 THEN 'Argentina'
        ELSE 'Colombia'
    END,
    CASE (RANDOM() * 3)::INT
        WHEN 0 THEN 'Realismo'
        WHEN 1 THEN 'Surrealismo'
        WHEN 2 THEN 'Arte abstracto'
        ELSE 'Pop Art'
    END
FROM Usuario WHERE tipo_usuario = 'vendedor' LIMIT 10;

-- 3. Categorías (6 registros básicos)
INSERT INTO Categoria (nombre, descripcion) VALUES
('Pintura', 'Obras realizadas con técnicas de pintura'),
('Escultura', 'Arte tridimensional'),
('Fotografía', 'Arte visual capturado con cámara'),
('Arte Digital', 'Creaciones usando medios digitales'),
('Arte Textil', 'Obras con tejidos o fibras'),
('Mixed Media', 'Combinación de múltiples técnicas');

-- 4. Obras de Arte (30 registros)
INSERT INTO ObraArte (titulo, descripcion, año_creacion, precio_referencia, estado, artista_id)
SELECT 
    'Obra ' || TO_CHAR(obra_num, 'FM999'),
    'Descripción detallada de la obra ' || TO_CHAR(obra_num, 'FM999'),
    2000 + (RANDOM() * 24)::INT,
    (500 + RANDOM() * 9500)::NUMERIC(10,2),
    CASE 
        WHEN RANDOM() < 0.6 THEN 'en venta'
        WHEN RANDOM() < 0.8 THEN 'subasta'
        ELSE 'vendida'
    END,
    artista_id
FROM PerfilArtista
CROSS JOIN generate_series(1, 3) as obra_num; -- 3 obras por artista

-- 5. Asignación de Categorías (60 relaciones obra-categoría)
INSERT INTO ObraCategoria (obra_id, categoria_id)
SELECT 
    obra_id,
    (1 + (RANDOM() * 5)::INT) -- Categorías del 1 al 6
FROM ObraArte
CROSS JOIN generate_series(1, 2); -- 2 categorías por obra en promedio

-- 6. Ventas (15 registros)
INSERT INTO Venta (usuario_id, obra_id, fecha_venta, monto, metodo_pago)
SELECT 
    (SELECT usuario_id FROM Usuario WHERE tipo_usuario = 'comprador' ORDER BY RANDOM() LIMIT 1),
    obra_id,
    CURRENT_DATE - (FLOOR(RANDOM() * 90) || ' days')::INTERVAL,
    precio_referencia * (0.8 + RANDOM() * 0.4), -- Monto +/- 20% del precio referencia
    CASE (RANDOM() * 2)::INT
        WHEN 0 THEN 'tarjeta'
        WHEN 1 THEN 'transferencia'
        ELSE 'paypal'
    END
FROM ObraArte 
WHERE estado = 'vendida' LIMIT 15;

-- 7. Subastas (5 registros)
INSERT INTO Subasta (obra_id, fecha_inicio, fecha_fin, monto_inicial)
SELECT 
    obra_id,
    CURRENT_TIMESTAMP - (FLOOR(RANDOM() * 30) || ' days')::INTERVAL,
    CURRENT_TIMESTAMP + (FLOOR(RANDOM() * 7) || ' days')::INTERVAL,
    precio_referencia * 0.5 -- Monto inicial al 50% del precio referencia
FROM ObraArte 
WHERE estado = 'subasta' LIMIT 5;

-- 8. Ofertas en Subastas (20 registros)
INSERT INTO OfertaSubasta (subasta_id, usuario_id, monto_ofertado)
WITH SubastasActivas AS (
    SELECT subasta_id, monto_inicial FROM Subasta
)
SELECT 
    subasta_id,
    (SELECT usuario_id FROM Usuario WHERE tipo_usuario = 'comprador' ORDER BY RANDOM() LIMIT 1),
    monto_inicial * (1 + (RANDOM() * 0.5)) -- Ofertas hasta +50% del monto inicial
FROM SubastasActivas
CROSS JOIN generate_series(1,4); -- 4 ofertas por subasta en promedio

-- 9. Envíos (15 registros)
INSERT INTO Envio (venta_id, direccion, fecha_envio, estado_envio)
SELECT 
    venta_id,
    'Calle ' || (RANDOM() * 1000)::INT || ', Ciudad ' || (RANDOM() * 10)::INT,
    fecha_venta + INTERVAL '1 day' + (FLOOR(RANDOM() * 5) || ' days')::INTERVAL,
    CASE (RANDOM() * 2)::INT
        WHEN 0 THEN 'pendiente'
        WHEN 1 THEN 'enviado'
        ELSE 'entregado'
    END
FROM Venta;