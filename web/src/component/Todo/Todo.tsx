import React, { useState } from 'react';

interface Props {
  todo: string;
}

export const Todo: React.VFC<Props> = ({ todo }) => {
  const [active, setActive] = useState(false);

  return (
    <div>
      <input type='checkbox' onChange={() => setActive((prev) => !prev)} />
      <button className={`${active ? 'line-through' : undefined}`}>{todo}</button>
    </div>
  );
};
