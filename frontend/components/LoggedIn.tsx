import type { NextPage } from 'next'

import { Button, Space } from '@mantine/core'
import { AuthContext } from '../contexts/AuthContext'
import { useContext } from 'react'
import Selector, { SelectPageProps } from './SelectPage'
import Publicos from './LoggedInTabs/Publicos'
import { Photo } from 'tabler-icons-react'
import Grupos from './LoggedInTabs/Grupos'

const LoggedIn: NextPage = () => {

  const { logout, user } = useContext(AuthContext)

  const pages: SelectPageProps[] = [
    {
      component: <Publicos />,
      icon: <Photo size={14} />,
      label: 'Publicos'
    },
    {
      component: <Grupos />,
      icon: <Photo size={14} />,
      label: 'Grupos'
    },
  ]

  return (
    <>
      <p>Welcome {user.user}</p>

      <Button onClick={() => logout()}>
        Logout
      </Button>

      <Space h={'xl'} />

      <Selector tabs={pages} />
    </>
  )
}

export default LoggedIn
