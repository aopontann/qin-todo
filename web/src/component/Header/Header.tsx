import React, { useState } from 'react';
import Image from 'next/image';
import { useSession, signIn, signOut } from 'next-auth/react';
import { Popover } from '@mantine/core';
import { useUserContext } from 'context/UserContext';

export const Header: React.VFC = () => {
  const { user } = useUserContext();
  const { data: session } = useSession();
  const [opened, setOpened] = useState(false);

  return (
    <header className='mx-break bg-filter sticky top-0 mb-5 grid items-center bg-white/60 px-6 py-2.5'>
      <div className='col-span-full row-span-full grid items-center justify-self-center'>
        <Image src='/logo.svg' alt='Qin Todo' width='113' height='24' layout='fixed' />
      </div>

      <Popover
        className='col-span-full row-span-full justify-self-end'
        opened={opened}
        onClose={() => setOpened(false)}
        target={
          <button className='block' onClick={() => setOpened((o) => !o)}>
            <div className='bg-gradient grid items-center overflow-hidden rounded-full p-[2px]'>
              <Image
                className='rounded-full'
                src={user.avatar_url !== '' ? user.avatar_url : 'https://placehold.jp/150x150.png'}
                alt='personLogo'
                width={32}
                height={32}
                layout='fixed'
              />
            </div>
          </button>
        }
        width={260}
        position='bottom'
      >
        {session ? (
          <button onClick={() => signOut()}>解除する</button>
        ) : (
          <button onClick={() => signIn()}>連携する</button>
        )}
      </Popover>
    </header>
  );
};
