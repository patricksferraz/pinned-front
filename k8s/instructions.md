# Kubernetes Deployment Guide

This guide provides instructions for deploying the Pinned Front application to a Kubernetes cluster.

## Prerequisites

- Kubernetes cluster up and running
- `kubectl` configured and connected to your cluster
- Docker registry access (if using private images)
- Basic understanding of Kubernetes concepts

## Configuration

### 1. Environment Variables

1. Navigate to the `k8s` directory:
```bash
cd k8s
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

3. Edit the `.env` file with your specific configuration values:
```bash
# Example of required variables
APP_ENV=production
APP_PORT=8080
# Add other required environment variables
```

### 2. Creating Kubernetes Secrets

#### Application Secrets

Create a generic secret containing your application's environment variables:

```bash
kubectl create secret generic pinned-front-secrets \
  --from-env-file k8s/.env \
  --namespace your-namespace
```

#### Docker Registry Secrets

If you're using a private Docker registry, create a registry secret:

```bash
kubectl create secret docker-registry regsecret \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL \
  --namespace your-namespace
```

Required variables:
- `$DOCKER_REGISTRY_SERVER`: Docker registry URL (e.g., docker.io, gcr.io)
- `$DOCKER_USER`: Registry username
- `$DOCKER_PASSWORD`: Registry password
- `$DOCKER_EMAIL`: (Optional) Email associated with the registry account

## Deployment

### 1. Deploy All Resources

To deploy all Kubernetes resources:

```bash
kubectl apply -f ./k8s
```

This will create:
- Deployments
- Services
- ConfigMaps
- Ingress rules (if configured)
- Other Kubernetes resources

### 2. Verify Deployment

Check the status of your deployment:

```bash
# Check pods
kubectl get pods -n your-namespace

# Check services
kubectl get services -n your-namespace

# Check deployments
kubectl get deployments -n your-namespace
```

### 3. View Logs

To view application logs:

```bash
# Get pod name
kubectl get pods -n your-namespace

# View logs
kubectl logs -f <pod-name> -n your-namespace
```

## Maintenance

### Updating the Deployment

1. Update your application code
2. Build and push the new Docker image
3. Update the deployment:
```bash
kubectl set image deployment/pinned-front pinned-front=new-image:tag -n your-namespace
```

### Scaling

Scale your deployment:

```bash
kubectl scale deployment pinned-front --replicas=3 -n your-namespace
```

### Troubleshooting

Common issues and solutions:

1. **Pods in CrashLoopBackOff**:
```bash
kubectl describe pod <pod-name> -n your-namespace
kubectl logs <pod-name> -n your-namespace
```

2. **Service not accessible**:
```bash
kubectl describe service pinned-front -n your-namespace
```

3. **Secret issues**:
```bash
kubectl describe secret pinned-front-secrets -n your-namespace
```

## Cleanup

To remove all deployed resources:

```bash
kubectl delete -f ./k8s
```

To remove specific secrets:

```bash
kubectl delete secret pinned-front-secrets -n your-namespace
kubectl delete secret regsecret -n your-namespace
```

## Best Practices

1. Always use namespaces to isolate resources
2. Implement resource limits and requests
3. Use health checks (readiness/liveness probes)
4. Implement proper logging and monitoring
5. Use secrets for sensitive data
6. Implement proper backup strategies
7. Use version control for your Kubernetes manifests

## Additional Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Kubernetes Best Practices](https://kubernetes.io/docs/concepts/configuration/overview/)
