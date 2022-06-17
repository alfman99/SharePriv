export const generateRandomString = (length: number) => {
  let text = '';
  const possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

  for (let i = 0; i < length; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }

  return text;
}

export const copiarElementoAlPortapapeles = (valor: string) => {
  navigator.clipboard.writeText(valor)
    .then(() => {
      alert('Texto copiado al portapapeles');
    }
    ).catch(err => {
      alert(err);
    });
}

export const formatDateBonitoDisplay = (date: Date) => {
  const d = new Date(date);
  return d.toLocaleDateString(undefined, { year: "numeric", month: "long", day: "numeric", hour: "numeric", minute: "numeric" })
}

const padTo2Digits = (num: number) => {
  return num.toString().padStart(2, '0');
}

export const formatDateEnviarAPI = (date: Date) => {
  const d = new Date(date);
  return [
    d.getFullYear(),
    padTo2Digits(d.getMonth() + 1),
    padTo2Digits(d.getDate()),
  ].join('-');
}

export const b64ToBlob = (b64Data: string, contentType: string) => {

  const byteCharacters = atob(b64Data);
  const byteNumbers = new Array(byteCharacters.length);
  for (let i = 0; i < byteCharacters.length; i++) {
    byteNumbers[i] = byteCharacters.charCodeAt(i);
  }
  const byteArray = new Uint8Array(byteNumbers);
  const blob = new Blob([byteArray], { type: contentType });
  return blob;

}