import type { NextPage } from 'next'

import { Button, Space } from '@mantine/core'
import { AuthContext } from '../contexts/AuthContext'
import { useContext } from 'react'
import Selector, { SelectPageProps } from './SelectPage'
import Publicos from './LoggedInTabs/Publicos'
import { Lego, Network, PlaylistAdd } from 'tabler-icons-react'
import Grupos from './LoggedInTabs/Grupos'
import CrearInvitacionRegistro from './CrearInvitacionRegistro'

const LoggedIn: NextPage = () => {

  const { logout, user } = useContext(AuthContext)

  const pages: SelectPageProps[] = [
    {
      component: <Publicos />,
      icon: <Network size={20} />,
      label: 'Publicos'
    },
    {
      component: <Grupos />,
      icon: <Lego size={20} />,
      label: 'Grupos'
    },
    {
      component: <CrearInvitacionRegistro />,
      icon: <PlaylistAdd size={20} />,
      label: 'Invitaciones Registro'
    },
  ]

  return (
    <>
      <p>Welcome {user.user}</p>

      <Button color={'red'} onClick={() => logout()}>
        Logout
      </Button>

      <Space h={'xl'} />

      <Selector tabs={pages} />
    </>
  )
}

export default LoggedIn
