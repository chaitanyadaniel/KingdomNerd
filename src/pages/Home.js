import React, { useEffect, useState } from "react";
import { useAuth } from "../contexts/auth-context";

const HomePage = () => {
  const [serverResponse, setServerResponse] = useState();
  const auth = useAuth();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("/home", {
          method: "GET",
          headers: {
            Authorization: `Bearer ${window.localStorage.getItem("token")}`,
          },
        });
        const body = await response.text();
        setServerResponse(body);
      } catch (error) {
        setServerResponse("Server is not running");
      }
    };
    fetchData();
  }, [auth.user]);

  return (
    <div>
      {serverResponse && auth.user.authorized ? (
        <span>
          {" "}
          The server says:{" "}
          <code>
            {serverResponse} {auth.user.first_name} {auth.user.last_name}
          </code>
        </span>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default HomePage;
