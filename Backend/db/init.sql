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


