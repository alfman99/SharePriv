insert into invitacion_registro (codigo, fecha_caducidad, maximo_usos) values ('aVMkUtYQTHikmmcy', '2023-05-05', 1);
insert into usuario(username, password, invitacion_registro_codigo) values ('admin', 'admin_admin', 'aVMkUtYQTHikmmcy');
UPDATE "invitacion_registro" SET "usos"=1,"propietario"='admin' WHERE "codigo" = 'aVMkUtYQTHikmmcy';
