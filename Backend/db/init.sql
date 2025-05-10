-- Crear la base de datos solo si no existe
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'p3db') THEN
        CREATE DATABASE p3db;
    END IF;
END $$;

-- Crear el usuario solo si no existe
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'dbuser') THEN
        CREATE USER dbuser WITH PASSWORD 'dbpassword';
    END IF;
END $$;

-- Otorgar privilegios al usuario sobre la base de datos
GRANT ALL PRIVILEGES ON DATABASE p3db TO dbuser;

-- Script DDL
-- Tabla Usuario
CREATE TABLE IF NOT EXISTS Usuario (
    usuario_id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    correo VARCHAR(100) UNIQUE NOT NULL,
    fecha_registro DATE DEFAULT CURRENT_DATE,
    tipo_usuario VARCHAR(15) NOT NULL CHECK (tipo_usuario IN ('comprador', 'vendedor', 'admin'))
);

-- Tabla PerfilArtista
CREATE TABLE IF NOT EXISTS PerfilArtista (
    artista_id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL UNIQUE,
    biografia TEXT,
    pais_origen VARCHAR(50),
    estilo_principal VARCHAR(50),
    FOREIGN KEY (usuario_id) REFERENCES Usuario(usuario_id)
);

-- Tabla ObraArte
CREATE TABLE IF NOT EXISTS ObraArte (
    obra_id SERIAL PRIMARY KEY,
    titulo VARCHAR(100) NOT NULL,
    descripcion TEXT,
    año_creacion INT CHECK (año_creacion > 1000),
    precio_referencia NUMERIC(10,2) DEFAULT 0,
    estado VARCHAR(20) NOT NULL CHECK (estado IN ('en venta', 'subasta', 'vendida', 'reservada')),
    artista_id INT NOT NULL,
    FOREIGN KEY (artista_id) REFERENCES PerfilArtista(artista_id)
);

-- Tabla Categoria
CREATE TABLE IF NOT EXISTS Categoria (
    categoria_id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) UNIQUE NOT NULL,
    descripcion TEXT
);

-- Tabla de cruce ObraCategoria (atributo multivaluado)
CREATE TABLE IF NOT EXISTS ObraCategoria (
    obra_id INT NOT NULL,
    categoria_id INT NOT NULL,
    PRIMARY KEY (obra_id, categoria_id),
    FOREIGN KEY (obra_id) REFERENCES ObraArte(obra_id),
    FOREIGN KEY (categoria_id) REFERENCES Categoria(categoria_id)
);

-- Tabla Venta
CREATE TABLE IF NOT EXISTS Venta (
    venta_id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL, -- comprador
    obra_id INT NOT NULL,
    fecha_venta DATE NOT NULL,
    monto NUMERIC(10,2) NOT NULL CHECK (monto > 0),
    metodo_pago VARCHAR(20) NOT NULL CHECK (metodo_pago IN ('tarjeta', 'transferencia', 'paypal')),
    FOREIGN KEY (usuario_id) REFERENCES Usuario(usuario_id),
    FOREIGN KEY (obra_id) REFERENCES ObraArte(obra_id)
);

-- Tabla Subasta
CREATE TABLE IF NOT EXISTS Subasta (
    subasta_id SERIAL PRIMARY KEY,
    obra_id INT NOT NULL UNIQUE,
    fecha_inicio TIMESTAMP NOT NULL,
    fecha_fin TIMESTAMP NOT NULL,
    monto_inicial NUMERIC(10,2) NOT NULL CHECK (monto_inicial > 0),
    FOREIGN KEY (obra_id) REFERENCES ObraArte(obra_id)
);

-- Tabla OfertaSubasta
CREATE TABLE IF NOT EXISTS OfertaSubasta (
    oferta_id SERIAL PRIMARY KEY,
    subasta_id INT NOT NULL,
    usuario_id INT NOT NULL,
    monto_ofertado NUMERIC(10,2) NOT NULL CHECK (monto_ofertado > 0),
    fecha_oferta TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (subasta_id) REFERENCES Subasta(subasta_id),
    FOREIGN KEY (usuario_id) REFERENCES Usuario(usuario_id)
);

-- Tabla Envio
CREATE TABLE IF NOT EXISTS Envio (
    envio_id SERIAL PRIMARY KEY,
    venta_id INT NOT NULL UNIQUE,
    direccion TEXT NOT NULL,
    fecha_envio DATE,
    estado_envio VARCHAR(20) CHECK (estado_envio IN ('pendiente', 'enviado', 'entregado')),
    FOREIGN KEY (venta_id) REFERENCES Venta(venta_id)
);

-- Tabla Transaccion (para registrar eventos automáticos)
CREATE TABLE IF NOT EXISTS Transaccion (
    transaccion_id SERIAL PRIMARY KEY,
    tipo VARCHAR(30) NOT NULL CHECK (tipo IN ('venta', 'subasta_ganada', 'oferta', 'cambio_estado')),
    detalle TEXT NOT NULL,
    fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    entidad_afectada_id INT NOT NULL -- puede hacer referencia a obra_id, venta_id, etc. (flexible)
);

-- TRIGGERS

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
