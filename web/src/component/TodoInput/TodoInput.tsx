import React from 'react';
import { useTodoContext } from 'context/TodoContext';

export const TodoInput = () => {
  const { todosList, text, ref, input, addTask, currentTask, updateTask } = useTodoContext();

  return (
    <div className='grid gap-2 py-8'>
      <input className='border border-gray' value={text} onChange={input} ref={ref} type='text' />
      <div className='grid gap-1 grid-flow-col'>
        {todosList.map((todosItem, index) => (
          <button
            className={`${todosItem.bg} text-white`}
            onClick={() => {
              currentTask.active ? updateTask(todosItem.label) : addTask(todosItem.label);
            }}
            key={index}
          >
            {todosItem.label}
          </button>
        ))}
      </div>
    </div>
  );
};
