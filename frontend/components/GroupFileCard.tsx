import { Card, Image, Text, Badge, Button, Group, useMantineTheme, Grid } from '@mantine/core';
import Link from 'next/link';
import { b64ToBlob, formatDateBonitoDisplay } from '../util/Util';
import { ArchivoGrupo } from './DisplayGroupArchivos';
import ShowArchivo from './ShowArchivo';

function GroupFileCard(props: { file: ArchivoGrupo; view: (url: string) => void; }) {

  const { file, view } = props;

  return (
    <div style={{ height: 'auto', backgroundColor: '#80ffdf9a', borderRadius: '0.3em' }}>
      <div style={{ position: 'relative', height: '10em', top: '1em'}}>
        <ShowArchivo textRows={6} documento={b64ToBlob(file.Data, file.Mime)} />
      </div>
      <div style={{ padding: '1em', wordBreak: 'break-all', paddingTop: '2em' }}>
        <Text><strong>Id: </strong>{file.Id}</Text>
        <Text><strong>Propietario: </strong>{file.Propietario}</Text>
        <Text><strong>Fecha: </strong>{formatDateBonitoDisplay(file.FechaPublicacion)}</Text>
        <Text><strong>Visualizaciones: </strong>{file.Visualizaciones}</Text>
        <Button onClick={() => view(`http://localho.st:3000/api/archivos/grupo/${file.Id}`)} style={{ marginRight: '0em', marginLeft: 'auto', display: 'block', marginTop: '1em' }} color={'dark'}>Abrir en grande</Button>
      </div>
    </div>
  );
}

export default GroupFileCard