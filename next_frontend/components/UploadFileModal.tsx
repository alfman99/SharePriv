import { Modal, Text } from "@mantine/core"
import { useState } from "react";
import { GroupData } from "../contexts/AuthContext";
import UploadArchivoGrupo from "./UploadArchivoGrupo";

const UploadFileToGrupoModal = (props: { opened: boolean; closeModal: () => void; grupo: GroupData }) => {

  const { opened, closeModal, grupo } = props;

  const [ uploadModalFile, setUploadModalFile ] = useState<File | null>(null);

  return (
    <Modal 
      opened={opened}
      transition="fade"
      transitionDuration={400}
      transitionTimingFunction="ease"
      onClose={() => closeModal()}
      style={{ backgroundColor: 'transparent' }}
      >
        <UploadArchivoGrupo grupo={grupo} />
    </Modal>
  )
}

export default UploadFileToGrupoModal;