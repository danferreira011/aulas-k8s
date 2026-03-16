# **SUPER GUIA KUBERNETES – DEVOPS EDITION**

### **Guia de consulta para projetos DevOps**

Este guia reúne os **principais conceitos usados em ambientes Kubernetes profissionais**, organizados por categoria.

---

# **1\. Conceitos Fundamentais**

## **Cluster**

Conjunto de máquinas (nodes) que executam workloads Kubernetes.

Componentes principais:

* Control Plane  
* Worker Nodes

---

## **Node**

Máquina que executa os pods.

Pode ser:

* VM  
* servidor físico  
* instância cloud

Componente principal:

kubelet

---

## **Pod**

Menor unidade executável do Kubernetes.

Características:

* pode conter 1 ou mais containers  
* compartilham rede  
* compartilham storage

---

## **Container**

Aplicação empacotada com suas dependências.

Normalmente criada com:

Docker

---

# **2\. Workloads**

## **Deployment**

Gerencia aplicações **stateless**.

Funções:

* cria pods  
* atualiza pods  
* controla réplicas

Exemplo:

replicas: 3

---

## **ReplicaSet**

Mantém um número desejado de pods rodando.

Normalmente gerenciado pelo **Deployment**.

---

## **StatefulSet**

Usado para aplicações **stateful**.

Características:

* identidade persistente  
* storage persistente  
* ordem de criação

Usado em:

* bancos  
* clusters distribuídos

---

## **DaemonSet**

Garante que **um pod rode em todos os nodes**.

Muito usado para:

* monitoramento  
* logs  
* segurança

---

## **Job**

Executa tarefas **uma única vez**.

Exemplo:

* processamento batch  
* migração de banco

---

## **CronJob**

Executa tarefas **agendadas**.

Exemplo:

backup diário

---

# **3\. Configuração de Aplicações**

## **ConfigMap**

Armazena **configurações não sensíveis**.

Exemplo:

variáveis de ambiente  
arquivos de configuração

---

## **Secret**

Armazena **dados sensíveis**.

Exemplo:

senhas  
tokens  
certificados

---

## **Environment Variables**

Permite passar configurações para containers.

Exemplo:

env:  
  \- name: DATABASE\_URL

---

# **4\. Networking**

## **Service**

Permite comunicação entre pods.

Tipos principais:

---

### **ClusterIP**

Acesso interno no cluster.

---

### **NodePort**

Exposição via porta do node.

---

### **LoadBalancer**

Cria um load balancer externo na cloud.

---

### **ExternalName**

Aponta para serviço externo.

---

## **Ingress**

Gerencia **acesso HTTP/HTTPS externo**.

Permite:

* múltiplos domínios  
* TLS  
* roteamento por path

---

## **Gateway API**

Evolução do Ingress.

Permite:

* melhor controle de tráfego  
* arquitetura mais modular

---

## **Endpoint**

Representa os **pods atrás de um service**.

---

## **EndpointSlice**

Versão moderna e escalável do endpoint.

---

# **5\. Storage**

## **Volume**

Armazenamento dentro do pod.

Exemplo:

emptyDir

---

## **Persistent Volume (PV)**

Representa **storage físico no cluster**.

Pode vir de:

* disco cloud  
* NFS  
* storage local

---

## **Persistent Volume Claim (PVC)**

Pedido de armazenamento feito por um pod.

---

## **StorageClass**

Define **tipos de storage disponíveis**.

Permite provisionamento dinâmico.

---

# **6\. Saúde das Aplicações**

## **Liveness Probe**

Verifica se o container está saudável.

Se falhar:

container é reiniciado

---

## **Readiness Probe**

Verifica se o pod está pronto para receber tráfego.

---

## **Startup Probe**

Usado para aplicações que demoram para iniciar.

---

# **7\. Escalabilidade**

## **Horizontal Pod Autoscaler (HPA)**

Escala número de pods baseado em métricas.

Exemplo:

CPU \> 70%

---

## **Vertical Pod Autoscaler (VPA)**

Ajusta recursos de containers.

Exemplo:

mais CPU  
mais memória

---

## **Cluster Autoscaler**

Escala número de nodes.

Muito usado em cloud.

---

# **8\. Gerenciamento de Recursos**

## **Requests**

Quantidade mínima garantida de recursos.

---

## **Limits**

Quantidade máxima permitida.

---

## **QoS (Quality of Service)**

Define prioridade de pods.

Tipos:

* Guaranteed  
* Burstable  
* BestEffort

---

## **LimitRange**

Define limites padrão em um namespace.

---

## **ResourceQuota**

Limita consumo total de recursos.

---

# **9\. Organização do Cluster**

## **Namespace**

Permite separar recursos.

Exemplo:

dev  
staging  
prod

---

## **Labels**

Pares chave/valor para organização.

Exemplo:

app=frontend  
env=production

---

## **Selectors**

Usados para selecionar recursos com base em labels.

---

# **10\. Scheduling Avançado**

## **Node Selector**

Define em qual node o pod pode rodar.

---

## **Node Affinity**

Agendamento mais avançado baseado em regras.

---

## **Pod Affinity**

Agrupa pods próximos.

---

## **Pod Anti-Affinity**

Evita pods no mesmo node.

---

## **Taints**

Impede pods de rodar em nodes.

---

## **Tolerations**

Permite pods ignorarem taints.

---

# **11\. Observabilidade**

Ferramentas comuns:

| Ferramenta | Função |
| ----- | ----- |
| Prometheus | métricas |
| Grafana | dashboards |
| ELK | logs |
| Jaeger | tracing |

---

# **12\. Segurança**

## **RBAC**

Controle de acesso baseado em papéis.

---

## **Service Account**

Identidade usada por pods.

---

## **Network Policies**

Firewall interno do cluster.

---

# **13\. Deploy Strategies**

## **Rolling Update**

Atualiza pods gradualmente.

---

## **Blue-Green**

Duas versões da aplicação.

---

## **Canary**

Libera nova versão para poucos usuários.

---

# **14\. GitOps**

Prática moderna de deploy.

Ferramentas:

* ArgoCD  
* FluxCD

Fluxo:

Git  
↓  
ArgoCD  
↓  
Kubernetes

---

# **15\. Ferramentas DevOps Comuns**

| Ferramenta | Função |
| ----- | ----- |
| Helm | gerenciamento de charts |
| ArgoCD | GitOps |
| Kustomize | customização YAML |
| Istio | service mesh |
| Linkerd | service mesh |

---

# **16\. Comandos Essenciais kubectl**

Ver pods:

kubectl get pods

Ver logs:

kubectl logs POD

Entrar no container:

kubectl exec \-it POD \-- bash

Aplicar manifesto:

kubectl apply \-f file.yaml

Deletar recurso:

kubectl delete pod POD

---

# **Fluxo de Arquitetura Kubernetes em Produção**

Internet  
   ↓  
Load Balancer  
   ↓  
Ingress / Gateway  
   ↓  
Service  
   ↓  
Pods  
   ↓  
Persistent Storage

---

# **Como usar este guia nos seus projetos**

Sugestão:

* salvar como **Markdown**  
* manter no GitHub  
* consultar durante deploys

