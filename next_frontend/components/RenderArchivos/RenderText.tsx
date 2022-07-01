import { Textarea } from "@mantine/core";
import { useEffect, useState } from "react";


const RenderTxt = (props: {
  documento: Blob;
  rows: number;
}) => {

  const { rows } = props;

  const [text, setText] = useState<string>("");

  useEffect(() => {
    props.documento.text()
      .then((text) => {
        setText(text);
      });
  }, [props.documento]);

  return (
    <Textarea
      value={text}
      maxRows={rows}
      autosize
      readOnly
    />
  )
}

export default RenderTxt;