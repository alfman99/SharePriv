import Image from "next/image";
import { useState } from "react";


const RenderImage = (props: {
  documento: Blob;
}) => {
  return (
    <div style={{
      marginTop: '1em',
      height: '15em',
      width: '100%',
      position: 'relative'
    }}>
      <Image
        src={URL.createObjectURL(props.documento)}
        alt="archivo"
        layout={'fill'}
        objectFit={'contain'}
      />
    </div>
  )

}

export default RenderImage;