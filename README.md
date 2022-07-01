# Enunciado del problema
Página web en la que se pueden compartir y gestionar archivos, basado en usuarios y grupos. 

La única manera de registrarse es utilizando un código de invitación, el cual se genera por parte del usuario que lo invita, un usuario nuevo no pertenecerá a ningún grupo de usuarios nada más registrarse. Los usuarios pueden subir archivos publicos a través de la página web, estos archivos son encriptados con una clave simétrica (AES) a determinar en el momento de subir el archivo, se guarda en la base de datos el archivo encriptado y cualquier persona que tenga la clave de encriptación y el ID único del archivo podrá verlo. De los archivos publicos se almacena, el número de veces que se ha visto el archivo, la fecha de publicación y la privacidad de esta (pública o para un grupo de usuarios). Todos los usuarios que pertenecen a un grupo pueden listar y ver los archivos asociados a ese grupo, por lo que los archivos compartidos para el grupo no están encriptados, un archivo de grupo puede pertenecer a mas de un grupo diferente, el método de privacidad en este caso es pertenecer al grupo. Los usuarios pueden pertenecer a ningún grupo o a múltiples grupos a la vez. Los grupos de usuarios pueden ser creados únicamente por usuarios de la plataforma y las invitaciones a estos grupos de igual manera tienen fecha de caducidad y número de usos máximos, también queremos saber qué día se unió un usuario a un grupo. Queremos saber quien es el propietario del grupo para que en caso de que se tenga que eliminar el grupo, este sea el que lo borre, si se borra un grupo se borran también todos los archivos asociados al grupo. Los archivos pueden ser eliminados por el propietario de los archivos únicamente. Las invitaciones al grupo las puede hacer únicamente el propietario del grupo.

# Diagrama UML de la base de datos
![Diagrama UML](https://github.com/alfman99/SharePriv/blob/main/diagrama_UML.png)

# Datos introducidos a la base de datos para hacer la DEMO
- archivo_grupo: 25969
- archivo_publico: 23975
- archivos_de_grupos: 26979
- grupo: 7685
- invitacion_grupo: 7139
- invitacion_registro: 8484
- usuario: 13353
- usuarios_grupos: 9094

# Preview del frontend
![Preview Frontend NEXTjs](https://github.com/alfman99/SharePriv/blob/main/imagen_preview.png)

# Video DEMO
[![YouTube Video DEMO](https://img.youtube.com/vi/C-PFvMHBG5w/0.jpg)](https://www.youtube.com/watch?v=C-PFvMHBG5w)

# Nota y correcciones del profesor
- Nota final: 9.0
- Comentarios:
  - No se utilizan los miles de datos creados en la DEMO
  - Se podría haber añadido paginación a la lista de archivos del grupo y a la tabla de invitaciones, para que en caso de haber muchos datos no cargarlos todos de manera inecesaria
  - No se comenta si hay SQL injection en la parte del backend
  - El comentario del codigo podría haber sido mas extenso, para explicar como funciona la conexión a la base de datos con Go

# Comentarios personales
  - Mi experiencia utilizando Golang previa a este proyecto había sido para hacer aplicaciones de consola.
  - El video de la demo estaba limitado a 8 minutos para qeu no fuese excesivamente largo, pero podría haber estado durante horas hablando sobre el proyecto.
  - Me ha gustado mucho trabajar en Golang y creo que he aprendido mucho.
  - Estoy seguro de que será un lenguaje que utilice para hacer otros proyectos.