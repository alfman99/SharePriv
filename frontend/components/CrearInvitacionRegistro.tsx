/* eslint-disable react-hooks/exhaustive-deps */
import { Table } from "@mantine/core";
import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../contexts/AuthContext";
import { formatDate } from "../util/Util";



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

  useEffect(() => {

    const fetchInvitaciones = async () => {
      const response = await requestAuthenticated('http://localho.st:3000/api/invitaciones/registro/listar');
      const data = await response.json();
      setListaInvitaciones(data.data);
    }

    fetchInvitaciones();

  }, [])

  return (
    <div>
      <h1>Invitaciones Registro</h1>
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
        <tbody>
          {listaInvitaciones.map((invitacion) => (
            <tr key={invitacion.Codigo}>
              <td>{invitacion.Codigo}</td>
              <td>{formatDate(invitacion.FechaCreacion)}</td>
              <td>{formatDate(invitacion.FechaCaducidad)}</td>
              <td>{invitacion.MaximoUsos}</td>
              <td>{invitacion.Usos}</td>
              <td>{invitacion.Propietario}</td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
}

export default CrearInvitacionRegistro