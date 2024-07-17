import React from "react";
import {
  createBrowserRouter,
  createRoutesFromChildren,
  Route,
  RouterProvider,
} from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import PrivateRoute from "./components/PrivateRoute";

const router = createBrowserRouter(
  createRoutesFromChildren(
    <>
      <Route path="/" element={<PrivateRoute/>}/>
      <Route path="/sign-in" element={<Login />} />
      <Route path="/sign-up" element={<Register />} />
    </>
  )
);

const App = () => {
  return <div className="min-h-screen max-w-screen overflow-y-auto overflow-x-hidden">
    <RouterProvider router={router} />
  </div>;
};

export default App;
