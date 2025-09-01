# gocv

We use a configuration defined at `` as AppConfig
that contains:

```yaml
cvs:
    farmer:
        description: 'Currículum base: backend'
        last_compile: 
        long_description: |
            Currículum técnico con enfoque académico y profesional, orientado a roles backend e 
            infraestructura. Incluye experiencia enseñando algoritmos en UNAM, desarrollo de APIs REST
            con Go y Python, automatización con Terraform y CI/CD, y proyectos de software abiertos
            enfocados en microservicios, visualización geoespacial y votación electrónica. Formación
            académica en Matemáticas y Ciencias de la Computación (UNAM) con especialización en geometría
            computacional, programación concurrente y aprendizaje automático.
        template: cvBase.typ
    devops:
        description: 'Currículum base: backend'
        last_compile: 
        long_description: |
            Currículum técnico con enfoque académico y profesional, orientado a
            roles backend e infraestructura. Incluye experiencia enseñando algoritmos
            en UNAM, desarrollo de APIs REST con Go y Python, automatización con
            Terraform y CI/CD, y proyectos de software abiertos enfocados en
            microservicios, visualización geoespacial y votación electrónica.
            Formación académica en Matemáticas y Ciencias de la Computación (UNAM)
            con especialización en geometría computacional, programación concurrente
            y aprendizaje automático.
        template: cvFarmer.typ
default_template: cvBase.typ
output_dir: ~/cvs/
templates: ~/.config/gocv/templates
```

### Comandos

- init:
- compile:
- health:
- get:
- compile:
