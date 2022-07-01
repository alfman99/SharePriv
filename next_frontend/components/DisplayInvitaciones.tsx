import { Group, Table } from "@mantine/core";
import { formatDateBonitoDisplay } from "../util/Util";
import { Invitacion } from "./CrearInvitacionRegistro";

const DisplayInvitaciones = (props: { listaInvitaciones: Invitacion[] }) => {

  const { listaInvitaciones } = props;

  return (
    <>
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
    </>
  )

}

export default DisplayInvitaciones;