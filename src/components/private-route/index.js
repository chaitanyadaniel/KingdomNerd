import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../contexts/auth-context";

const PrivateRoute = ({ children, ...rest }) => {
  let { user } = useAuth();
  let navigate = useNavigate();

  useEffect(() => {
    if (!user || !user.authorized) {
        navigate("/login");
      }
  }, [user]);

  return children;
};

export default PrivateRoute;
