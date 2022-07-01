import { Button, Container, Modal, Text } from "@mantine/core";
import { useContext, useState } from "react";
import { AuthContext, GroupData } from "../contexts/AuthContext";
import { formatDateBonitoDisplay } from "../util/Util";
import CrearInvitacionRegistro from "./CrearInvitacionRegistro";
import GenerarInvitacion from "./GenerarInvitacion";


const DisplayGroupInfo = (props: { group: GroupData; }) => {

  const { group } = props;

  return (
    <Container>
      <Text><strong>Nombre: </strong>{group.Nombre}</Text>
      <Text><strong>Propietario: </strong>{group.Propietario}</Text>
      <Text><strong>Fecha creación: </strong>{formatDateBonitoDisplay(group.FechaCreacion)}</Text>
    </Container>
  )
}

export default DisplayGroupInfo;