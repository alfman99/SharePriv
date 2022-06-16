import { Container, Tabs } from "@mantine/core";
import { Eye, FileUpload } from "tabler-icons-react";
import GetArchivoPublico from "../GetArchivoPublico";
import UploadArchivoPublico from "../UploadArchivoPublico";

const Publicos = () => {

  return (
    <Container>
      <Tabs>
        <Tabs.Tab label={'Subir archivo'} icon={<FileUpload size={20} />}>
          <UploadArchivoPublico />
        </Tabs.Tab>
        <Tabs.Tab label={'Ver archivo'} icon={<Eye size={20} />}>
          <GetArchivoPublico />
        </Tabs.Tab>
      </Tabs>
    </Container>
  )

}
export default Publicos;