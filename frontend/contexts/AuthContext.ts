import { createContext } from "react";

interface Claims {
  exp: number;
}

export interface PerfilData {
  StandardClaims: Claims,
  user: string;
}

export interface IAuth {
  user: PerfilData;
  login: (username: string, password: string) => any,
  logout: () => void,
  signup: (username: string, password: string, invitacion: string) => any
}

export const PerfilVacio: PerfilData = {
  StandardClaims: {
    exp: 0
  },
  user: ""
}

export const AuthContext = createContext<IAuth>({
  user: PerfilVacio,
  login: (username: string, password: string) => {},
  logout: () => {},
  signup: (username: string, password: string, invitacion: string) => {}
});