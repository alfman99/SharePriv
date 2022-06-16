import { Button, Container, Tabs } from "@mantine/core";
import { InfoCircle } from "tabler-icons-react";
import { GroupData } from "../contexts/AuthContext";
import { formatDate } from "../util/Util";


const DisplayGroupInfo = (props: { group: GroupData; }) => {

  const { group } = props;

  return (
    <Container>

      <h2>Nombre: {group.Nombre}</h2>
      <h2>Propietario: {group.Propietario}</h2>
      <h2>Fecha creación: {formatDate(group.FechaCreacion)}</h2>

    </Container>
  )
}

export default DisplayGroupInfo;