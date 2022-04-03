import React from 'react';
import { useTodoContext } from 'context/TodoContext';

type Props = {
  todo: {
    id: string;
    content: string;
    completed: boolean;
    execution_date: string | null;
  };
};

export const Todo: React.VFC<Props> = ({ todo }) => {
  const { removeTask, completedTask, currentTask, editTask } = useTodoContext();

  return (
    <div
      className={`${
        currentTask === todo.id ? 'bg-[#FBBF24]/10' : 'bg-transparent'
      } -mx-2 grid grid-cols-[max-content_1fr_max-content] items-center gap-3 px-2 transition-colors`}
    >
      <div className='h-6 w-6 rounded-full border-2 border-gray p-[2px]'>
        <div
          className={`h-full w-full rounded-full transition-colors ${todo.completed ? 'bg-current' : 'bg-transparent'}`}
          onClick={() => completedTask(todo.id)}
        ></div>
      </div>
      <button
        className={`${
          todo.completed ? ' text-gray line-through' : 'text-black'
        } items-stretch justify-self-start text-left font-normal transition-colors`}
        onClick={() => editTask(todo.id, todo.content)}
      >
        {todo.content}
      </button>
      <button className='-mr-6 h-full bg-red px-5 py-2 text-white' onClick={() => removeTask(todo.id)}>
        削除
      </button>
    </div>
  );
};
