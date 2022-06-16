import { Button, Group, Text, TextInput } from "@mantine/core";
import { Dropzone, DropzoneStatus, IMAGE_MIME_TYPE } from "@mantine/dropzone";
import { useContext, useState } from "react";
import { AuthContext } from "../contexts/AuthContext";
import { Upload, Photo, X, Icon as TablerIcon } from 'tabler-icons-react';
import ShowArchivo from "./ShowArchivo";
import { useForm } from "@mantine/form";
import { generateRandomString } from "../util/Util";


function ImageUploadIcon({
  status,
  ...props
}: React.ComponentProps<TablerIcon> & { status: DropzoneStatus }) {
  if (status.accepted) {
    return <Upload {...props} />;
  }

  if (status.rejected) {
    return <X {...props} />;
  }

  return <Photo {...props} />;
}

export const dropzoneChildren = (status: DropzoneStatus) => (
  <Group position="center" spacing="xl" style={{ minHeight: 220, pointerEvents: 'none' }}>
    <ImageUploadIcon status={status} style={{ color: '#C0FFEA' }} size={80} />

    <div>
      <Text size="xl" inline>
        Arrastra el archivo aqui o haga click para seleccionar el archivo
      </Text>
      <Text size="sm" color="dimmed" inline mt={7}>
        Adjunte un archivo, no puede exceder los 10 Mb
      </Text>
    </div>
  </Group>
);

const UploadArchivoPublico = () => {

  const { requestAuthenticated } = useContext(AuthContext)
  const [documento, setDocumento] = useState<File>()

  const handleOnDrop = (file: File[]) => {
    setDocumento(file[0])
  }

  const handleOnReject = (fileRej: any) => {
    alert(fileRej[0].errors[0].message)
  }

  const mantineForm = useForm({
    initialValues: {
      clave: ''
    },
    validate: {
      clave: (value: string) => value.length === 32 ? undefined : 'La clave debe ser de 32 caracteres'
    }
  })

  const handleGenerarClave = () => {
    const clave = generateRandomString(32)
    mantineForm.setValues({ clave })
  }

  const handleUploadArchivo = async (event: {
    clave: string;
  }) => {
    const formData = new FormData()
    formData.append('file', documento!)
    formData.append('clave', event.clave)

    const response = await requestAuthenticated('http://localho.st:3000/api/archivos/publico/upload', {
      method: 'POST',
      body: formData
    });

    const data = await response.json()

    if (response.status === 200) {
      alert(data.message)
    }
    else {
      alert(data.message)
    }

  }

  return (
    <>
      <h2>Subir archivo publico</h2>
      {
        documento ? (
          <>
            <form onSubmit={mantineForm.onSubmit((values) => handleUploadArchivo(values))}>
              <Group style={{ alignItems: 'flex-end' }}>
                <TextInput {...mantineForm.getInputProps('clave')}  label={'Clave de encriptacion (IMPORTANTE NO PERDERLA!!)'} placeholder={'Ej. WaZ8UYUfVly3RH6WfQdEMfde0kL4wCJ2'} />
                <Button onClick={() => handleGenerarClave()}>Generar aleatoriamente</Button>
              </Group>
              <Button style={{ marginTop: '0.5em', marginRight: '0.5em'}} color={'red'} onClick={() => setDocumento(undefined)}>
                Cancel
              </Button>
              <Button style={{ marginTop: '0.5em', marginRight: '0.5em'}} color={'green'} type={'submit'}>
                Subir
              </Button>
            </form>
            <ShowArchivo documento={documento} />
          </>
        ) : (
          <Dropzone
            onDrop={(file: File[]) => handleOnDrop(file)}
            onReject={(file: any) => handleOnReject(file)}
            maxSize={10000000}
            accept={IMAGE_MIME_TYPE}
            multiple={false}
            >
            {dropzoneChildren}
          </Dropzone>
        )
      }

    </>
  )
}

export default UploadArchivoPublico;