import { Button, Container, Space, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import React, { useContext, useState } from "react";
import { AuthContext } from "../../contexts/AuthContext";
import ShowArchivo from "../ShowArchivo";

const Publicos = () => {

  const { requestAuthenticated } = useContext(AuthContext)

  const [documento, setDocumento] = useState<Blob>()

  const handleGetArchivo = async (event: {
    id_archivo: string;
    clave: string;
  }) => {
    const response = await requestAuthenticated(`http://localho.st:3000/api/archivos/publico/${event.id_archivo}/${event.clave}`) as Response;
    
    if (response.status === 200) {
      const blob = await response.blob()
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
      clave: ''
    },
    validate: {
      id_archivo: (value: string) => value.length === 32 ? undefined : 'El id debe tener 32 caracteres',
      clave: (value: string) => value.length === 32 ? undefined : 'La clave debe ser de 32 caracteres'
    }
  })

  return (
    <Container>
      <h2>Publicos</h2>
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
        <Button type={'submit'}>Get Archivo</Button>
      </form>
      <ShowArchivo documento={documento} />
    </Container>
  );
}
export default Publicos;