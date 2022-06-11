import os
import random
import string
import requests

differenciator = "1"
total_archivos = 100
archivos_por_usuario = 10
usuarios_total = total_archivos // archivos_por_usuario
all_users_gen = []
all_archivos_public_uploaded = []
all_archivos_grupo_uploaded = []
carpeta_archivos = "./archivos/1/"
cookies_login = None

def gen_invitaciones_registro(cookies):
  url = 'http://localhost:3000/api/invitaciones/registro/crear'
  data = {'fechaCaducidad': '2023-05-05', 'maximoUsos': str(usuarios_total)}
  response = requests.post(url, data=data, cookies=cookies)
  return response.json()


def gen_clave_encript(length):
  return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))


def upload_file_public(file_path, clave_encripcion, cookies):
  url = 'http://localhost:3000/api/archivos/publico/upload'
  files = {'file': open(file_path, 'rb')}
  data = {'clave': clave_encripcion}
  response = requests.post(url, files=files, data=data, cookies=cookies)
  return response.json()


def get_cookies_login_user(username, password):
  url = 'http://localhost:3000/api/auth/login'
  data = {'username': username, 'password': password}
  response = requests.post(url, data=data)
  return response.cookies


def generate_usuarios(invitacion):
  for i in range(usuarios_total):
    username = differenciator + 'usuario_' + str(i)
    password = gen_clave_encript(16)
    all_users_gen.append({
      "username": username,
      "password": password,
    })
    url = 'http://localhost:3000/api/usuarios'
    data = {'username': username, 'password': password, 'invitacion': invitacion}
    response = requests.post(url, data=data)
    print(response.json())


def upload_files_carpeta_and_join_group(carpeta, codigo_invitacion_grupo, uuid_grupo):
  cookie_usuario = None
  i = 0
  usuario_act = len(all_users_gen) - 1
  for file_name in os.listdir(carpeta):
    if i % archivos_por_usuario == 0:
      usuario = all_users_gen[usuario_act]
      cookie_usuario = get_cookies_login_user(usuario["username"], usuario["password"])
      unirse_grupo(cookie_usuario, codigo_invitacion_grupo)
      usuario_act -= 1
    if file_name.endswith('.png'):
      if random.randint(0, 1) == 0:
        clave = file_name.split('.')[0]
        response = upload_file_public(carpeta_archivos + file_name, clave, cookie_usuario)
        all_archivos_public_uploaded.append({
          "clave": clave,
          "uuid": response
        })
      else:
        response_grupo = upload_file_group(carpeta_archivos + file_name, cookie_usuario, uuid_grupo)
        all_archivos_grupo_uploaded.append({
          "uuid": response_grupo
        })
    i += 1



def upload_file_group(file_path, cookies, uuid_grupo):
  url = 'http://localhost:3000/api/archivos/grupo/upload'
  files = {'file': open(file_path, 'rb')}
  data = {'grupo': uuid_grupo}
  response = requests.post(url, files=files, data=data, cookies=cookies)
  return response.json()


def create_grupo(nombre, cookies):
  url = 'http://localhost:3000/api/grupos/'
  data = {'nombre': nombre}
  response = requests.post(url, data=data, cookies=cookies)
  return response.json()


def crear_invitacion_grupo(cookies, uuid_grupo):
  url = 'http://localhost:3000/api/invitaciones/grupo/crear'
  data = {"fechaCaducidad": "2023-05-05", "maximoUsos": str(usuarios_total), "grupoUuid": uuid_grupo}
  response = requests.post(url, data=data, cookies=cookies)
  data = response.json()
  print("Invitacion creada: " + str(data))
  return data


def unirse_grupo(cookies, invitacion):
  url = 'http://localhost:3000/api/grupos/join'
  data = {'invitacion': invitacion}
  response = requests.post(url, data=data, cookies=cookies)
  return response.json()


def guardar_info_subida(archivo, datos):
  with open('./datos/' + archivo, 'w') as f:
    f.write(str(datos))


def main():

  inicio = 1
  final = 20

  for i in range(inicio, final + 1):

    global differenciator
    differenciator = str(i)

    global carpeta_archivos
    carpeta_archivos = "./archivos/" + differenciator + "/"

    if i == inicio:
      cookies_login = get_cookies_login_user('admin', 'admin_admin')
    else:
      cookies_login = get_cookies_login_user(all_users_gen[i]["username"], all_users_gen[i]["password"])

    codigo_registro = gen_invitaciones_registro(cookies_login)["data"]["Codigo"]
    uuid_grupo = create_grupo(gen_clave_encript(5), cookies_login)["data"]["uuid"]
    generate_usuarios(codigo_registro)
    codigo_invitacion_grupo = crear_invitacion_grupo(cookies_login, uuid_grupo)["data"]["Codigo"]
    upload_files_carpeta_and_join_group(carpeta_archivos, codigo_invitacion_grupo, uuid_grupo)
  
  guardar_info_subida('all_users_gen.json', all_users_gen)

if __name__ == '__main__':
  main()