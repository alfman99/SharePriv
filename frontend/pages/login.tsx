/* eslint-disable react-hooks/exhaustive-deps */
import { TextInput, Checkbox, Button, Group, Box, Container, PasswordInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { NextPage } from 'next';
import { useRouter } from 'next/router';
import { useContext, useEffect } from 'react';
import { AuthContext } from '../contexts/AuthContext';

const Login: NextPage = () => {

  const { login, user } = useContext(AuthContext)
  const router = useRouter()

  const form = useForm({
    initialValues: {
      username: '',
      password: ''
    },

    validate: {
      password: (value) => (value.length < 8 ? 'Password must be at least 8 characters' : null)
    },
  });

  useEffect(() => {
    if (user.user != '') {
      router.push('/')
    }
  }, [user])

  return (
    <Container>
      <form onSubmit={form.onSubmit((e) => login(e.username, e.password))}>
        <TextInput
          required
          label="Username"
          placeholder="your@username.com"
          {...form.getInputProps('username')}
        />
        <PasswordInput
          required
          label="Password"
          placeholder="********"
          {...form.getInputProps('password')}
        />

        <Group position="right" mt="md">
          <Button type="submit">Submit</Button>
        </Group>
      </form>
    </Container>
  );
}

export default Login