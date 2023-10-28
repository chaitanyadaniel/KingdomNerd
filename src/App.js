import React from "react";
import "./App.css";
import Main from "./components/main";
import AuthProvider from "./contexts/auth-context";

function App() {
  return (
    <AuthProvider>
      <Main />
    </AuthProvider>
  );
}

export default App;
