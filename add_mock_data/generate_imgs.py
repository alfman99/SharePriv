from concurrent.futures import ThreadPoolExecutor
from importlib.resources import path
import os
import random
import string
from randimage import get_random_image
import matplotlib
from faker import Faker

fakeInst = Faker('es_ES')

def gen_clave_encript(length):
  return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))

def proceso(carpeta):

  carpeta = str(carpeta)

  try:
    if not os.path.exists('./archivos/' + carpeta):
      os.makedirs('./archivos/' + carpeta)
    for _ in range(1000):
      clave_encripcion = gen_clave_encript(32)
      image_size = (random.randint(16, 32), random.randint(16, 32))
      image = get_random_image(image_size)
      matplotlib.image.imsave('./archivos/' + carpeta + '/' + clave_encripcion + '.' + fakeInst.file_extension(category='image'), image)
    return

  except Exception as e:
    print('EXCEPTION', e)
    

def main():
  num_carpetas = 10
  carpetas = list(range(num_carpetas))

  with ThreadPoolExecutor(num_carpetas) as executor:
    executor.map(proceso, carpetas)
    executor.shutdown(wait=True)     


if __name__ == '__main__':
  main()