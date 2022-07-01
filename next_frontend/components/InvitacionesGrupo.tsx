/* eslint-disable react-hooks/exhaustive-deps */
import { Button, Group, Modal } from "@mantine/core";
import { useContext, useEffect, useState } from "react";
import { AuthContext, GroupData } from "../contexts/AuthContext"
import { Invitacion } from "./CrearInvitacionRegistro";
import DisplayInvitaciones from "./DisplayInvitaciones";
import GenerarInvitacion from "./GenerarInvitacion";

const InvitacionesGrupo = (props: { group: GroupData }) => {

  const { group } = props;
  const { user, requestAuthenticated } = useContext(AuthContext);

  const [ listaInvitaciones, setListaInvitaciones ] = useState<Invitacion[]>([]);

  const [openedModal, setOpenedModal] = useState<boolean>(false);

  const fetchInvitacionesGrupo = async () => {
    const response = await requestAuthenticated(`http://localho.st:3000/api/grupos/${group.Id}/invitaciones`);
    const data = await response.json();
    
    if (response.status === 200) {
      setListaInvitaciones(data.data.invitaciones);
    }
    else {
      alert(data.message);
    }
  }

  useEffect(() => {
    if (group.Propietario == user.user) {
      fetchInvitacionesGrupo();
    }
  }, [openedModal])

  return (
    <>
      <Modal
        opened={openedModal}
        onClose={() => setOpenedModal(false)}
        title={`Crear Invitacion Grupo: ${group.Nombre}`}
        >
          <GenerarInvitacion group={group} closeForm={() => setOpenedModal(false)} />
      </Modal>

      <Group style={{ justifyContent: 'space-between' }}>
        <h2>Invitaciones del grupo</h2>
        <Button style={{ marginTop: '1em' }} onClick={() => setOpenedModal(true)}>Crear invitacion</Button>
      </Group>
      <DisplayInvitaciones listaInvitaciones={listaInvitaciones} />
    </>
  )

}

export default InvitacionesGrupo
