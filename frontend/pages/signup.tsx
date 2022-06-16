/* eslint-disable react-hooks/exhaustive-deps */
import { TextInput, Checkbox, Button, Group, Box, Container, PasswordInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { NextPage } from 'next';
import { useRouter } from 'next/router';
import { useContext, useEffect } from 'react';
import { AuthContext } from '../contexts/AuthContext';

const Signup: NextPage = () => {

  const { signup, user } = useContext(AuthContext)
  const router = useRouter()

  const form = useForm({
    initialValues: {
      username: '',
      password: '',
      invitacion: ''
    },

    validate: {
      password: (value) => (value.length < 8 ? 'Password must be at least 8 characters' : null),
      invitacion: (value) => (value.length < 16 ? 'Invitation code must be 16 characters' : null)
    },
  });

  useEffect(() => {
    if (user.user != '') {
      router.push('/')
    }
  }, [user])

  return (
    <Container>
      <h1>SharePriv</h1>
      <form onSubmit={form.onSubmit((e) => signup(e.username, e.password, e.invitacion))}>
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

        <TextInput
          required
          label="Invitacion"
          placeholder="Ej. njRJWiLlObdknXxt"
          {...form.getInputProps('invitacion')}
        />

        <Group position="right" mt="md">
          <Button type="submit">Submit</Button>
        </Group>
      </form>
    </Container>
  );
}

export default Signup