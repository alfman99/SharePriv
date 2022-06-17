import { Button, Container, Tabs } from "@mantine/core";
import { GroupData } from "../contexts/AuthContext";
import { formatDateBonitoDisplay } from "../util/Util";


const DisplayGroupInfo = (props: { group: GroupData; }) => {

  const { group } = props;

  return (
    <Container>
      <h2>Nombre: {group.Nombre}</h2>
      <h2>Propietario: {group.Propietario}</h2>
      <h2>Fecha creación: {formatDateBonitoDisplay(group.FechaCreacion)}</h2>
    </Container>
  )
}

export default DisplayGroupInfo;