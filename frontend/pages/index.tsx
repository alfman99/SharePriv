import type { NextPage } from 'next'
import Link from 'next/link'

import { Button, Container } from '@mantine/core'
import { AuthContext } from '../contexts/AuthContext'
import { useContext } from 'react'

const Home: NextPage = () => {

  const { logout, user } = useContext(AuthContext)

  return (
    <Container>
      <h1>SharePriv</h1>
      {
        user.user == "" ? (
          <>
            <Link href="/signup">
              <Button>
                Signup
              </Button>
            </Link>
            <Link href="/login">
              <Button>
                Login
              </Button>
            </Link>
          </>
        ) : (
          <>
            <p>Welcome {user.user}</p>
            <Button onClick={() => logout()}>
              Logout
            </Button>
          </>
        )
      }
    </Container>
  )
}

export default Home
