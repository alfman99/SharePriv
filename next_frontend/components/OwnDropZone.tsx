import { Group, Text } from "@mantine/core";
import { Dropzone, DropzoneStatus, IMAGE_MIME_TYPE } from "@mantine/dropzone";
import { Upload, Photo, X, Icon as TablerIcon } from 'tabler-icons-react';

function ImageUploadIcon({
  status,
  ...props
}: React.ComponentProps<TablerIcon> & { status: DropzoneStatus }) {
  if (status.accepted) {
    return <Upload {...props} />;
  }

  if (status.rejected) {
    return <X {...props} />;
  }

  return <Photo {...props} />;
}

export const dropzoneChildren = (status: DropzoneStatus) => (
  <Group position="center" spacing="xl" style={{ minHeight: 220, pointerEvents: 'none' }}>
    <ImageUploadIcon status={status} style={{ color: '#C0FFEA' }} size={80} />
    <div>
      <Text size="xl" inline>
        Arrastra el archivo aqui o haga click para seleccionar el archivo
      </Text>
      <Text size="sm" color="dimmed" inline mt={7}>
        Adjunte un archivo, no puede exceder los 10 Mb
      </Text>
    </div>
  </Group>
);

const OwnDropZone = (props: { handleOnDrop: (file: any) => void; handleOnReject: (file: any) => void; }) => {

  const { handleOnDrop, handleOnReject } = props

  return (
    <Dropzone
      onDrop={(file: File[]) => handleOnDrop(file)}
      onReject={(file: any) => handleOnReject(file)}
      maxSize={10000000}
      accept={['image/*', 'text/*']}
      multiple={false}
      >
      {dropzoneChildren}
    </Dropzone>
  )


}

export default OwnDropZone;