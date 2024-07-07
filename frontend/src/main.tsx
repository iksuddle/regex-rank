import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")!).render(
    // StrictMode calls each component's function twice during development to find impure functions
    <React.StrictMode>
        <App />
    </React.StrictMode>,
)
