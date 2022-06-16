import type { NextPage } from 'next'
import Link from 'next/link'

import { Button, Container } from '@mantine/core'
import { AuthContext } from '../contexts/AuthContext'
import { useContext } from 'react'
import LoggedIn from '../components/LoggedIn'
import LoggedOut from '../components/LoggedOut'

const Home: NextPage = () => {

  const { logout, user } = useContext(AuthContext)

  return (
    <Container>
      <h1>SharePriv</h1>
      { user.user == "" ? <LoggedOut /> : <LoggedIn />}
    </Container>
  )
}

export default Home
