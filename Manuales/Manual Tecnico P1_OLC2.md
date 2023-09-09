Universidad de San Carlos de Guatemala 

Facultad de Ingeniería Escuela de Ingeniería en Ciencias y Sistemas 

Organización de Lenguajes y Compiladores 2 Segundo semestre 2023

Sebastian Alejandro de Leon Tenaz	
201906085

# Manual Técnico de T-Swift

## Introducción

T-Swift es un lenguaje de programación basado en Swift que ha ganado popularidad debido a su sintaxis moderna y características avanzadas. Este manual técnico proporcionará una guía detallada sobre el proyecto de T-Swift, que se desarrolló con el patrón intérprete. El proyecto involucra el uso de herramientas de Golang, ANTLR4 para el backend, React para el frontend y Graphviz DOT para la generación de informes.

## Objetivos

### Generales

-   Permitir a los estudiantes realizar el análisis de un compilador para un lenguaje de programación de alto nivel, centrándose en el procesamiento del lenguaje mediante herramientas para la generación de analizadores léxicos y sintácticos.

### Específicos

-   Ayudar a los estudiantes a utilizar herramientas léxicas y sintácticas para el reconocimiento y análisis del lenguaje.
-   Facilitar que los estudiantes implementen la traducción dirigida por la sintaxis utilizando reglas semánticas y atributos sintetizados/heredados.
-   Familiarizar a los estudiantes con las herramientas y estructuras de datos disponibles para la creación de un intérprete.

## Estructura del Proyecto

El proyecto de T-Swift se divide en las siguientes carpetas:

### 1. Environment

Esta carpeta contiene las definiciones de entorno y contexto utilizadas para el análisis semántico y la ejecución de código. Aquí se gestionan las variables, funciones y otros elementos del entorno de ejecución.

### 2. Expressions

En esta carpeta, se definen las expresiones del lenguaje T-Swift. Las expresiones son componentes fundamentales de cualquier lenguaje de programación y se utilizan para evaluar valores.

### 3. Instructions

Las instrucciones representan las acciones que puede realizar un programa en T-Swift. Estas instrucciones se ejecutan secuencialmente para llevar a cabo una tarea específica.

### 4. Interfaces

Las interfaces del proyecto proporcionan métodos y estructuras que se utilizan para la interacción entre componentes del intérprete. Esto incluye interfaces para el lexer, el parser y otros componentes clave.

### 5. Parser

En esta carpeta se encuentra el código relacionado con el análisis léxico y sintáctico de T-Swift. Aquí se definen las reglas de gramática y se utiliza ANTLR4 para generar el parser.

## Tecnologías Utilizadas

-   **Golang (Go):** Se utiliza como lenguaje principal para la implementación del intérprete de T-Swift, incluyendo el análisis semántico y la ejecución de código.
    
-   **ANTLR4:** Se utiliza para generar el lexer y parser a partir de la gramática definida para T-Swift.
    
-   **React:** Se emplea para desarrollar la interfaz de usuario (UI) del intérprete, permitiendo a los usuarios ingresar código T-Swift y visualizar los resultados.
    
-   **Graphviz DOT:** Se utiliza para generar informes visuales, como el árbol de sintaxis concreta (CST) y otros informes relacionados con la estructura del código fuente.
    

## Flujo de Trabajo

El flujo de trabajo típico del intérprete de T-Swift es el siguiente:

1.  El usuario ingresa el código T-Swift en la interfaz de usuario desarrollada en React.
    
2.  El código fuente se envía al backend implementado en Golang.
    
3.  El backend utiliza ANTLR4 para realizar el análisis léxico y sintáctico, generando un CST y llevando a cabo la traducción dirigida por la sintaxis.
    
4.  El análisis semántico verifica la corrección del código y evalúa las expresiones.
    
5.  Se genera el resultado de la ejecución del código T-Swift.
    
6.  Se utilizan Graphviz DOT para generar informes visuales, como el árbol CST, la tabla de símbolos y la tabla de errores.
    
7.  Los resultados se muestran al usuario a través de la interfaz de usuario.
    
## Código
### Ejemplos instructions:
#### Declaration.go 
La clase `Declaration` representa una declaración de variable en el lenguaje T-Swift. Esta declaración incluye información sobre la ubicación en el código fuente (línea y columna), el nombre de la variable, su tipo de datos, una expresión inicial (si la tiene) y una indicación de si es una variable global.

El método `Ejecutar` se encarga de ejecutar la declaración de la variable. Realiza validaciones de tipo y, si es una variable de tipo array, realiza una validación específica. Luego, almacena la variable en el entorno y agrega información sobre ella a la tabla de símbolos del árbol de sintaxis abstracta (AST).

El método `ArrayValidation` valida la estructura de un array, aunque en este ejemplo siempre devuelve `true`


![](https://github.com/sebbbasdl/OLC2_P1_201906085/blob/main/Manuales/Imagenes%20Tecnico/declaration.png)

#### Print.go
  
La clase `Print` representa una instrucción de impresión en el lenguaje T-Swift. Esta instrucción se utiliza para mostrar valores en la consola.

El método `Ejecutar` se encarga de ejecutar la instrucción de impresión. Dependiendo de si se proporciona un solo valor (`Value`) o dos valores (`Value` y `Value2`), obtiene el valor o valores correspondientes utilizando el método `Ejecutar` de la interfaz `Expression`. Luego, formatea los valores y los imprime en la consola a través del AST.

En resumen, la clase `Print` se utiliza para imprimir uno o dos valores en la consola del programa T-Swift.

![](https://github.com/sebbbasdl/OLC2_P1_201906085/blob/main/Manuales/Imagenes%20Tecnico/print.png)

### Ejemplos expressions:
#### Conversion.go
La clase `Conversion` representa una expresión de conversión de tipo en el lenguaje T-Swift. Esta expresión se utiliza para convertir un valor de un tipo de datos a otro tipo de datos. La conversión puede ser entre tipos numéricos (enteros y flotantes), entre numéricos y cadenas (strings) o entre booleanos y cadenas.

El método `Ejecutar` se encarga de realizar la conversión de tipo según las reglas del lenguaje. Depende del tipo de conversión (por ejemplo, de `INTEGER` a `STRING`), obtiene el valor original y realiza la conversión apropiada. En caso de que la conversión no sea posible, se maneja el error y se devuelve un valor nulo.

En resumen, la clase `Conversion` se utiliza para realizar conversiones de tipo en el lenguaje T-Swift, garantizando que los valores sean compatibles antes de la conversión. 


![](https://github.com/sebbbasdl/OLC2_P1_201906085/blob/main/Manuales/Imagenes%20Tecnico/conversion.png)

#### Count.go
La clase `Count` representa una expresión en el lenguaje T-Swift que se utiliza para contar el número de elementos en un array identificado por su ID.

El método `Ejecutar` se encarga de ejecutar esta expresión. Primero, intenta obtener la variable correspondiente al ID del array del entorno de ejecución. Si la variable no existe o no es un array, se emiten mensajes de error apropiados. Si es un array, se cuenta el número de elementos en el array y se devuelve un valor de tipo `INTEGER` que representa la cantidad de elementos en el array.

En resumen, la clase `Count` se utiliza para contar los elementos de un array en el lenguaje T-Swift y devuelve el resultado como un valor numérico.


![](https://github.com/sebbbasdl/OLC2_P1_201906085/blob/main/Manuales/Imagenes%20Tecnico/count.png)


## Conclusiones

El proyecto de T-Swift proporciona una plataforma educativa efectiva para que los estudiantes aprendan sobre la creación de un intérprete de lenguaje de programación. Combina varias tecnologías, incluyendo Golang, ANTLR4, React y Graphviz DOT, para ofrecer una experiencia completa de desarrollo de lenguajes de programación.
