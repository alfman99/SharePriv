import Image from "next/image";


const RenderImage = (props: {
  documento: Blob;
}) => {

  return (
    <div style={{
      paddingTop: '1em',
    }}>
      <Image
        src={URL.createObjectURL(props.documento)}
        alt="archivo"
        width={200}
        height={200}
      />
    </div>
  )

}

export default RenderImage;