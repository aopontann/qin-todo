import React from 'react';
import Image from 'next/image';

export const Header = () => {
  return (
    <header className='max-w-5xl mx-auto flex justify-between items-center h-22'>
      <div>
        {/* <img src='/logo.svg' alt='Qin Todo' /> */}
        <Image src='/logo.svg' alt='Qin Todo' width={100} height={100} />
        {/* https://nextjs.org/docs/api-reference/next/image */}
      </div>
      <div>
        <Image src='/person.png' alt='personLogo' width={36} height={36} />
      </div>
    </header>
  );
};
