import { Button, Container } from "@mantine/core"
import { useContext } from "react"
import { AuthContext } from "../../contexts/AuthContext"
import ExploreGroups from "../ExploreGroups"

const Grupos = () => {

  const { groups } = useContext(AuthContext)

  return (
    <Container>
      <h2>Mis Grupos</h2>
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