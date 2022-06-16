import { Button, Checkbox, Container, Space, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useContext, useState } from "react";
import { AuthContext } from "../contexts/AuthContext";
import ShowArchivo from "./ShowArchivo";

const GetArchivoPublico = () => {
  
  const { requestAuthenticated } = useContext(AuthContext)

  const [documento, setDocumento] = useState<Blob>()

  const handleGetArchivo = async (event: {
    id_archivo: string;
    clave: string;
    download: boolean;
  }) => {

    const response = await requestAuthenticated(`http://localho.st:3000/api/archivos/publico/${event.id_archivo}/${event.clave}`) as Response;
    
    if (response.status === 200) {
      const blob = await response.blob()

      if (event.download) {
        const link = document.createElement('a')
        link.href = window.URL.createObjectURL(blob)
        link.download = event.id_archivo
        link.click()
        link.remove();
      }

      setDocumento(blob)
    }
    else {
      const error = await response.json()
      alert(error.message)
    }
  }

  const form = useForm({
    initialValues: {
      id_archivo: '',
      clave: '',
      download: false
    },
    validate: {
      id_archivo: (value: string) => value.length === 32 ? undefined : 'El id debe tener 32 caracteres',
      clave: (value: string) => value.length === 32 ? undefined : 'La clave debe ser de 32 caracteres'
    }
  })

  return (
    <div style={{ marginBottom: '5em' }}>
      <h2>Ver archivo publico</h2>
      <form onSubmit={form.onSubmit((values) => handleGetArchivo(values))}>

        <TextInput 
          placeholder={'Ej. 79e3d6c206630d2293948131c78af9e6'} 
          {...form.getInputProps('id_archivo')} 
          label={'ID Archivo'} 
          required />

        <TextInput 
          placeholder={'Ej. VVLebfq7E8WmkRXuDimQ8MgWNrVIGiDH'} 
          {...form.getInputProps('clave')} 
          label={'Clave desencriptar'} 
          required />

        <Space h={'xs'} />
        <Checkbox style={{ marginBottom: '0.5em' }}
          {...form.getInputProps('download', { type: 'checkbox' })}
          label={'Descargar Archivo'} />
        <Button type={'submit'}>Ver archivo</Button>

      </form>
      <ShowArchivo documento={documento} />
    </div>
  );
}

export default GetArchivoPublico;