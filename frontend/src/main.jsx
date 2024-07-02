import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.jsx";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Login from "./components/LoginPage.jsx";
import About from "./components/About.jsx";
import "./styles/index.css";
import Dashboard from "./components/Dashboard.jsx";
import SearchResults from "./components/SearchResults.jsx";
import PrivacyPolicy from "./components/PrivacyPolicy.jsx";
import TermsOfService from "./components/TermsOfService.jsx";
import Home from "./components/Home.jsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      { path: "/about", element: <About /> },
      { path: "/", element: <Home /> },
      { path: "dashboard", element: <Dashboard /> },
      { path: "search/:type", element: <SearchResults /> },
      { path: "privacy", element: <PrivacyPolicy /> },
      { path: "terms", element: <TermsOfService /> },
    ],
  },
  {
    path: "/login",
    element: <Login />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
