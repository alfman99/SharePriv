from concurrent.futures import ThreadPoolExecutor
import requests
import json

num_threads = 200
all_files_uploaded = []

imagenes = None

def guardar_info_subida(archivo, datos):
  with open(archivo, 'w') as f:
    f.write(str(datos))

def get_cookies_login_user(username, password):
  url = 'http://localhost:3000/api/auth/login'
  data = {'username': username, 'password': password}
  response = requests.post(url, data=data)
  return response.cookies

def add_file_to_group(cookies, group_id, img_id):
  url = 'http://localhost:3000/api/archivos/grupo/add'
  data = {
    'grupo': group_id,
    'id': img_id
    }
  try:
    response = requests.post(url, data=data, cookies=cookies)
    return response.json()
  except Exception as e:
    print("Exception!!", e)
    return None

def process_user(obj_user):
  try:
    cookies = get_cookies_login_user(obj_user["username"], obj_user["password"])
    for grupo in obj_user["grupos"]:
      for imagen in imagenes:
        response = add_file_to_group(cookies, grupo, imagen["archivo"])
        if response['message'] == 'El usuario no es el propietario del archivo':
          print(imagen)
          print(grupo)
          print(obj_user)
          print("El usuario no es el propietario del archivo")

    print("Added file to groups")
  except KeyboardInterrupt:
    print("Interrupted by user")
    exit()

def start_things(data):
  with ThreadPoolExecutor(num_threads) as executor:
    executor.map(process_user, data)
    executor.shutdown(wait=True)

def find_common_elements_json(json_usuarios, json_grupos):
  elements = []
  for usuario in json_usuarios:
    if "username" in usuario and "password" in usuario:
      element = {
        "username": usuario["username"],
        "password": usuario["password"],
        "grupos": []
      }
      for grupo in json_grupos:
        if usuario["username"] == grupo["username"]:
          element["grupos"].append(grupo["group_id"])
      if len(element['grupos']) >= 2:
        elements.append(element)
  return elements
       

# bruteforce xd
def main():
  file_usuarios = open("X:/Carrera/DABD/SharePriv/app_add_mock_data/datos/all_users_gen.json", "r")
  data_usuarios = json.load(file_usuarios)

  file_grupos = open("X:/Carrera/DABD/SharePriv/app_add_mock_data/datos/all_users_in_groups.json", "r")
  data_grupos = json.load(file_grupos)

  intersect = find_common_elements_json(data_usuarios, data_grupos)

  global imagenes
  file_imagenes = open("X:/Carrera/DABD/SharePriv/app_add_mock_data/datos/imagenes_pertenece_grupos.json", "r")
  imagenes = json.load(file_imagenes)
  # guardar_info_subida("X:/Carrera/DABD/SharePriv/app_add_mock_data/datos/test.json", json.dumps(intersect))

  start_things(intersect)
  # guardar_info_subida("X:/Carrera/DABD/SharePriv/app_add_mock_data/datos/all_files_uploaded_rapido_grupo.json", all_files_uploaded)

if __name__ == '__main__':
  main()