import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";
import "./index.css";
import App from "./App";
import RegexPage from "./RegexPage";
import Profile from "./Profile";

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<App />}>
                    <Route index element={<RegexPage />} />
                    <Route path="/profile" element={<Profile />} />
                </Route>
            </Routes>
        </BrowserRouter>
    </StrictMode>
);
