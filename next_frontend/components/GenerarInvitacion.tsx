import { Button, NumberInput, TextInput } from "@mantine/core";
import { DatePicker } from "@mantine/dates";
import { useForm } from "@mantine/form";
import { useContext } from "react";
import { AuthContext, GroupData } from "../contexts/AuthContext";
import { formatDateEnviarAPI } from "../util/Util";


const GenerarInvitacion = (props: { group?: GroupData; closeForm: () => void }) => {

  const { group, closeForm } = props;
  const { requestAuthenticated } = useContext(AuthContext);

  const form = useForm({
    initialValues: {
      fechaCaducidad: new Date(),
      maximoUsos: 0,
    },
    validate: {
      fechaCaducidad: (value) => new Date () >= new Date(value) ? "La fecha de caducidad no puede ser menor a la fecha actual" : undefined,
      maximoUsos: (value) => value <= 0 ? "El valor debe ser mayor a 0" : undefined,
    }
  });

  const handleSubmit = async (values: {
    fechaCaducidad: Date;
    maximoUsos: number;
  }) => {

    const multipartFormData = new FormData();
    multipartFormData.append("fechaCaducidad", formatDateEnviarAPI(values.fechaCaducidad));
    multipartFormData.append("maximoUsos", values.maximoUsos.toString());
    
    if (group) {
      
      multipartFormData.append("grupoId", group.Id.toString());
      
      const response = await requestAuthenticated(`http://localho.st:3000/api/invitaciones/grupo/crear`, {
        method: 'POST',
        body: multipartFormData,
      }) as Response

      if (response.status !== 200) {
        const dataResponse = await response.json();
        alert(dataResponse.message);
      }
      else {
        closeForm()
      }

    }
    else {
      const response = await requestAuthenticated(`http://localho.st:3000/api/invitaciones/registro/crear`, {
        method: 'POST',
        body: multipartFormData,
      }) as Response

      if (response.status !== 200) {
        console.log(response)
        const dataResponse = await response.json();
        alert(dataResponse.message);
      }
      else {
        closeForm()
      }
    }
    
  }

  return (
    <form onSubmit={form.onSubmit((values) => handleSubmit(values))}>
      <DatePicker placeholder={'Fecha de caducidad'} label={'Fecha de caducidad'} required {...form.getInputProps('fechaCaducidad')} />
      <NumberInput label={'Usos máximos'} {...form.getInputProps('maximoUsos')} />
      <Button style={{ marginTop: '1em' }} type="submit">Generar</Button>
    </form>
  )

}

export default GenerarInvitacion;