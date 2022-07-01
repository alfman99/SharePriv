import Image from "next/image";
import { useState } from "react";


const RenderImage = (props: {
  documento: Blob;
}) => {
  return (
    <Image
      src={URL.createObjectURL(props.documento)}
      alt="archivo"
      layout={'fill'}
      objectFit={'contain'}
    />
  )

}

export default RenderImage;