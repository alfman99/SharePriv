/* eslint-disable react-hooks/exhaustive-deps */
import { Button, Group, Modal, Table } from "@mantine/core";
import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../contexts/AuthContext";
import { formatDateBonitoDisplay } from "../util/Util";
import GenerarInvitacion from "./GenerarInvitacion";



export interface Intivacion {
  Codigo: string;  
	FechaCreacion: Date;
	FechaCaducidad: Date;
	MaximoUsos: number;
	Usos: number
  Propietario: string;
}

const CrearInvitacionRegistro = () => {

  const [listaInvitaciones, setListaInvitaciones] = useState<Intivacion[]>([]);

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
      <Table>
        <thead>
          <tr>
            <th>Codigo</th>
            <th>Fecha Creacion</th>
            <th>Fecha Caducidad</th>
            <th>Maximo Usos</th>
            <th>Usos</th>
            <th>Propietario</th>
          </tr>
        </thead>
        {
          <tbody>
            {listaInvitaciones.map((invitacion) => (
              <tr key={invitacion.Codigo}>
                <td>{invitacion.Codigo}</td>
                <td>{formatDateBonitoDisplay(invitacion.FechaCreacion)}</td>
                <td>{formatDateBonitoDisplay(invitacion.FechaCaducidad)}</td>
                <td>{invitacion.MaximoUsos}</td>
                <td>{invitacion.Usos}</td>
                <td>{invitacion.Propietario}</td>
              </tr>
            ))}
          </tbody>
        }
      </Table>
      {
        listaInvitaciones.length === 0 ? (
          <Group style={{ justifyContent: 'center', width: '100%' }}>
          <p>No tienes ninguna invitación</p>
        </Group>
        ) : null
      }
    </div>
  );
}

export default CrearInvitacionRegistro