import RenderImage from "./RenderArchivos/RenderImage";
import RenderText from "./RenderArchivos/RenderText";

const ShowArchivo = (props: {
  documento: Blob | undefined;
}) => {
  const { documento } = props

  const selectRenderer = (documento: Blob) => {

    switch (documento.type) {
      case 'image/jpeg':
      case 'image/png':
      case 'image/gif':
      case 'image/bmp':
        return <RenderImage documento={documento} />
      case 'text/plain':
        return <RenderText documento={documento} />
      default:
        return <h1>Error type not handled: {documento.type}</h1>
    }
}

  if (documento) {
    return selectRenderer(documento)
  }
  else {
    return null
  }

}

export default ShowArchivo