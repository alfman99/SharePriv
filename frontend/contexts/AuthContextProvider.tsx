import React, { useEffect, useState } from 'react';
import { AuthContext, GroupData, GruposVacios, PerfilData, PerfilVacio } from './AuthContext';

export const AuthContextProvider = ({ children }: any) => {

  const [user, setUser] = useState<PerfilData>(PerfilVacio);
  const [groups, setGroups] = useState<GroupData[]>(GruposVacios);

  const fetchUserInfo = async () => {

    const token = localStorage.getItem('token')

    if (!token) {
      setUser(PerfilVacio)
      return;
    }

    const response = await fetch(`http://localho.st:3000/api/auth/validate`, {
      headers: {
        'Authorization': token
      }
    });
    if (response.status === 200) {
      const data = await response.json()
      setUser(data.datos)
    }
    else {
      setUser(PerfilVacio)
      return;
    }


    const response2 = await fetch(`http://localho.st:3000/api/auth/grupos`, {
      headers: {
        'Authorization': token
      }
    });

    if (response2.status === 200) {
      const data = await response2.json()
      setGroups(data.data)
    } else {
      setGroups(GruposVacios)
      return;
    }
  }

  useEffect(() => {
    fetchUserInfo()
  }, [])

  const login = async (username: string, password: string) => {
    const response = await fetch(`http://localho.st:3000/api/auth/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username,
        password
      })
    })
    if (response.status === 200) {
      // store token in local storage
      const data = await response.json()
      localStorage.setItem('token', data.token)
      await fetchUserInfo();
    }
    else {
      alert('Usuario o contraseña incorrectos')
    }
  }

  const logout = async () => {
    localStorage.removeItem('token')
    document.cookie = "";
    setUser(PerfilVacio)
  }

  const signup = async (username: string, password: string, invitacion: string) => {
    const response = await fetch(`http://localho.st:3000/api/usuarios`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username,
        password,
        invitacion
      })
    })

    const data = await response.json()

    if (response.status === 200) {
      login(username, password)
      return;
    }
    else {
      alert(data.message)
    }

  }

  const requestAuthenticated = async (url: string, options?: any) => {

    let token = localStorage.getItem('token')

    if (!token) {
      setUser(PerfilVacio)
      return;
    }

    const response = fetch(url, {
      headers: {
        'Authorization': token
      },
      ...options
    })

    return response;
  }

  return (
    <AuthContext.Provider value={{
      user,
      groups,
      login,
      logout,
      signup,
      requestAuthenticated
    }}>
      {children}
    </AuthContext.Provider>
  )
}