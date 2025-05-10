// models/datos.go
package models

import "time"

type TipoUsuario string

const (
	Comprador TipoUsuario = "comprador"
	Vendedor  TipoUsuario = "vendedor"
	Admin     TipoUsuario = "admin"
)

type Usuario struct {
	ID            int        `json:"usuario_id"`
	Nombre        string     `json:"nombre"`
	Correo        string     `json:"correo"`
	FechaRegistro time.Time  `json:"fecha_registro"`
	TipoUsuario   TipoUsuario `json:"tipo_usuario"`
}

// models/artist.go
type PerfilArtista struct {
	ID             int    `json:"artista_id"`
	UsuarioID      int    `json:"usuario_id"`
	Biografia      string `json:"biografia"`
	PaisOrigen     string `json:"pais_origen"`
	EstiloPrincipal string `json:"estilo_principal"`
}

// models/artwork.go
type EstadoObra string

const (
	EnVenta   EstadoObra = "en venta"
	Subasta   EstadoObra = "subasta"
	Vendida   EstadoObra = "vendida"
	Reservada EstadoObra = "reservada"
)

type ObraArte struct {
	ID             int       `json:"obra_id"`
	Titulo         string    `json:"titulo"`
	Descripcion    string    `json:"descripcion"`
	AñoCreacion    int       `json:"año_creacion"`
	PrecioReferencia float64  `json:"precio_referencia"`
	Estado         EstadoObra `json:"estado"`
	ArtistaID      int       `json:"artista_id"`
	Categorias     []Categoria `json:"categorias,omitempty"`
}

// models/category.go
type Categoria struct {
	ID          int    `json:"categoria_id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

// models/sale.go
type MetodoPago string

const (
	Tarjeta       MetodoPago = "tarjeta"
	Transferencia MetodoPago = "transferencia"
	Paypal        MetodoPago = "paypal"
)

type Venta struct {
	ID          int        `json:"venta_id"`
	UsuarioID   int        `json:"usuario_id"`
	ObraID      int        `json:"obra_id"`
	FechaVenta  time.Time  `json:"fecha_venta"`
	Monto       float64    `json:"monto"`
	MetodoPago  MetodoPago `json:"metodo_pago"`
}

// models/auction.go
type Subasta_ struct {
	ID           int       `json:"subasta_id"`
	ObraID       int       `json:"obra_id"`
	FechaInicio  time.Time `json:"fecha_inicio"`
	FechaFin     time.Time `json:"fecha_fin"`
	MontoInicial float64   `json:"monto_inicial"`
	Ofertas      []OfertaSubasta `json:"ofertas,omitempty"`
}

type OfertaSubasta struct {
	ID           int       `json:"oferta_id"`
	SubastaID    int       `json:"subasta_id"`
	UsuarioID    int       `json:"usuario_id"`
	MontoOfertado float64   `json:"monto_ofertado"`
	FechaOferta  time.Time `json:"fecha_oferta"`
}

// models/shipping.go
type EstadoEnvio string

const (
	Pendiente EstadoEnvio = "pendiente"
	Enviado   EstadoEnvio = "enviado"
	Entregado EstadoEnvio = "entregado"
)

type Envio struct {
	ID          int         `json:"envio_id"`
	VentaID     int         `json:"venta_id"`
	Direccion   string      `json:"direccion"`
	FechaEnvio  *time.Time  `json:"fecha_envio,omitempty"`
	EstadoEnvio EstadoEnvio `json:"estado_envio"`
}

// models/transaction.go
type TipoTransaccion string

const (
	VentaTransaccion    TipoTransaccion = "venta"
	SubastaGanada       TipoTransaccion = "subasta_ganada"
	OfertaTransaccion   TipoTransaccion = "oferta"
	CambioEstado        TipoTransaccion = "cambio_estado"
)

type Transaccion struct {
	ID               int            `json:"transaccion_id"`
	Tipo             TipoTransaccion `json:"tipo"`
	Detalle          string         `json:"detalle"`
	Fecha            time.Time      `json:"fecha"`
	EntidadAfectadaID int           `json:"entidad_afectada_id"`
}