import React, { ReactNode } from 'react';

interface Props {
  children: ReactNode;
}

export const Body: React.VFC<Props> = ({ children }) => {
  return <div>{children}</div>;
};
