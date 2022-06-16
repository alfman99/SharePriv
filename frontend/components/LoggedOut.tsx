import type { NextPage } from 'next'

import { Button } from '@mantine/core'
import { AuthContext } from '../contexts/AuthContext'
import { useContext } from 'react'
import Link from 'next/link'

const LoggedOut: NextPage = () => {

  const { logout, user } = useContext(AuthContext)

  return (
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
  )
}

export default LoggedOut
