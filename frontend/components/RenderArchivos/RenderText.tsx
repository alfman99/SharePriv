import { Textarea } from "@mantine/core";
import { useEffect, useState } from "react";


const RenderTxt = (props: {
  documento: Blob;
}) => {

  const [text, setText] = useState<string>("");

  useEffect(() => {
    props.documento.text()
      .then((text) => {
        setText(text);
      });
  }, [props.documento]);

  return (
    <div style={{
      paddingTop: '1em',
      paddingBottom: '1em',
    }}>
      <Textarea
        value={text}
        autosize
        readOnly
      />
    </div>
  )
}

export default RenderTxt;