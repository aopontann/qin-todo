import React, { useState } from 'react';
import { InputWrapper, Input, Button } from '@mantine/core';
import { useRouter } from 'next/router';

const Login = () => {
  const [input, setInput] = useState({
    email: '',
    password: '',
  });

  const router = useRouter();

  const handleLogin = () => {
    const body = {
      email: input.email,
      password: input.password,
    };

    const param: RequestInit = {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json; charset=utf-8',
        Accept: 'application/json',
      },
      body: JSON.stringify(body),
    };

    fetch('/auth/login', param)
      .then((res) => res.json())
      .then((res) => {
        if (!res.error) {
          router.push('/');
        }
      });
  };

  return (
    <div className='grid min-h-screen content-center gap-4 px-6'>
      <InputWrapper className='grid gap-2' id='input-demo' label='Qin Todo にログイン' size='md'>
        <Input
          id='input-demo'
          placeholder='Your email'
          value={input.email}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setInput((prev) => ({ ...prev, email: e.target.value }))
          }
        />
        <Input
          id='input-demo'
          placeholder='Your password'
          value={input.password}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setInput((prev) => ({ ...prev, password: e.target.value }))
          }
        />
      </InputWrapper>
      <Button onClick={handleLogin}>ログインする</Button>
      <div className='flex items-center'>
        <div className='h-[1px] flex-1 bg-gray'></div>
        <span className='mx-4 text-black'>または</span>
        <div className='h-[1px] flex-1 bg-gray'></div>
      </div>
      <Button variant='outline'>アカウント登録する</Button>
    </div>
  );
};

export default Login;
