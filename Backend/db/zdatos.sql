-- 1)
INSERT INTO Usuario (nombre, correo, fecha_registro, tipo_usuario)
SELECT
  -- "Usuario 1", "Usuario 2", …
  'Usuario ' || i AS nombre,
  -- "user1@example.com", "user2@example.com", …
  'user' || i || '@example.com' AS correo,
  -- Fecha de registro: hoy menos (i-1)%30 días
  CURRENT_DATE - ((i - 1) % 30) * INTERVAL '1 day' AS fecha_registro,
  -- Rota los tipos: comprador, vendedor, admin
  CASE (i - 1) % 3
    WHEN 0 THEN 'comprador'
    WHEN 1 THEN 'vendedor'
    ELSE 'admin'
  END
FROM generate_series(1,30) AS s(i);

-- 2) PerfilArtista: un artista para los primeros 20 usuarios
INSERT INTO PerfilArtista (usuario_id, biografia, pais_origen, estilo_principal)
SELECT
  i,
  'Biografía del artista ' || i,
  CASE (i - 1) % 5
    WHEN 0 THEN 'España'
    WHEN 1 THEN 'México'
    WHEN 2 THEN 'Argentina'
    WHEN 3 THEN 'Colombia'
    ELSE 'Chile'
  END,
  CASE (i - 1) % 4
    WHEN 0 THEN 'Realismo'
    WHEN 1 THEN 'Impresionismo'
    WHEN 2 THEN 'Surrealismo'
    ELSE 'Abstracto'
  END
FROM generate_series(1,30) AS s(i);

-- 3) ObraArte: 50 obras distribuidas entre los 20 artistas
INSERT INTO ObraArte (titulo, descripcion, año_creacion, precio_referencia, estado, artista_id)
SELECT
  'Obra ' || i,
  'Descripción de la obra ' || i,
  1500 + ((i - 1) % 526),
  (i * 100)::NUMERIC(10,2),
  CASE (i - 1) % 4
    WHEN 0 THEN 'en venta'
    WHEN 1 THEN 'subasta'
    WHEN 2 THEN 'vendida'
    ELSE 'reservada'
  END,
  ((i - 1) % 20) + 1
FROM generate_series(1,30) AS s(i);

-- 4) Categoria: 10 categorías
INSERT INTO Categoria (nombre, descripcion)
SELECT
  'Categoría ' || i,
  'Descripción de la categoría ' || i
FROM generate_series(1,30) AS s(i);
