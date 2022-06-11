from concurrent.futures import ThreadPoolExecutor
import os
import random
import string
import matplotlib
from randimage import get_random_image
import requests
import json

num_threads = 10
all_files_uploaded = []

def guardar_info_subida(archivo, datos):
  with open(archivo, 'w') as f:
    f.write(str(datos))

def gen_clave_encript(length):
  return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))

def gen_file():
  clave_encripcion = gen_clave_encript(32)
  carpeta = 'X:/Carrera/DABD/SharePriv/app_add_mock_data/tmp/'

  random_num = random.randint(1, 5)

  if random_num == 1:
    image_size = (128, 128)
    image = get_random_image(image_size)
    matplotlib.image.imsave(carpeta + clave_encripcion + '.png', image)
    return open(carpeta + clave_encripcion + '.png', 'rb')
  elif random_num == 2:
    txt_content = gen_clave_encript(random.randint(1, 2048))
    with open(carpeta + clave_encripcion + '.txt', 'w') as f:
      f.write(txt_content)
    return open(carpeta + clave_encripcion + '.txt', 'rb')
  else:
    image_size = (random.randint(100, 500), random.randint(100, 500))
    image = get_random_image(image_size)
    matplotlib.image.imsave(carpeta + clave_encripcion + '.jpeg', image)
    return open(carpeta + clave_encripcion + '.jpeg', 'rb')

def get_cookies_login_user(username, password):
  url = 'http://localhost:3000/api/auth/login'
  data = {'username': username, 'password': password}
  response = requests.post(url, data=data)
  return response.cookies

def upload_public_file(cookies):
  archivo = gen_file()
  url = 'http://localhost:3000/api/archivos/publico/upload'
  clave_encriptacion = gen_clave_encript(32)
  files = {'file': archivo}
  data = {'clave': clave_encriptacion}
  try:
    response = requests.post(url, files=files, data=data, cookies=cookies)
    archivo.close()
    os.remove(archivo.name)
    return {
      "id": response.text,
      "clave": clave_encriptacion
    }
  except Exception as e:
    print("Exception!!", e)
    return None

def process_user(obj_user):
  try:
    cookies = get_cookies_login_user(obj_user["username"], obj_user["password"])
    uploaded_file = upload_public_file(cookies)
    all_files_uploaded.append(uploaded_file)
    print("Se han subido unas cuantas imagenes")
  except KeyboardInterrupt:
    print("Interrupted by user")
    exit()

def start_things(data):
  with ThreadPoolExecutor(num_threads) as executor:
    executor.map(process_user, data)
    executor.shutdown(wait=True)      

def main():
  file_usuarios = open("X:/Carrera/DABD/SharePriv/app_add_mock_data/datos/all_users_gen.json", "r")
  data = json.load(file_usuarios)
  start_things(data)
  guardar_info_subida("X:/Carrera/DABD/SharePriv/app_add_mock_data/datos/all_files_uploaded_rapido.json", all_files_uploaded)

if __name__ == '__main__':
  main()