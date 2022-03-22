import React from 'react';
import { useTodoContext } from 'context/TodoContext';

export const TodosList = () => {
  const { todosList, ref, removeTask, completedTask, currentTask, editTask } = useTodoContext();

  return (
    <div className='grid gap-8'>
      {todosList.map((todosItem, i) => (
        <section key={i}>
          <p className={todosItem.color}>{todosItem.label}</p>
          <ul>
            {todosItem.todos.map((todo, i) => (
              <li className={currentTask.id === todo.id ? 'bg-[#FBBF24]/10' : ''} key={i}>
                <input
                  onChange={() => completedTask(todo.id)}
                  checked={todo.completed}
                  type='checkbox'
                />
                <button onClick={() => editTask(todo.id, todo.content)}>{todo.content}</button>
                <button onClick={() => removeTask(todo.id)}>削除</button>
              </li>
            ))}
            <li className={todosItem.todos.length ? 'hidden md:block' : ''}>
              <button onClick={() => ref.current?.focus()}>タスクを追加する</button>
            </li>
          </ul>
        </section>
      ))}
    </div>
  );
};
