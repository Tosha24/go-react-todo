import React from "react";
import Todos from "./Todos";

const Home = ({token}) => {  
  const logout = () => {
    localStorage.removeItem("token");
    window.location.reload();
  };

  return (
    <div className="w-full h-full">
      <nav className="fixed flex flex-row bg-teal-400 w-full top-0 justify-between p-3 px-10 md:px-20 items-center z-40">
        <div className="text-2xl font-bold">Tosha Patel</div>
        <div>
          <button
            className="border-2 border-black p-2 font-bold rounded-xl"
            onClick={logout}
          >
            Logout
          </button>
        </div>
      </nav>
      <Todos token={token}/>
    </div>
  );
};

export default Home;