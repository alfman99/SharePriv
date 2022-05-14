insert into usuario(username, password, invitacion_registro_codigo) values ('hola_como', 'test242', 'aVMkUtYQTHikmmcy');
insert into invitacion_registro (codigo, fecha_caducidad, maximo_usos, propietario) values ('aVMkUtYQTHikmmcy', '2023-05-05', 50, 'hola_como');
insert into grupo(nombre, propietario_username) values ('grupaso', 'hola_como');
insert into usuarios_grupos values ('hola_como', '7ea15850-fd2b-4abc-88a3-5e2bfb869ce4');

insert into usuario(username, password, invitacion_registro_codigo) values ('hola_como_2', '346346', 'aVMkUtYQTHikmmcy');
insert into usuarios_grupos values ('hola_como_2', '0d0b27e1-7fc1-402d-ab16-e1a2329682d6');