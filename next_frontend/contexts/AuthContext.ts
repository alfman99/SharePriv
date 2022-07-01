import { createContext } from "react";

interface Claims {
  exp: number;
}

export interface PerfilData {
  StandardClaims: Claims,
  user: string;
}

export interface GroupData {
  Id: string;
  Nombre: string;
  FechaCreacion: Date;
  Propietario: string;
}

export interface IAuth {
  user: PerfilData;
  groups: GroupData[];
  fetchUserInfo: () => void,
  login: (username: string, password: string) => any,
  logout: () => void,
  signup: (username: string, password: string, invitacion: string) => any,
  requestAuthenticated: (url: string, options?: any) => any,
}

export const PerfilVacio: PerfilData = {
  StandardClaims: {
    exp: 0
  },
  user: ""
}

export const GruposVacios: GroupData[] = [];

export const AuthContext = createContext<IAuth>({
  user: PerfilVacio,
  groups: GruposVacios,
  fetchUserInfo: () => {},
  login: (username: string, password: string) => {},
  logout: () => {},
  signup: (username: string, password: string, invitacion: string) => {},
  requestAuthenticated: (url: string, options: any) => {}
});