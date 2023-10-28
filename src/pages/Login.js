import React, { useEffect } from 'react';
import { useAuth } from '../contexts/auth-context';
import { useNavigate } from 'react-router-dom';

const LoginPage = () => {
  const auth = useAuth();
  const user = auth.user;
  const navigate = useNavigate();

  useEffect(() => {

  }, [auth.user]);

  const handleSubmit = async (event) => {
    event.preventDefault();
    const { username, password } = event.target.elements;
    await auth.login(username.value, password.value);
  };

  return (
    <>
      {user.authorized ? (
        navigate('/home')
      ) : (
        <form onSubmit={handleSubmit}>
          <input name="username" type="text" placeholder="Username" />
          <input name="password" type="password" placeholder="Password" />
          <button type="submit">Login</button>
        </form>
      )}
    </>
  );

}
export default LoginPage;
