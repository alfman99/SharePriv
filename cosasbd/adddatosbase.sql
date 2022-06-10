insert into invitacion_registro (codigo, fecha_caducidad, maximo_usos) values ('aVMkUtYQTHikmmcy', '2023-05-05', 50);
insert into usuario(username, password, invitacion_registro_codigo) values ('admin', 'admin_admin', 'aVMkUtYQTHikmmcy');
insert into grupo(nombre, propietario_username) values ('grupaso', 'hola_como');
insert into usuarios_grupos values ('hola_como', '7ea15850-fd2b-4abc-88a3-5e2bfb869ce4');