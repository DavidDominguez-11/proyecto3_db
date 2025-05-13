# Proyecto3 DB - Sistema de Gestión de Galería de Arte
Plataforma web para gestión de obras de arte, ventas, subastas y reportes analíticos mediante filtros.

## Grupo 14
Anggie Quezada 23643  
Gabriel Bran 23590  
David Dominguez 23712  

## Características Principales

- 🖼️ CRUD completo para Artistas, Obras y Usuarios
- 📊 Reportes avanzados con filtros:
  - Ventas realizadas
  - Ofertas por subasta
  - Transacciones del sistema
  - Obras por categoría y popularidad
- 🛒 Sistema de subastas con ofertas en tiempo real
- 📦 Seguimiento de envíos
- 📈 Análisis de tendencias artísticas

## Tecnologías

**Backend:**
- Go 1.24
- PostgreSQL 15
- Docker

**Frontend:**
- HTML5/CSS3
- JavaScript ES6
- Python HTTP Server (para desarrollo)

**Infraestructura:**
- Docker Compose
- PostgreSQL con inicialización automática
- Configuración CORS

## Instalación

1. **Prerrequisitos:**
   - Docker 20.10+
   - Docker Compose 2.20+

2. **Clonar y ejecutar:**
```bash
git clone https://github.com/DavidDominguez-11/proyecto3_db.git
cd proyecto3_db
docker-compose up --build
ir a: http://localhost:8000/ o http://127.0.0.1:8000/
```
En la pantalla principal, elige el filtro deseado. Esto dirige a la pantalla del filtro.  
Agrega los filtros deseados. Cuando alguno sea requerido al momento de filtrar mediante el botón "Filtrar", saldrá una advertencia si falta alguno o si algún valor no es correcto.  
Para ver todos los datos, filtra sin parámetros y solo agrega el necesario si se requiere.  

3. **Parar y ejecutar de nuevo:**
```bash
ctrl + c
docker-compose down -v
docker-compose up --build
```
