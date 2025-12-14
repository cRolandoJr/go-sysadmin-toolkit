# Go Systems Engineering Toolkit

Este repositorio aloja una colección de herramientas de sistema (CLI) de alto rendimiento escritas en **Go**, diseñadas para operar en entornos Linux siguiendo la filosofía UNIX.

El objetivo principal de este proyecto es la **reimplementación desde cero** de utilidades core (`coreutils`) y herramientas de administración, enfocándose en la gestión eficiente de memoria, interacción directa con el Kernel (syscalls) y concurrencia, sin depender de librerías externas pesadas.

##  Herramientas Incluidas

### 1. GoMonitor (`cmd/gomonitor`)
Un monitor de recursos en tiempo real que interactúa directamente con el sistema de archivos virtual `/proc` de Linux.
- **Ingeniería:** Parsing manual de `/proc/stat` para calcular deltas de tiempo de CPU.
- **Rendimiento:** Zero-allocation parsing para minimizar la presión sobre el Garbage Collector.
- **Funcionalidad:** Cálculo preciso de uso de CPU (User + Nice + System vs Idle).

### 2. GoRotator (`cmd/gorotator`)
Herramienta de automatización para la rotación y compresión de logs.
- **Ingeniería:** Uso de `filepath.WalkDir` para recorrido eficiente del sistema de archivos.
- **Streams:** Implementación de pipelines de compresión (`io.Pipe`, `compress/gzip`) para procesar archivos mayores a la RAM disponible.
- **Seguridad:** Manejo atómico de archivos para evitar pérdida de datos durante la rotación.

### 3. GoLs (`cmd/gols`)
Listado de archivos con inspección de metadatos de sistema.
- **Syscalls:** Uso de aserciones de tipo `syscall.Stat_t` para acceder a inodos, UID y GID específicos de Linux.
- **User Space:** Mapeo de IDs numéricos a nombres de usuario/grupo mediante `os/user`.

### 4. GoCat (`cmd/gocat`)
Concatenación de flujos de datos.
- **Memoria O(1):** Uso de buffers fijos (32KB) para transferir Gigabytes de datos sin picos de consumo de RAM.
- **Interoperabilidad:** Manejo agnóstico de `Stdin` y archivos en disco mediante interfaces `io.Reader`.

##  Instalación y Uso

Requisitos: Go 1.25+ (Linux environment recommended)

```bash
# Clonar el repositorio
git clone [https://github.com/TU_USUARIO/go-sysadmin-toolkit](https://github.com/TU_USUARIO/go-sysadmin-toolkit)
cd go-sysadmin-toolkit

# Ejecutar una herramienta (Ej: Monitor)
go run ./cmd/gomonitor/main.go
