import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";

// createRoot().render() triggers the initial render
ReactDOM.createRoot(document.getElementById("root")!).render(
    // StrictMode calls each component's function twice during development to find impure functions
    <React.StrictMode>
        <App />
    </React.StrictMode>,
)
