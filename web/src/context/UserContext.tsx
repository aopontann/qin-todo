import { createContext, useContext, useState } from 'react';

type User = {
  avatar_url: string;
  email: string;
  id: string;
  name: string;
};

type DafeultValue = {
  user: User;
  setUser: React.Dispatch<
    React.SetStateAction<{
      avatar_url: string;
      email: string;
      id: string;
      name: string;
    }>
  >;
};

type Props = {
  children: React.ReactNode;
};

export const UserContext = createContext<DafeultValue | null>(null);
export const useUserContext = () => {
  const context = useContext(UserContext);
  if (context === null) {
    throw new Error('useCount must be used within a CountProvider');
  }
  return context;
};
export const UserContextProvider: React.VFC<Props> = ({ children }) => {
  const [user, setUser] = useState({
    avatar_url: '',
    email: '',
    id: '',
    name: '',
  });
  return <UserContext.Provider value={{ user, setUser }}>{children}</UserContext.Provider>;
};
