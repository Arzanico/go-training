# Algoritmosn mas comunes 

## 1. Big-O

**Qué estudiar:** complejidad temporal y espacial.
**Fundamento lógico:** medir cómo escala una solución cuando crece el input.
**Caso de uso real:** comparar si conviene buscar usuarios en una lista lineal o indexarlos para responder más rápido en una API.
**Estructura de dato conveniente en Go:** depende del problema; normalmente `slice`, `map`, árbol o heap según el costo buscado.

## 2. Hash map / set

**Qué estudiar:** detección de duplicados, conteo, búsquedas rápidas.
**Fundamento lógico:** acceso directo por clave para evitar recorrer todo repetidamente.
**Caso de uso real:** detectar IDs repetidos en un lote de requests o contar frecuencia de errores por código.
**Estructura de dato conveniente en Go:** `map[K]V` y para set usualmente `map[K]struct{}`.

## 3. Arrays y strings

**Qué estudiar:** recorridos, índices, subarrays, substrings.
**Fundamento lógico:** secuencia ordenada por posición, acceso por índice y control de límites.
**Caso de uso real:** recorrer resultados de una query, paginar elementos o parsear un identificador recibido por HTTP.
**Estructura de dato conveniente en Go:** `[]T`, `string`, y según el caso `[]byte` o `[]rune`.

## 4. Binary search

**Qué estudiar:** búsqueda en arrays ordenados y problemas de condición monótona.
**Fundamento lógico:** dividir el espacio de búsqueda en mitades y descartar una mitad en cada paso.
**Caso de uso real:** encontrar rápidamente una versión mínima válida, un umbral de capacidad o un valor en una lista ordenada.
**Estructura de dato conveniente en Go:** `slice` ordenado, por ejemplo `[]int` o `[]struct` con criterio comparable.

## 5. Two pointers

**Qué estudiar:** pares, arrays ordenados, reversas, deduplicación.
**Fundamento lógico:** mover dos índices coordinadamente para evitar comparaciones innecesarias.
**Caso de uso real:** buscar dos importes que sumen un target o limpiar duplicados en una lista ordenada.
**Estructura de dato conveniente en Go:** `slice` o `string` indexable.

## 6. Sliding window

**Qué estudiar:** substring/subarray más largo o corto bajo una condición.
**Fundamento lógico:** mantener una ventana válida y ajustarla sin recomputar todo desde cero.
**Caso de uso real:** encontrar la secuencia más larga de eventos sin repetidos o el tramo mínimo que cumple un límite de suma.
**Estructura de dato conveniente en Go:** `slice`, `string` y apoyo con `map` para conteos o posiciones.

## 7. Sorting

**Qué estudiar:** cuándo ordenar simplifica un problema.
**Fundamento lógico:** imponer orden para hacer visibles patrones, overlaps, duplicados o búsquedas eficientes.
**Caso de uso real:** ordenar logs por timestamp antes de consolidarlos o intervalos de agenda antes de mergearlos.
**Estructura de dato conveniente en Go:** `slice` ordenable con `sort.Slice` o `slices.Sort`.

## 8. BFS

**Qué estudiar:** recorrido por niveles, shortest path en grafos no ponderados o grillas.
**Fundamento lógico:** exploración por capas.
**Caso de uso real:** hallar la menor cantidad de pasos en una grilla o distancia entre nodos de una red no ponderada.
**Estructura de dato conveniente en Go:** grafo como lista de adyacencia `map[K][]K` o grilla `[][]T`, más una queue implementada con `slice`.

## 9. DFS

**Qué estudiar:** exploración profunda, componentes conectadas, ciclos.
**Fundamento lógico:** seguir un camino hasta el fondo y luego retroceder.
**Caso de uso real:** descubrir todos los servicios conectados a uno dado o recorrer dependencias de forma exhaustiva.
**Estructura de dato conveniente en Go:** lista de adyacencia `map[K][]K`, árbol, o grilla; stack implícito por recursión o explícito con `slice`.

## 10. Trees / BST

**Qué estudiar:** recorridos, altura, búsqueda, propiedades del BST.
**Fundamento lógico:** estructura jerárquica donde cada nodo conecta decisiones o subdivisiones.
**Caso de uso real:** modelar menús jerárquicos, reglas de decisión o búsquedas ordenadas por rango.
**Estructura de dato conveniente en Go:** `struct` con punteros a hijos, por ejemplo `type Node struct { Left, Right *Node }`.

## 11. Heap / Priority Queue

**Qué estudiar:** top K, scheduling, prioridades.
**Fundamento lógico:** mantener acceso eficiente al mínimo o máximo sin ordenar todo.
**Caso de uso real:** procesar primero el job más prioritario o conservar los top K scores de millones de registros.
**Estructura de dato conveniente en Go:** heap sobre `slice` usando `container/heap`.

## 12. Intervals

**Qué estudiar:** merge intervals, overlaps, meeting rooms.
**Fundamento lógico:** razonar sobre rangos y solapamientos después de ordenar.
**Caso de uso real:** consolidar franjas horarias de reservas o detectar solapamientos de ventanas de mantenimiento.
**Estructura de dato conveniente en Go:** `[]Interval` donde `Interval` es un `struct{ Start, End int }` o similar.

## 13. Recursión

**Qué estudiar:** caso base, llamada recursiva, stack.
**Fundamento lógico:** resolver un problema en términos de una versión más pequeña del mismo.
**Caso de uso real:** recorrer directorios anidados, árboles JSON o estructuras jerárquicas.
**Estructura de dato conveniente en Go:** árboles, grafos, estructuras anidadas y funciones recursivas sobre `struct` o `map`.

## 14. Backtracking

**Qué estudiar:** combinaciones, permutaciones, búsqueda de caminos.
**Fundamento lógico:** explorar opciones, probar, deshacer y seguir.
**Caso de uso real:** generar combinaciones válidas de configuración o encontrar rutas posibles bajo restricciones.
**Estructura de dato conveniente en Go:** `slice` para estado actual, más `map` o matriz para marcar visitados.

## 15. Graphs

**Qué estudiar:** nodos, aristas, lista de adyacencia, ciclos, componentes.
**Fundamento lógico:** modelar relaciones y conexiones entre entidades.
**Caso de uso real:** dependencias entre servicios, rutas entre ciudades o relaciones entre usuarios y recursos.
**Estructura de dato conveniente en Go:** `map[K][]K`, `[]Edge`, o `[][]int` según densidad y tipo de problema.

## 16. Greedy

**Qué estudiar:** problemas donde una decisión local puede llevar a una buena solución global.
**Fundamento lógico:** tomar la mejor decisión inmediata cuando el problema lo permite.
**Caso de uso real:** elegir tareas que no se solapan maximizando cantidad o asignar recursos con criterio local óptimo.
**Estructura de dato conveniente en Go:** suele apoyarse en `slice` ordenado; a veces `heap`.

## 17. Concurrency en Go

**Qué estudiar:** worker pool, channels, mutex, context cancellation.
**Fundamento lógico:** coordinar trabajo simultáneo sin romper consistencia.
**Caso de uso real:** procesar múltiples requests externos en paralelo con límite de workers y cancelación por timeout.
**Estructura de dato conveniente en Go:** `chan T`, `sync.Mutex`, `sync.WaitGroup`, `context.Context`.

## 18. Strings en Go

**Qué estudiar:** `byte` vs `rune`, UTF-8, indexing.
**Fundamento lógico:** una representación en memoria no siempre coincide con “caracteres humanos”.
**Caso de uso real:** validar nombres con acentos, cortar texto sin romper caracteres multibyte o contar caracteres visibles.
**Estructura de dato conveniente en Go:** `string`, `[]byte`, `[]rune` según si priorizas bytes, mutabilidad o caracteres Unicode.

## 19. Complejidad de estructuras en Go

**Qué estudiar:** costo de `map`, `slice`, `append`, copias, memoria.
**Fundamento lógico:** la estructura elegida define el costo real de la solución.
**Caso de uso real:** decidir entre usar `map` o `slice` en una API de alto tráfico para balancear latencia y memoria.
**Estructura de dato conveniente en Go:** `slice`, `map`, `array`, `struct`, canales y heaps según patrón de acceso.

## Prioridad en los algorithmos a saber
* Big-O
* Hash map / set
* Arrays y strings
* Binary search
* Two pointers
* Sliding window
* BFS / DFS
* Heap
* Intervals
* Concurrency básica en Go

## Que debes saber responder antes de decidir que algorithmo usar

1. **Qué resuelve**
2. **Cuál es la lógica detrás**
3. **Cuándo usarlo**
4. **Qué complejidad tiene**
