import { createBrowserRouter, RouterProvider } from "react-router-dom";
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import ErrorPage from "./error-page.tsx";
import Profile from "./routes/profile.tsx";
import "./index.css";
import RegexInput from "./components/RegexInput.tsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App />,
        errorElement: <ErrorPage />,
        children: [
            {
                path: "/login",
                element: <Profile />,
            },
            {
                path: "/",
                element: <RegexInput />,
            }
        ]
    }
]);


// createRoot().render() triggers the initial render
ReactDOM.createRoot(document.getElementById("root")!).render(
    // StrictMode calls each component's function twice during development to find impure functions
    <React.StrictMode>
        <RouterProvider router={router} />
    </React.StrictMode>,
)
