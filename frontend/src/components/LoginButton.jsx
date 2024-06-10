import "../styles/LoginButton.css";

const LoginButton = () => {
  const handleLogin = () => {
    window.location.href = `https://discord.com/oauth2/authorize?client_id=${
      import.meta.env.VITE_CLIENT_ID
    }&redirect_uri=${
      import.meta.env.VITE_REDIRECT_URI
    }&response_type=code&scope=identify%20email`;
  };
  return (
    <div className="container">
      <button onClick={handleLogin}>Login with Discord</button>
    </div>
  );
};

export default LoginButton;
