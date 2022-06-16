import { Container } from "@mantine/core";
import GetArchivoPublico from "../GetArchivoPublico";
import UploadArchivoPublico from "../UploadArchivoPublico";

const Publicos = () => {

  return (
    <Container>
      <UploadArchivoPublico />
      <GetArchivoPublico />
    </Container>
  )

}
export default Publicos;