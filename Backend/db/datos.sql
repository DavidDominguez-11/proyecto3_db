--Datos Usuario
DO $$
BEGIN
    FOR i IN 1..100 LOOP
        INSERT INTO Usuario (nombre, correo, tipo_usuario)
        VALUES (
            'Usuario' || i,
            'usuario' || i || '@correo.com',
            CASE 
                WHEN i % 3 = 0 THEN 'admin'
                WHEN i % 2 = 0 THEN 'vendedor'
                ELSE 'comprador'
            END
        );
    END LOOP;
END $$;


--Datos PerfilArtista
DO $$
DECLARE
    uid INT;
BEGIN
    FOR uid IN SELECT usuario_id FROM Usuario WHERE tipo_usuario = 'vendedor' LIMIT 30 LOOP
        INSERT INTO PerfilArtista (usuario_id, biografia, pais_origen, estilo_principal)
        VALUES (uid, 'Biografía del artista ' || uid, 'País' || uid, 'Estilo' || uid);
    END LOOP;
END $$;


--Datos Categoria
DO $$
BEGIN
    FOR i IN 1..10 LOOP
        INSERT INTO Categoria (nombre, descripcion)
        VALUES ('Categoria' || i, 'Descripción de la categoría ' || i);
    END LOOP;
END $$;

--Datos ObraArte
DO $$
DECLARE
    aid INT;
    i INT := 1;
BEGIN
    FOR aid IN SELECT artista_id FROM PerfilArtista LOOP
        FOR j IN 1..4 LOOP  -- ajustamos para que 25 artistas * 4 = 100
            EXIT WHEN i > 100;
            INSERT INTO ObraArte (titulo, descripcion, año_creacion, precio_referencia, estado, artista_id)
            VALUES (
                'Obra ' || i,
                'Descripción de la obra ' || i,
                2000 + (i % 21),
                1000 + (i * 10),
                'en venta',
                aid
            );
            i := i + 1;
        END LOOP;
        EXIT WHEN i > 100;
    END LOOP;
END $$;


--Datos ObraCategoria
DO $$
DECLARE
    oid INT;
BEGIN
    FOR oid IN SELECT obra_id FROM ObraArte LOOP
        INSERT INTO ObraCategoria (obra_id, categoria_id)
        VALUES (oid, ((oid - 1) % 10) + 1);
    END LOOP;
END $$;


--Datos Venta
DO $$
DECLARE
    cid INT;
    oid INT;
    i INT := 1;
BEGIN
    FOR cid IN SELECT usuario_id FROM Usuario WHERE tipo_usuario = 'comprador' LIMIT 40 LOOP
        SELECT obra_id INTO oid FROM ObraArte WHERE estado = 'en venta' LIMIT 1;
        INSERT INTO Venta (usuario_id, obra_id, fecha_venta, monto, metodo_pago)
        VALUES (
            cid,
            oid,
            CURRENT_DATE - (i || ' days')::interval,
            1200 + i * 5,
            CASE WHEN i % 3 = 0 THEN 'paypal' WHEN i % 2 = 0 THEN 'transferencia' ELSE 'tarjeta' END
        );
        i := i + 1;
    END LOOP;
END $$;


--Datos Envio
DO $$
DECLARE
    vid INT;
    estado TEXT;
BEGIN
    FOR vid IN SELECT venta_id FROM Venta LOOP
        -- Determinar el estado aleatoriamente
        IF random() < 0.4 THEN
            estado := 'enviado';
        ELSE
            estado := 'pendiente';
        END IF;

        INSERT INTO Envio (venta_id, direccion, fecha_envio, estado_envio)
        VALUES (
            vid,
            'Dirección de envío para venta ' || vid,
            CURRENT_DATE + INTERVAL '1 day',
            estado
        );
    END LOOP;
END $$;


--Datos Subasta
DO $$
DECLARE
    oid INT;
BEGIN
    FOR oid IN SELECT obra_id FROM ObraArte WHERE estado != 'vendida' LIMIT 30 LOOP
        INSERT INTO Subasta (obra_id, fecha_inicio, fecha_fin, monto_inicial)
        VALUES (
            oid,
            CURRENT_TIMESTAMP,
            CURRENT_TIMESTAMP + INTERVAL '1 day',
            500 + oid
        );
    END LOOP;
END $$;


--Datos OfertaSubasta
DO $$
DECLARE
    sid INT;
    uid INT;
    monto NUMERIC := 600;
BEGIN
    FOR sid IN SELECT subasta_id FROM Subasta LOOP
        FOR uid IN SELECT usuario_id FROM Usuario WHERE tipo_usuario = 'comprador' LIMIT 3 LOOP
            monto := monto + 10;
            INSERT INTO OfertaSubasta (subasta_id, usuario_id, monto_ofertado)
            VALUES (sid, uid, monto);
        END LOOP;
    END LOOP;
END $$;