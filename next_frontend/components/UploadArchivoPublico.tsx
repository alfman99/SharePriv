import { Button, Group, Text, TextInput } from "@mantine/core";
import { useContext, useState } from "react";
import { AuthContext } from "../contexts/AuthContext";
import ShowArchivo from "./ShowArchivo";
import { useForm } from "@mantine/form";
import { copiarElementoAlPortapapeles, generateRandomString } from "../util/Util";
import OwnDropZone from "./OwnDropZone";

const UploadArchivoPublico = () => {

  const { requestAuthenticated } = useContext(AuthContext)
  const [documento, setDocumento] = useState<File>()

  const [uploadResponse, setUploadResponse] = useState<any>()

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
      setUploadResponse(data)
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
          uploadResponse ? (
            <>
              <h3>Status: {uploadResponse.message}</h3>
              <Group style={{ alignItems: 'flex-end' }}>
                <TextInput style={{ flexGrow: 1 }} readOnly={true} label="Id archivo" name="id_archivo" value={uploadResponse.data.Id} />
                <Button onClick={() => handleGenerarClave()}>Copiar ID</Button>
              </Group>
              <Group style={{ alignItems: 'flex-end' }}>
                <TextInput style={{ flexGrow: 1 }} readOnly={true} label="Clave" name="clave" {...mantineForm.getInputProps('clave')} />
                <Button onClick={() => handleGenerarClave()}>Copiar clave</Button>
              </Group>
              <Group style={{ alignItems: 'flex-end' }}>
                <TextInput style={{ flexGrow: 1 }} label="Compartir URL" readOnly={true} value={`http://localho.st:3000/api/archivos/publico/${uploadResponse.data.Id}/${mantineForm.values.clave}`} />
                <Button onClick={() => copiarElementoAlPortapapeles(`http://localho.st:3000/api/archivos/publico/${uploadResponse.data.Id}/${mantineForm.values.clave}`)}>Copiar URL</Button>
              </Group>
            </>
          ) : (
            <>
              <form onSubmit={mantineForm.onSubmit((values) => handleUploadArchivo(values))}>
                <Group style={{ alignItems: 'flex-end' }}>
                  <TextInput style={{ flexGrow: 1 }} {...mantineForm.getInputProps('clave')}  label={'Clave de encriptacion (IMPORTANTE NO PERDERLA!!)'} placeholder={'Ej. WaZ8UYUfVly3RH6WfQdEMfde0kL4wCJ2'} />
                  <Button onClick={() => handleGenerarClave()}>Generar aleatoriamente</Button>
                </Group>
                <Button style={{ marginTop: '0.5em', marginRight: '0.5em'}} color={'red'} onClick={() => setDocumento(undefined)}>
                  Cancel
                </Button>
                <Button style={{ marginTop: '0.5em', marginRight: '0.5em'}} color={'green'} type={'submit'}>
                  Subir
                </Button>
              </form>
              <div style={{ marginTop: '1em', height: '15em', width: '100%', position: 'relative' }}>
                <ShowArchivo textRows={10} documento={documento} />
              </div>
            </>
          )
        ) : (
          <OwnDropZone handleOnDrop={handleOnDrop} handleOnReject={handleOnReject} />
        )
      }
    </>
  )
}

export default UploadArchivoPublico;