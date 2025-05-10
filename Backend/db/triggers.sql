-- Insertar en Transaccion cuando se realiza una venta
CREATE OR REPLACE FUNCTION log_transaccion_venta()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO Transaccion (tipo, detalle, entidad_afectada_id)
    VALUES ('venta',
            'Venta registrada para la obra ID ' || NEW.obra_id || ' por usuario ID ' || NEW.usuario_id,
            NEW.venta_id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_log_transaccion_venta
AFTER INSERT ON Venta
FOR EACH ROW
EXECUTE FUNCTION log_transaccion_venta();

-- Cambiar estado de obra a "vendida" cuando se realiza una venta
CREATE OR REPLACE FUNCTION actualizar_estado_obra_vendida()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE ObraArte
    SET estado = 'vendida'
    WHERE obra_id = NEW.obra_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_actualizar_estado_obra
AFTER INSERT ON Venta
FOR EACH ROW
EXECUTE FUNCTION actualizar_estado_obra_vendida();

-- Validar que la oferta sea mayor a las anteriores en una subasta
CREATE OR REPLACE FUNCTION validar_oferta_subasta()
RETURNS TRIGGER AS $$
DECLARE
    max_monto NUMERIC;
BEGIN
    SELECT MAX(monto_ofertado) INTO max_monto
    FROM OfertaSubasta
    WHERE subasta_id = NEW.subasta_id;

    IF max_monto IS NOT NULL AND NEW.monto_ofertado <= max_monto THEN
        RAISE EXCEPTION 'La oferta debe ser mayor que la oferta actual más alta (%.2f)', max_monto;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_validar_oferta_subasta
BEFORE INSERT ON OfertaSubasta
FOR EACH ROW
EXECUTE FUNCTION validar_oferta_subasta();

-- Validar que fecha de fin de subasta sea posterior a la de inicio
CREATE OR REPLACE FUNCTION validar_fechas_subasta()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.fecha_fin <= NEW.fecha_inicio THEN
        RAISE EXCEPTION 'La fecha de fin de la subasta debe ser posterior a la de inicio.';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_validar_fechas_subasta
BEFORE INSERT ON Subasta
FOR EACH ROW
EXECUTE FUNCTION validar_fechas_subasta();
