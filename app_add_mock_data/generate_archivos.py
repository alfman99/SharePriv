import random
import string
from randimage import get_random_image
import matplotlib

def gen_clave_encript(length):
  return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))

def main():

  inicio = 20
  final = 20

  for i in range(inicio, final + 1):
    for _ in range(100):
      clave_encripcion = gen_clave_encript(32)
      carpeta = './archivos/' + str(i) + '/'
      image_size = (128, 128)
      image = get_random_image(image_size)
      matplotlib.image.imsave(carpeta + clave_encripcion + '.png', image)

if __name__ == '__main__':
  main()