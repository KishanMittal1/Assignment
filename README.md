# Multi-Service Reverse Proxy with Docker Compose

## Overview

This project demonstrates a simple multi-service Docker Compose setup with an Nginx reverse proxy routing to:

- **service_1**: Golang backend (`/service1/`)
- **service_2**: Python backend (`/service2/`)

All services are accessible via:  
`http://localhost:8080/service1/` and  
`http://localhost:8080/service2/`

---

## Setup Instructions

1. **Clone this repository**  
   Make sure `service_1` and `service_2` contain the code as described in their respective folders.

2. **Build and start all services**
   ```bash
   docker-compose up --build
   ```

3. **Test routing**
   - Open [http://localhost:8080/service1/](http://localhost:8080/service1/) for Go service
   - Open [http://localhost:8080/service2/](http://localhost:8080/service2/) for Python service

---

## Routing

Nginx reverse proxy uses path-based routing:
- `/service1/` → `service1` container (`port 8000`)
- `/service2/` → `service2` container (`port 8000`)

See [`nginx/nginx.conf`](nginx/nginx.conf) for details.

---

## Logging

Nginx logs all incoming requests with timestamp and path to `/var/log/nginx/access.log` inside the container.

---

## Health Checks

Docker Compose uses HTTP health checks:
- `service1`: `http://localhost:8000/health`
- `service2`: `http://localhost:8000/health`

Check container health with:
```bash
docker ps
```

---

## Bonus

- Health checks for both backend services.
- Custom access log format for clarity.
- Modular, clean Docker setup.

---

## Troubleshooting

- Ensure ports aren't in use.
- Rebuild with `docker-compose build --no-cache` if changes aren't reflected.

---

## License

MIT