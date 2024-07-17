import React, { useEffect, useState } from "react";
import Login from "./Login";
import Home from "./Home";

const isAuthenticated = () => {
  const token = localStorage.getItem("token");
  if (!token) {
    return "";
  }
  return token; 
};

const PrivateRoute = () => {
  const [tokenString, setTokenString] = useState("");

  useEffect(() => {
    const getTokenString = () => {
      if(tokenString === ""){
        setTokenString(isAuthenticated())
      }
    }
    
    getTokenString()
  }, [tokenString])
  return tokenString === "" ? <Login/> : <Home token={tokenString}/>;
};

export default PrivateRoute;