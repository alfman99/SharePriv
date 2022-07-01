from concurrent.futures import ThreadPoolExecutor
import random
import string
import time
import requests
from faker import Faker

num_threads = 500
users_gen = []

codigos_invitaciones = [
  'aVMkUtYQTHikmmcy', 
  'GsQcZokuxIkNDlNO', 
  'hqWeQijYUiueLIqY', 
  'OnzOUSVBUecXxVIM', 
  'eYwYyhkfPnQVRAHg', 
  'UbDehClAIZHHTPer', 
  'evtRQqrGKqSTjYLn', 
  'WNgWCHPtEKQJtpse', 
  'XNruNQEXwxPdlIQl', 
  'zwVWUjDXSxXUCYxw', 
  'EtovGBdoxNzwWmcQ', 
  'GrHBgijVrbbbCLqx', 
  'EmGXgYDQddQYsrvw', 
  'hIqvYPHUAJEfkRpg', 
  'UvsjqcGHeFrbDHMN', 
  'AbqEwyUfCraYkTPC', 
  'OHKWmnNnxrnBntPb', 
  'nupzUInJkFLpTlmo', 
  'TrzHUQkqPKHPLqlO', 
  'ixvZJFbWHawPAJri', 
  'lLZzRPrelQIKslXn', 
  'cmNtBMAfUaCirHAm', 
  'dJggqfezrNkExkId', 
  'PIONbgDxAfdCnmcS', 
  'sQfDQtrvetwUDQeW', 
  'TWeGxzZFKyTWXprM', 
  'uUfRCCUFTnfegegg', 
  'njRJWiLlObdknXxt', 
  'wjlnGwjrKfPUIiiE', 
  'eXakcHsDJlrWXdwq', 
  'ecsSBFQMgymmIQmW', 
  'qWXRlgfkCkuzrzSA', 
  'LnzPwJlKIURWrReL', 
  'SFYjadeEtPJyStQy', 
  'SjFHCrdMbSbmZZAj', 
  'AFnsKmAVhNpdSVKj', 
  'pJumsvQnNPSwEiTJ', 
  'iIMMlOYqiWfBxOUa', 
  'uNWWuwnFrwjxEsUa', 
  'tUTSPBbbrjeTVTNT', 
  'GNYMTLVkqtQbDcxG', 
  'ZnbddRwyElJRseDu', 
  'tegJAodOywwaipqB', 
  'hUBcWOOxKfZySoPJ', 
  'OahxqturexIPaPcD', 
  'VTgdYdOxdgtqLriC', 
  'ZLPLqOEfSGVZzfFy', 
  'RyqMqpwzWWjDrolq', 
  'WewHKDDLHasJmSxp', 
  'npligfDmnaaEUcwO', 
  'ZZqtcuPFTdHBUkqV', 
  'zNcPdOYspJzUYvvH', 
  'VhRLVxfZYFzZApzD', 
  'NIsijJQcENwvaEeJ', 
  'ymQiywSQHYqFPoBZ', 
  'qGPMLTlXBoLDpfht', 
  'UyroyqteMJyaBEoO', 
  'qdiqqMqtNvsTpNkh', 
  'sEHeHbOPOsxrJKLn', 
  'ccSpLkuaIPEboMTi',
  'YQqEWDJXPefQJBhK',
  'FKvqsHzuQdSmWoyc',
  'LHPexWcTvfmGyrNO',
  'HAbpuLfbgeQHUHsX',
  'RPOEIPaJHoXEOoCy',
  'qZByKqiNWKMQjCJa',
  'bOSxaDExsAlWaVJt',
  'aRvgREIhUbfydDtE',
  'lDzhgKzVabsqrHCA',
  'eDemGgkZrVWvyQyL',
  'ORtRCfEKZWtsIyYh',
  'YVPgLnPiucgQCzkf',
  'qXYemkTmmvTrVtOp',
  'leTfElrtohnSDBBx',
  'jRdmDAmwlupkJsKi',
  'HCsAAcQOLJlDMFyc',
  'UZFEoZVxJfakCyug',
  'RGeGvnwdcLajiauU',
  'tRxtRCBnwLdQxZBk',
  'iNRQAhwRDRIoYbjx',
  'vxzjEvmDirxlYDNR',
  'RaUYnPnjaAwubgpY',
  'WyWvBmNIKxQyxSbs',
  'HYdaMwHSXXAxfxCR',
  'kFzMrtpNgkZeZEaZ',
  'zvRYWmLlFmgMnwYZ'
  ]

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

  # time.sleep(random.randint(1, 30))

  # print (codigos_invitaciones)

  try:
    profile['Invitacion'] = random.choice(codigos_invitaciones)

    url = 'http://localhost:3000/api/usuarios'
    response = requests.post(url, data=profile)

    # print(response)

    if True:
      # print("GENERANDO NUEVO CODIGO")
      cookies = get_cookies_login_user(profile['username'], profile['password'])
      # print('COOKIES', cookies)
      max_usos = random.randint(3, 400)
      # print('MAX_USOS', max_usos)
      response_code = requests.post('http://localhost:3000/api/invitaciones/registro/crear', cookies=cookies, data={'fechaCaducidad': random_date("2022-06-10", "2050-1-1", random.random()), 'maximoUsos': max_usos})
      # print('RESPONSE?', response_code.json())
      codigos_invitaciones.append(response_code.json()['data']['Codigo'])

    return

  except KeyboardInterrupt:
    # print("Interrupted by user")
    exit()

def start_things(num_users):
  users_to_gen = [{"username": fakeInst.user_name(), "password": fakeInst.password()} for _ in range(num_users)]
  with ThreadPoolExecutor(num_threads) as executor:
    executor.map(process_user, users_to_gen)
    executor.shutdown(wait=True)

def main():
  start_things(5000)

if __name__ == '__main__':
  main()