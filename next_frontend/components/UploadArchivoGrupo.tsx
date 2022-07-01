import { Button, Group, Text, TextInput } from "@mantine/core";
import { Dropzone, DropzoneStatus, IMAGE_MIME_TYPE } from "@mantine/dropzone";
import { useContext, useState } from "react";
import { AuthContext, GroupData } from "../contexts/AuthContext";
import ShowArchivo from "./ShowArchivo";
import { useForm } from "@mantine/form";
import { copiarElementoAlPortapapeles, generateRandomString } from "../util/Util";
import OwnDropZone from "./OwnDropZone";


const UploadArchivoGrupo = (props: { grupo: GroupData }) => {

  const { grupo } = props

  const { requestAuthenticated } = useContext(AuthContext)
  const [documento, setDocumento] = useState<File>()

  const [uploadResponse, setUploadResponse] = useState<any>()

  const handleOnDrop = (file: File[]) => {
    setDocumento(file[0])
  }

  const handleOnReject = (fileRej: any) => {
    alert(fileRej[0].errors[0].message)
  }

  const handleUploadArchivo = async () => {
    const formData = new FormData()
    formData.append('file', documento!)
    formData.append('grupo', grupo.Id)

    const response = await requestAuthenticated('http://localho.st:3000/api/archivos/grupo/upload', {
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
      <h2>Subir archivo grupo</h2>
      {
        documento ? (
          uploadResponse ? (
            <>
              <h3>Status: {uploadResponse.message}</h3>
            </>
          ) : (
            <>
                <Button style={{ marginTop: '0.5em', marginRight: '0.5em'}} color={'red'} onClick={() => setDocumento(undefined)}>
                  Cancel
                </Button>
                <Button style={{ marginTop: '0.5em', marginRight: '0.5em'}} color={'green'} onClick={() => handleUploadArchivo()}>
                  Subir
                </Button>
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

export default UploadArchivoGrupo;