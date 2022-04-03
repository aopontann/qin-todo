import React, { useState } from 'react';
import { useTodoContext } from 'context/TodoContext';

export const TodoInput = () => {
  const { todosList, text, setText, ref, input, addTask, currentTask, setCurrentTask, updateTask } = useTodoContext();
  const [onClick, setOnclick] = useState(false);
  const [current, setCurrent] = useState(false);
  const handleBlur = () => {
    if (onClick) return false;
    setCurrent(false);
    setCurrentTask('');
    setText('');
  };

  return (
    <div className='grid gap-2.5'>
      <input
        className='rounded-2xl bg-surface py-3 px-4 text-sm'
        value={text}
        onChange={input}
        onFocus={() => setCurrent(true)}
        onBlur={handleBlur}
        ref={ref}
        placeholder='input placeholder'
        type='text'
      />
      <div className='grid grid-flow-col gap-1'>
        {current &&
          todosList.map((todosItem, index) => (
            <button
              className={`${todosItem.color} grid items-center rounded-full bg-current py-2.5`}
              onMouseDown={() => setOnclick(true)}
              onMouseUp={() => setOnclick(false)}
              onClick={() => (currentTask !== '' ? updateTask(todosItem.label) : addTask(todosItem.label))}
              key={index}
            >
              <span className='text-xs font-bold text-white'>{todosItem.label}</span>
            </button>
          ))}
      </div>
    </div>
  );
};
