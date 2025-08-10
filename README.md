# Kubernetes Go CRUD API

A simple CRUD API built with Go and deployed on Kubernetes, demonstrating:
- Basic Go REST API with in-memory storage
- Kubernetes deployment configuration
- Docker containerization with security best practices
- Multiple service types (ClusterIP, NodePort)
- Ingress configuration

## Project Structure
```
.
├── Dockerfile           # Multi-stage build with security focus
├── main.go             # API entry point
├── go.mod              # Go module definition
├── pkg/
│   ├── api/           # API handlers
│   │   └── handlers.go
│   └── models/        # Data models
│       └── item.go
└── deployments/       # Kubernetes configurations
    ├── deployment.yaml
    ├── service.yaml
    ├── nodeport-service.yaml
    └── ingress.yaml
```

## Features
- CRUD operations for items
- Multiple ways to access the API:
  - NodePort: http://localhost:30000
  - Ingress: http://localhost/api
- Secure Docker image using distroless base
- Kubernetes resources with proper configuration

## Getting Started

### Prerequisites
- Go 1.21+
- Docker
- Kubernetes (via Colima/Docker Desktop/Minikube)
- kubectl
- Helm (for Ingress Controller)

### Running Locally
1. Start Kubernetes:
   ```bash
   colima start --kubernetes
   ```

2. Install NGINX Ingress Controller:
   ```bash
   helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
   helm repo update
   helm install ingress-nginx ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace
   ```

3. Build and deploy:
   ```bash
   docker build -t go-crud-api .
   kubectl apply -f deployments/
   ```

4. Access the API:
   - Via NodePort: http://localhost:30000/api/items
   - Via Ingress: http://localhost/api/items

## API Endpoints

- `GET /api/items` - List all items
- `GET /api/items/{id}` - Get a specific item
- `POST /api/items` - Create a new item
- `PUT /api/items/{id}` - Update an item
- `DELETE /api/items/{id}` - Delete an item

## Development

The project uses Tilt for development workflow:

```bash
tilt up
```

This will automatically:
- Build the Docker image
- Deploy to Kubernetes
- Set up port forwarding
- Watch for changes
