import { Button, Container, Group, TextInput } from "@mantine/core"
import { useContext, useState } from "react"
import { AuthContext } from "../../contexts/AuthContext"
import ExploreGroups from "../ExploreGroups"

const Grupos = () => {

  const { groups, requestAuthenticated, fetchUserInfo } = useContext(AuthContext)

  const [invitacionCodigo, setInvitacionCodigo] = useState<string>("")

  const [nombreNuevoGrupo, setNombreNuevoGrupo] = useState<string>("")

  const handleJoinGroup = async () => {
    if (invitacionCodigo.length != 16) {
      alert("El codigo de invitacion es invalido")
      return
    }

    const formData = new FormData()
    formData.append("invitacion", invitacionCodigo)

    const res = await requestAuthenticated(`http://localho.st:3000/api/grupos/join`, {
      method: "POST",
      body: formData,
    }) as Response

    if (res.status !== 200) {
      const dataResponse = await res.json()
      alert(dataResponse.message)
    }
    else {
      fetchUserInfo()
      setInvitacionCodigo('')
    }
  }

  const handleCrearGrupo = async () => {
    if (nombreNuevoGrupo.length == 0) {
      alert("El nombre del grupo no puede estar vacio")
      return
    }

    const formData = new FormData()
    formData.append("nombre", nombreNuevoGrupo)

    const res = await requestAuthenticated(`http://localho.st:3000/api/grupos`, {
      method: "POST",
      body: formData,
    }) as Response

    if (res.status !== 201) {
      const dataResponse = await res.json()
      alert(dataResponse.message)
    }
    else {
      fetchUserInfo()
      setNombreNuevoGrupo('')
    }
  }

  return (
    <Container>

      <Group style={{ justifyContent: 'space-between' }}>
        <h1>Mis Grupos</h1>
        <Group style={{ flexDirection: 'column', alignItems: 'flex-start' }}>
          <Group>
            <TextInput placeholder="Nombre del grupo" value={nombreNuevoGrupo} onChange={event => setNombreNuevoGrupo(event.target.value)} />
            <Button color={'green'} onClick={() => handleCrearGrupo()}>Crear grupo</Button>
          </Group>
          <Group>
            <TextInput placeholder="Invtacion a grupo" value={invitacionCodigo} onChange={event => setInvitacionCodigo(event.target.value)} />
            <Button onClick={() => handleJoinGroup()}>Unirse a grupo</Button>
          </Group>
        </Group>
      </Group>
      {
        groups.length > 0 ? (
          <ExploreGroups groups={groups} />
        ) : (
          <p>No tienes grupos</p>
        )
      }
    </Container>
  )
}

export default Grupos