import React, { LegacyRef, ReactNode } from 'react';

export const Footer: React.VFC<{ children: ReactNode; _ref: React.LegacyRef<HTMLElement> }> = ({ children, _ref }) => {
  return (
    <footer className='shadow-t fixed bottom-0 left-0 min-h-[100px] w-full bg-white py-3 px-6' ref={_ref}>
      {children}
    </footer>
  );
};
