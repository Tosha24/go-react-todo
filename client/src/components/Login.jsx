import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.id]: e.target.value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    axios
      .post(`${import.meta.env.VITE_SERVER_URL}/api/login`, formData)
      .then((response) => {
        localStorage.setItem("token", response.data.token);
        navigate("/");
        window.location.reload();
      })
      .catch((error) => {
        alert(error.response.data.error || "Cannot Login User");
      })
      .finally(() => {
        setFormData({
          email: "",
          password: "",
        });
      });
  };

  return (
    <div className="flex justify-center items-center h-screen w-full">
      <div className="bg-teal-200 p-8 rounded shadow-md w-96 border-4 border-teal-900">
        <h1 className="text-2xl font-bold mb-6 flex justify-center">Login</h1>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label
              className="block text-gray-800 text-sm font-bold mb-2"
              htmlFor="email"
            >
              Email
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-teal-800 focus:shadow-outline"
              id="email"
              type="email"
              placeholder="abc@xyz.com"
              value={formData.email}
              onChange={handleChange}
            />
          </div>
          <div className="mb-6">
            <label
              className="block text-gray-800 text-sm font-bold mb-2"
              htmlFor="password"
            >
              Password
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-teal-800 focus:shadow-outline"
              id="password"
              type="password"
              placeholder="**********"
              value={formData.password}
              onChange={handleChange}
            />
          </div>
          <div className="flex items-center justify-between">
            <button
              className="bg-teal-800 hover:bg-teal-900 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
              type="submit"
            >
              Sign In
            </button>
          </div>
        </form>

        <div className="flex justify-center mt-4 text-gray-800 text-lg font-bold">
          Don't have an Account? &nbsp;
          <a className="hover:underline" href="/sign-up">
            {" "}
            Register
          </a>
        </div>
      </div>
    </div>
  );
};

export default Login;
