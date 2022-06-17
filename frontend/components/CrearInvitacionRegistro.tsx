/* eslint-disable react-hooks/exhaustive-deps */
import { Button, Group, Modal, Table } from "@mantine/core";
import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../contexts/AuthContext";
import { formatDateBonitoDisplay } from "../util/Util";
import DisplayInvitaciones from "./DisplayInvitaciones";
import GenerarInvitacion from "./GenerarInvitacion";



export interface Invitacion {
  Codigo: string;  
	FechaCreacion: Date;
	FechaCaducidad: Date;
	MaximoUsos: number;
	Usos: number
  Propietario: string;
}

const CrearInvitacionRegistro = () => {

  const [listaInvitaciones, setListaInvitaciones] = useState<Invitacion[]>([]);

  const { requestAuthenticated } = useContext(AuthContext);

  const [openedModal, setOpenedModal] = useState<boolean>(false);

  const fetchInvitaciones = async () => {
    const response = await requestAuthenticated('http://localho.st:3000/api/invitaciones/registro/listar');
    const data = await response.json();
    setListaInvitaciones(data.data);
  }

  useEffect(() => {
    fetchInvitaciones();
  }, [openedModal])

  return (
    <div>
      <Modal
        opened={openedModal}
        onClose={() => setOpenedModal(false)}
        title="Crear Invitacion Registro"
        >
          <GenerarInvitacion closeForm={() => setOpenedModal(false)} />
      </Modal>
      <Group style={{ justifyContent: 'space-between' }}>
        <h1>Invitaciones Registro</h1>
        <Button onClick={() => setOpenedModal(true)}>Generar Invitación</Button>
      </Group>
      <DisplayInvitaciones listaInvitaciones={listaInvitaciones} />
    </div>
  );
}

export default CrearInvitacionRegistro