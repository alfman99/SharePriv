import RenderImage from "./RenderArchivos/RenderImage";
import RenderText from "./RenderArchivos/RenderText";

const ShowArchivo = (props: {
  documento: Blob | undefined;
  textRows: number;
}) => {
  const { documento, textRows } = props

  const selectRenderer = (documento: Blob) => {

    switch (documento.type) {
      case 'image/jpeg':
      case 'image/png':
      case 'image/gif':
      case 'image/bmp':
        return <RenderImage documento={documento} />
      case 'text/plain':
      case 'text/plain; charset=utf-8':
        return <RenderText rows={textRows} documento={documento} />
      default:
        return null
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