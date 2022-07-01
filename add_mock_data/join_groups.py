from concurrent.futures import ThreadPoolExecutor
import json
import random
import string
import time
import requests
from faker import Faker

num_threads = 150
users_gen = []

codigos_invitaciones_grupos = []

fakeInst = Faker('es_ES')

def gen_clave_encript(length):
  return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))

def str_time_prop(start, end, time_format, prop):
  stime = time.mktime(time.strptime(start, time_format))
  etime = time.mktime(time.strptime(end, time_format))
  ptime = stime + prop * (etime - stime)
  return time.strftime(time_format, time.localtime(ptime))

def random_date(start, end, prop):
  return str_time_prop(start, end, '%Y-%m-%d', prop)

def get_cookies_login_user(username, password):
  url = 'http://localhost:3000/api/auth/login'
  data = {'username': username, 'password': password}
  response = requests.post(url, data=data)
  return response.cookies

def process_user(profile):
  try:
    num_grups = random.randint(1, 5)
    cookies = get_cookies_login_user(profile['username'], profile['password'])

    for i in range(num_grups):
      print(i)
      invitacion_grupo = random.choice(codigos_invitaciones_grupos)
      requests.post('http://localhost:3000/api/grupos/join', cookies=cookies, data={'invitacion': invitacion_grupo})
    return

  except Exception as e:
    print('EXCEPTION', e)
    exit()

def start_things():
  users = {}
  # Load json with all users
  with open('./datos/all_users_gen.json') as f:
    users = json.load(f)
  
  # Load array from file with all invitations
  with open('./datos/all_group_invitations.json') as f:
    global codigos_invitaciones_grupos
    codigos_invitaciones_grupos = json.load(f)
  
  with ThreadPoolExecutor(num_threads) as executor:
    executor.map(process_user, users)
    executor.shutdown(wait=True)

def main():
  start_things()

if __name__ == '__main__':
  main()