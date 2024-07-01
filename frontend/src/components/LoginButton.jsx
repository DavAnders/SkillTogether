const LoginButton = () => {
  const link = `https://discord.com/oauth2/authorize?client_id=${
    import.meta.env.VITE_CLIENT_ID
  }&redirect_uri=${
    import.meta.env.VITE_REDIRECT_URI
  }&response_type=code&scope=identify%20email`;

  return (
    <div>
      <a
        className="dark:bg-slate-200 dark:text-slate-900 dark:hover:bg-slate-200/90 ring-offset-background focus-visible:ring-ring flex h-10 w-full items-center justify-center whitespace-nowrap rounded-md bg-black px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-black/90 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50"
        href={link}
      >
        Login with Discord
      </a>
    </div>
  );
};

export default LoginButton;
