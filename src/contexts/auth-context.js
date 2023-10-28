import React, { createContext, useContext, useState } from "react";

// Authentication context
const AuthContext = createContext();

const useAuth = () => {
  return useContext(AuthContext);
};

const logout = (setUser) => {
    window.localStorage.removeItem("token");
    setUser({registered: false})
}

const AuthProvider = ({ children }) => {
  const [user, setUser] = useState({registered: false});

  const contextValue = {
    user,
    setUser,
    login: async (username, password) => {
        try {
            const formData = new URLSearchParams();
            formData.append("username", username);
            formData.append("password", password);
          const response = await fetch("http://localhost:8000/login", {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
            method: "POST",
            body: formData,
          }).then((response) => response.json()).catch((error) => console.log(error));
          window.localStorage.setItem("token", response.token);
          const decodedToken = JSON.parse(atob(response.token.split('.')[1]));
          setUser(decodedToken);

        } catch (error) {
            console.log(error);
            throw new Error(error);
        }

    },
    loginCheck: async () =>{
        try {
            const token = window.localStorage.getItem("token");
            if(!token){
                logout(setUser);
            }
            const response = await fetch("/loginCheck", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            }).then((response) => response.json()).catch((error) => {
                console.log(error)
                logout(setUser);
            });
            const decodedToken = JSON.parse(atob(token.split('.')[1]));
            setUser(decodedToken);
        } catch (error) {
            console.log(error);
        }
    },
  };

  return <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>;
};

export {AuthContext, useAuth}
export default AuthProvider