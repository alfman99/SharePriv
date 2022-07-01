/* eslint-disable react-hooks/exhaustive-deps */
import { Button, Grid, Group, Modal, SimpleGrid } from "@mantine/core";
import { useContext, useEffect, useState } from "react";
import { AuthContext, GroupData } from "../contexts/AuthContext";
import { b64ToBlob } from "../util/Util";
import GroupFileCard from "./GroupFileCard";
import ShowArchivo from "./ShowArchivo";
import UploadFileToGrupoModal from "./UploadFileModal";

export interface ArchivoGrupo {
  Id: string;
  Data: string;
  Visualizaciones: number;
  FechaPublicacion: Date;
  Mime: string;
  Propietario: string;
}

const DisplayGroupArchivos = (props: { group: GroupData; }) => {

  const { group } = props;
  const { requestAuthenticated } = useContext(AuthContext);
  const [archivos, setArchivos] = useState<ArchivoGrupo[]>([]);

  const [openedModalArchivo, setOpenedModalArchivo] = useState<boolean>(false);
  const [ archivoModal, setArchivoModal ] = useState<Blob | null>(null);

  const [ openedUploadModal, setOpenedUploadModal ] = useState<boolean>(false);

  useEffect(() => {
    const fetchArchivosGrupo = async () => {
      const response = await requestAuthenticated(`http://localho.st:3000/api/grupos/${group.Id}/archivos`)
      const data = await response.json()
      setArchivos(data.data.archivos)
    }

    if (group.Id) {
      fetchArchivosGrupo();
    }
  }, [group])

  const viewOnModal = async (url: string) => {
    const response = await requestAuthenticated(url)
    const data = await response.blob()
    setOpenedModalArchivo(true)
    setArchivoModal(data)
  }

  const closeModalUploadFileToGrupo = () => {
    setOpenedUploadModal(false)
  }

  return (
    <>
      <Group style={{ justifyContent: 'space-between' }}>
        <h2>Archivos</h2>
        <Button onClick={() => setOpenedUploadModal(true)}>Subir archivo a éste grupo</Button>
        <UploadFileToGrupoModal opened={openedUploadModal} closeModal={closeModalUploadFileToGrupo} grupo={group}  />
      </Group>
      { 
        archivos.length > 0 ? 
        (
          <>
            <Modal withCloseButton={false}
              opened={openedModalArchivo}
              transition="fade"
              transitionDuration={400}
              transitionTimingFunction="ease"
              onClose={() => setOpenedModalArchivo(false)}
              style={{ backgroundColor: 'transparent' }}
            >
              <div style={{ height: '45em', backgroundColor: 'transparent' }}>
                {archivoModal ? <ShowArchivo textRows={31} documento={archivoModal} /> : <h1>Error</h1>}
              </div>
            </Modal>
            <SimpleGrid cols={3} style={{paddingBottom: '4em'}}>
              {
                archivos.map((ar: ArchivoGrupo, index: number) => {
                  return (
                    <div key={index}>
                      <GroupFileCard view={viewOnModal} file={ar} />
                    </div>
                  )
                })
              }
            </SimpleGrid>
          </>
        ) : (
          <p>No tienes archivos</p>
        )}
    </>
  );
}

export default DisplayGroupArchivos;