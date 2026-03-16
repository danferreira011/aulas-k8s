# **📌 Por que precisamos usar `ldd`**

Quando um programa é executado no Linux (como o **Node.js**), ele normalmente depende de **bibliotecas compartilhadas** do sistema.

Essas bibliotecas ficam geralmente em:

* `/lib`  
* `/usr/lib`  
* `/usr/local/lib`

Se você usar `scratch`, **essas bibliotecas não existem**.  
Então o binário quebra.

O comando **`ldd`** mostra exatamente **quais bibliotecas o binário precisa**.

---

# **🧠 Como funciona o `ldd`**

Execute dentro de um container Node ou em uma máquina Linux com Node instalado:

ldd $(which node)

### **Explicando o comando**

`which node`

Mostra o caminho do executável:

/usr/local/bin/node

Então o comando vira:

ldd /usr/local/bin/node

---

# **📊 Exemplo de saída**

Você verá algo parecido com:

linux-vdso.so.1 (0x00007ffd3b9d9000)  
libnode.so.115 \=\> /usr/local/lib/libnode.so.115  
libc.musl-x86\_64.so.1 \=\> /lib/libc.musl-x86\_64.so.1  
libstdc++.so.6 \=\> /usr/lib/libstdc++.so.6  
libgcc\_s.so.1 \=\> /usr/lib/libgcc\_s.so.1

Agora vamos entender isso.

---

# **🔎 Interpretando o resultado**

### **Exemplo**

libstdc++.so.6 \=\> /usr/lib/libstdc++.so.6

Significa:

* Node precisa da biblioteca **libstdc++**  
* Ela está em **/usr/lib**

Então se você usar `scratch`, precisa copiar:

COPY \--from=builder /usr/lib/libstdc++.so.6 /usr/lib/

---

Outro exemplo:

libc.musl-x86\_64.so.1 \=\> /lib/libc.musl-x86\_64.so.1

Então precisamos copiar:

COPY \--from=builder /lib/libc.musl-x86\_64.so.1 /lib/

---

# **📌 Dependências comuns do Node**

Normalmente você verá:

| Biblioteca | Função |
| ----- | ----- |
| libc | biblioteca padrão C |
| libstdc++ | C++ runtime |
| libgcc | suporte GCC |
| libcrypto | criptografia |
| libssl | TLS |

---

# **⚙️ Processo DevOps completo**

## **1️⃣ Rodar container temporário**

docker run \-it node:20-alpine sh

---

## **2️⃣ Descobrir onde está o Node**

which node

Saída:

/usr/local/bin/node

---

## **3️⃣ Ver dependências**

ldd /usr/local/bin/node

Exemplo:

libstdc++.so.6 \=\> /usr/lib/libstdc++.so.6  
libgcc\_s.so.1 \=\> /usr/lib/libgcc\_s.so.1  
libcrypto.so.3 \=\> /lib/libcrypto.so.3  
libssl.so.3 \=\> /lib/libssl.so.3

---

# **🧱 Agora sabemos o que copiar**

Dockerfile:

FROM node:20-alpine AS builder

WORKDIR /app

COPY package\*.json .  
RUN npm install \--omit=dev

COPY . .

FROM scratch

WORKDIR /app

COPY \--from=builder /usr/local/bin/node /usr/local/bin/node  
COPY \--from=builder /usr/lib /usr/lib  
COPY \--from=builder /lib /lib  
COPY \--from=builder /app /app

CMD \["/usr/local/bin/node", "server.js"\]

---

# **🚨 Erro clássico quando você esquece libs**

Quando falta biblioteca aparece algo como:

standard\_init\_linux.go:219: exec user process caused: no such file or directory

ou

error while loading shared libraries

Esse erro **quase sempre significa dependência faltando**.

---

# **🔬 Ferramenta melhor ainda: `scanelf`**

Em Alpine:

scanelf \-n /usr/local/bin/node

Isso mostra dependências de forma mais detalhada.

---

# **💡 Dica DevOps muito usada**

Quando alguém quer entender dependências de um binário:

ldd binary

Quando quer entender **qual pacote instalou a biblioteca**:

apk info \-W /usr/lib/libstdc++.so.6

---

# **🚀 Conceito muito importante (nível DevOps Sênior)**

Existem dois tipos de binários:

### **🔹 Dinâmicos**

Dependem de libs do sistema

Exemplo:

* Node  
* Python  
* Java

---

### **🔹 Estáticos**

Não dependem de libs

Exemplo:

* Go  
* Rust

Por isso Go funciona perfeitamente com **scratch**.

---

# **🔥 Exemplo Go (perfeito para scratch)**

FROM golang:1.22 AS builder

WORKDIR /app  
COPY . .

RUN CGO\_ENABLED=0 go build \-o app

FROM scratch

COPY \--from=builder /app/app /app

CMD \["/app"\]

Imagem final:

\~8MB

---

# **🧠 Insight de engenharia**

É por isso que muitas empresas:

* escrevem microservices em **Go**  
* usam **scratch**

