from concurrent.futures import ThreadPoolExecutor
import json
import random
import string
import time
import requests
from faker import Faker

num_threads = 250
fakeInst = Faker('es_ES')

def gen_clave_encript(length):
  return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))

def get_cookies_login_user(username, password):
  url = 'http://localhost:3000/api/auth/login'
  data = {'username': username, 'password': password}
  response = requests.post(url, data=data)
  return response.cookies

def str_time_prop(start, end, time_format, prop):
  stime = time.mktime(time.strptime(start, time_format))
  etime = time.mktime(time.strptime(end, time_format))
  ptime = stime + prop * (etime - stime)
  return time.strftime(time_format, time.localtime(ptime))

def random_date(start, end, prop):
  return str_time_prop(start, end, '%Y-%m-%d', prop)

def process_user(profile):
  try:
    cookies = get_cookies_login_user(profile['username'], profile['password'])

    rand_num = random.randint(1, 10)
    grupo_nombre = ""

    if rand_num == 1:
      grupo_nombre = "Grupo " + fakeInst.name()
    elif rand_num == 2:
      grupo_nombre = fakeInst.name() + " Grupo"
    elif rand_num == 3:
      grupo_nombre = fakeInst.name() + " " + fakeInst.name()
    elif rand_num == 4:
      grupo_nombre = fakeInst.name() + " " + fakeInst.name() + " Grupo"
    elif rand_num == 5:
      grupo_nombre = fakeInst.user_name() + " Group"
    elif rand_num == 6:
      grupo_nombre = fakeInst.user_name() + "_" + fakeInst.user_name()
    elif rand_num == 7:
      grupo_nombre = fakeInst.user_name() + "_" + fakeInst.user_name() + " Group"
    elif rand_num == 8:
      grupo_nombre = fakeInst.name() + "_" + fakeInst.user_name()
    elif rand_num == 9:
      grupo_nombre = fakeInst.name() + "_" + str(rand_num)
    else:
      grupo_nombre = fakeInst.name() + "_" + fakeInst.user_name() + "_" + fakeInst.user_name()


    url = 'http://localhost:3000/api/grupos/'
    response = requests.post(url, data={
      'nombre': grupo_nombre,
    }, cookies=cookies)

    json_datos = response.json()
    id_grupo = json_datos['data']['id']
    requests.post('http://localhost:3000/api/invitaciones/grupo/crear', cookies=cookies, data={
      'fechaCaducidad': random_date("2022-06-10", "2050-1-1", random.random()),
      'maximoUsos': random.randint(50, 400),
      'grupoId': id_grupo,
    })

    # print('RESPUESTA GRUPO', response)

    return

  except Exception as e:
    # print("Interrupted by user")
    print('EXCEPTION', e)
    exit()

def start_things():
  file_usuarios = open("X:/Carrera/DABD/SharePriv/app_add_mock_data/datos/all_users_gen.json", "r")
  data = json.load(file_usuarios)
  with ThreadPoolExecutor(num_threads) as executor:
    executor.map(process_user, data)
    executor.shutdown(wait=True)

def main():
  start_things()

if __name__ == '__main__':
  main()