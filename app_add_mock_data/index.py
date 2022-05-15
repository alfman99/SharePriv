import requests
from thispersondoesnotexist import get_online_person

def upload_file_public(file_path, clave_encripcion):
  url = 'localhost:3000/api/archivos/publico/upload'
  files = {'file': open(file_path, 'rb')}
  data = {'clave': clave_encripcion}
  response = requests.post(url, files=files, data=data)
  return response.json()

async def descargar_archivos_mock(clave_encripcion):
  picture = await get_online_person()
  print(picture)


def main():
  file_path = './data/archivo.txt'
  clave_encripcion = 'clave'
  upload_file_public(file_path, clave_encripcion)

if __name__ == '__main__':
  main()