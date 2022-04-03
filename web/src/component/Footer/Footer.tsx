import React, { LegacyRef, ReactNode } from 'react';

export const Footer: React.VFC<{ children: ReactNode; _ref: React.LegacyRef<HTMLElement> }> = ({ children, _ref }) => {
  return (
    <footer
      className='fixed bottom-0 left-0 min-h-[100px] w-full bg-white py-3 px-6'
      style={{ boxShadow: '0px -1px 1px rgba(0, 0, 0, 0.12)' }}
      ref={_ref}
    >
      {children}
    </footer>
  );
};
