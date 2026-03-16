# Kubernetes Gateway API

Roteamento de tráfego HTTP com **NGINX Gateway Fabric** usando a Gateway API do Kubernetes.

Suporta roteamento por **path** e por **hostname**:

| URL | Tipo |
|-----|------|
| `http://<dominio>/blue` | Path-based |
| `http://<dominio>/green` | Path-based |
| `http://blue.<dominio>` | Host-based |
| `http://green.<dominio>` | Host-based |

---

## Estrutura

```
.
├── kind-config.yaml    # Configuração do cluster Kind (ambiente local)
├── values.yaml         # Configuração do NGINX Gateway Fabric (Helm)
├── app.yaml            # Deployments, Services, Gateway e HTTPRoutes
└── host-routes.yaml    # HTTPRoutes por hostname (opcional no ambiente local)
```

---

## Pré-requisitos

- `kubectl`
- `helm`

---

## Ambiente Local (Kind)

### 1. Criar o cluster

```bash
kind create cluster --name kindcluster --config kind-config.yaml
```

> O `kind-config.yaml` mapeia `containerPort: 30000` → `hostPort: 80` no control-plane.

### 2. Instalar os CRDs da Gateway API

```bash
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.2.0/standard-install.yaml
```

### 3. Instalar o NGINX Gateway Fabric

```bash
helm upgrade --install ngf oci://ghcr.io/nginxinc/charts/nginx-gateway-fabric \
  --namespace nginx-gateway \
  --create-namespace \
  -f values.yaml
```

> O `values.yaml` configura `NodePort 30000`, `nodeSelector` e `toleration` para o pod rodar no control-plane.

### 4. Aplicar os recursos

```bash
kubectl apply -f app.yaml
```

### 5. Testar

```bash
curl http://localhost/blue
curl http://localhost/green
```

Para roteamento por hostname, use [nip.io](https://nip.io) (não requer alteração no `/etc/hosts`):

```bash
curl http://blue.127.0.0.1.nip.io
curl http://green.127.0.0.1.nip.io
```

---

## Nuvem (qualquer cloud provider)

### 1. Criar o cluster

Crie um cluster Kubernetes utilizando o cloud provider de sua escolha (AKS, EKS, GKE, etc.).

### 2. Instalar os CRDs da Gateway API

```bash
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.2.0/standard-install.yaml
```

### 3. Instalar o NGINX Gateway Fabric

```yaml
# values.yaml (nuvem — muito mais simples que o local)
service:
  type: LoadBalancer
```

```bash
helm upgrade --install ngf oci://ghcr.io/nginxinc/charts/nginx-gateway-fabric \
  --namespace nginx-gateway \
  --create-namespace \
  -f values.yaml
```

Aguarde o IP público ser provisionado:

```bash
kubectl get svc -n nginx-gateway --watch
# Aguarde aparecer um EXTERNAL-IP
```

### 4. Configurar o DNS

Crie registros **A** no seu provedor de DNS apontando para o `EXTERNAL-IP`:

| Tipo | Host | Valor |
|------|------|-------|
| A | `@` | `<EXTERNAL-IP>` |
| A | `blue` | `<EXTERNAL-IP>` |
| A | `green` | `<EXTERNAL-IP>` |

### 5. Aplicar os recursos

```bash
kubectl apply -f app.yaml
```

### 6. Testar

```bash
curl http://<dominio>/blue
curl http://<dominio>/green
curl http://blue.<dominio>
curl http://green.<dominio>
```

---

## Como funciona

### Gateway com dois listeners

O Gateway precisa de dois listeners separados porque `*.dominio.com` **não cobre** o domínio raiz `dominio.com`:

```yaml
listeners:
  - name: http-subdomains
    protocol: HTTP
    port: 80
    hostname: "*.dominio.com"    # cobre blue.dominio.com e green.dominio.com

  - name: http-root
    protocol: HTTP
    port: 80
    hostname: "dominio.com"      # cobre dominio.com/blue e dominio.com/green
```

### HTTPRoutes com sectionName

Cada HTTPRoute referencia explicitamente seu listener via `sectionName`:

```yaml
# Rota por path → listener http-root
parentRefs:
  - name: main-gateway
    sectionName: http-root

# Rota por hostname → listener http-subdomains
parentRefs:
  - name: main-gateway
    sectionName: http-subdomains
```

### URLRewrite nas rotas por path

Sem o filtro `URLRewrite`, a aplicação receberia `/blue` como path e retornaria 404. O filtro substitui `/blue` por `/` antes de encaminhar:

```yaml
filters:
  - type: URLRewrite
    urlRewrite:
      path:
        type: ReplacePrefixMatch
        replacePrefixMatch: /
```

### Fluxo do tráfego

```
Usuário
  └─► DNS
        └─► Load Balancer (IP público)
              └─► NGINX Gateway Fabric
                    └─► HTTPRoute (path ou hostname)
                          └─► Service
                                └─► Pod (blue ou green)
```

---

## Diferenças: Local vs Nuvem

| | Kind (local) | Nuvem |
|--|--|--|
| Exposição | NodePort 30000 | LoadBalancer (IP público automático) |
| nodeSelector | Obrigatório | Não necessário |
| Toleration | Obrigatório | Não necessário |
| DNS | nip.io | DNS real |
| values.yaml | Complexo | Simples |
