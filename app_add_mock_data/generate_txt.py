from concurrent.futures import ThreadPoolExecutor
import os
import random
import string
from faker import Faker

num_threads = 250
fakeInst = Faker('en_US')

def gen_clave_encript(length):
  return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))

def store_txt_in_file(txt):
  with open('./archivos/txt/' + gen_clave_encript(32) + '.txt', 'w') as f:
    f.write(txt + '\n')

def process_user(text):
  try:
    store_txt_in_file(text)
    return
  except Exception as e:
    print('EXCEPTION', e)
    exit()

def start_things(num_txt):
  txt_to_gen = [fakeInst.sentence(nb_words=random.randint(500, 1500)) for _ in range(num_txt)]
  if not os.path.exists('./archivos/txt/'):
    os.makedirs('./archivos/txt/')
  with ThreadPoolExecutor(num_threads) as executor:
    executor.map(process_user, txt_to_gen)
    executor.shutdown(wait=True)

def main():
  start_things(15000)

if __name__ == '__main__':
  main()