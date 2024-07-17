import axios from "axios";
import React, { useEffect, useState } from "react";
import { MdOutlineDelete } from "react-icons/md";
import { FiEdit } from "react-icons/fi";
import toast from "react-hot-toast";

const Todos = ({ token }) => {
  const [title, setTitle] = useState("");
  const [allTodos, setAllTodos] = useState([]);

  useEffect(() => {
    getAllTodos();
  }, [token]);

  function getAllTodos() {
    axios
      .get(`${import.meta.env.VITE_SERVER_URL}/api/todos`, {
        headers: {
          Authorization: `${token}`,
        },
      })
      .then((response) => {
        setAllTodos(response.data.todos);
      })
      .catch((error) => {
        toast.error(error.response.data.error || "Cannot Fetch todos! Try again...")
      });
  }

  const addTodo = () => {
    if (title === "") return;
    axios
      .post(
        `${import.meta.env.VITE_SERVER_URL}/api/todo`,
        {
          Title: title,
          Completed: false,
        },
        {
          headers: {
            Authorization: `${token}`,
          },
        }
      )
      .then((response) => {
        toast.success("Todo added successfully!");
        setAllTodos([...allTodos, response.data.todo]);
      })
      .catch((error) => {
        toast.error(error.response.data.error || "Cannot add Todo! Try again...")
      })
      .finally(() => {
        setTitle("");
      });
  };

  const toggleComplete = (id, value) => {
    axios
      .put(
        `${import.meta.env.VITE_SERVER_URL}/api/todo/mark/${id}`,
        {
          Completed: value,
        },
        {
          headers: {
            Authorization: `${token}`,
          },
        }
      )
      .then(() => {
        toast.success("Todo Updated Successfully!");
        getAllTodos();
      })
      .catch((error) => {
        console.error(error);
        toast.error(error.response.data.error || "Error occured!")
      });
  };

  const updateTodo = (id) => {
    let currentTodoTitle = "";
    // get currenttodo from id
    allTodos.forEach((todo) => {
      if(todo.id === id){
        currentTodoTitle = todo.title
      }
    })  

    const newTitle = prompt("Enter new title", currentTodoTitle);
    if (newTitle === "" || newTitle === null) return;
    axios
      .put(
        `${import.meta.env.VITE_SERVER_URL}/api/todo/${id}`,
        {
          Title: newTitle,
        },
        {
          headers: {
            Authorization: `${token}`,
          },
        }
      )
      .then(() => {
        toast.success("Todo updated Successfully!")
        getAllTodos();
      })
      .catch((error) => {
        console.error(error);
        toast.error(error.response.data.error || "Cannot Update Todo! Try again...")
      });
  };

  const deleteTodo = (id) => {
    axios
      .delete(`${import.meta.env.VITE_SERVER_URL}/api/todo/${id}`, {
        headers: {
          Authorization: `${token}`,
        },
      })
      .then(() => {
        toast.success("Todo deleted successfully!");
        getAllTodos();
      })
      .catch((error) => {
        console.error(error);
        toast.error(error.response.data.error || "Cannot Delete Todo! Try again...")
      });
  };

  return (
    <div className="mt-16 w-[100%]">
      <div className="w-[100%] md:w-[60%] mx-auto items-center justify-center flex mt-10 flex-col gap-5 p-5">
        <h1 className="text-2xl font-bold">My Todos</h1>
        <div className="gap-2 flex">
          <input
            type="text"
            placeholder="Add your task..."
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            autoFocus
            className="p-2 w-[200px] sm:w-[400px] rounded-md border-2 border-teal-950 focus:outline-teal-950"
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                addTodo();
              }
            }}
          />
          <button
            className="bg-teal-900 text-white p-2 rounded"
            onClick={addTodo}
          >
            Add
          </button>
        </div>
        <div className="w-full p-3 items-center flex flex-col justify-center gap-3">
          {allTodos && allTodos.length > 0 ? (
            allTodos.map((todo, index) => (
              <div
                key={todo.id}
                className="flex justify-between items-center w-full border border-teal-800 rounded-xl p-4"
              >
                <div className="flex items-center gap-3">
                  <span className="font-bold">{index + 1}.</span>
                  <span
                    className={`${
                      todo.completed ? "line-through" : ""
                    } hidden sm:block text-xl`}
                  >
                    {todo.title}
                  </span>
                  <span
                    className={`${
                      todo.completed ? "line-through" : ""
                    } sm:hidden text-xl text-ellipsis`}
                  >
                    {todo.title.slice(0, 6)}{todo.title.length > 6 ? "..." : ""}
                  </span>
                </div>
                <div className="flex items-center gap-3">
                  <input
                    type="checkbox"
                    checked={todo.completed}
                    onChange={(e) => toggleComplete(todo.id, e.target.checked)}
                    className="cursor-pointer w-5 h-5"
                  />
                  <button
                    className="text-red-500 text-2xl"
                    onClick={() => updateTodo(todo.id)}
                  >
                    <FiEdit />
                  </button>
                  <button
                    className="text-red-500 text-2xl"
                    onClick={() => deleteTodo(todo.id)}
                  >
                    <MdOutlineDelete />
                  </button>
                </div>
              </div>
            ))
          ) : (
            <div className="text-gray-500 text-center">No todos found</div>
          )}
        </div>
      </div>
    </div>
  );
};

export default Todos;
