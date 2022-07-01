import type { NextPage } from 'next'

import { Button } from '@mantine/core'
import Link from 'next/link'

const LoggedOut: NextPage = () => {
  return (
    <>
      <Link href="/signup">
        <Button color={'green'}>
          Signup
        </Button>
      </Link>
      <Link href="/login">
        <Button style={{ marginLeft: '0.5em' }}>
          Login
        </Button>
      </Link>
    </>
  )
}

export default LoggedOut
