import React, { useEffect } from "react";
import {
  Navigate,
  Route,
  BrowserRouter as Router,
  Routes,
} from "react-router-dom";
import { useAuth } from "../../contexts/auth-context";
import LoginPage from "../../pages/Login";
import PrivateRoute from "../private-route";
import HomePage from "../../pages/Home";

function Main() {
  const auth = useAuth();

  useEffect(() => {
    const reSign = async () => {
      await auth.loginCheck();
    };
    reSign();
  }, []);

  return (
    <>
      <Router>
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route
            path="/home"
            element={
              <PrivateRoute>
                <HomePage />
              </PrivateRoute>
            }
          />
          <Route path="/" element={<Navigate to="/home" />} />
        </Routes>
      </Router>
    </>
  );
}

export default Main;
