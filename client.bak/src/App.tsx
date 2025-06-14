import { Link, Outlet } from "react-router";
import { FaHouse, FaUser } from "react-icons/fa6";
import { useEffect, useState } from "react";
import { User } from "./types/types";
import "./App.css";


function App() {
    let [user, setUser] = useState<User | undefined>(undefined);

    // authenticate
    useEffect(() => {
        const fetch_user = async () => {
            try {
                let response = await fetch("http://localhost:3000/me", {
                    method: "GET",
                    credentials: "include"
                });
                if (!response.ok) {
                    throw new Error(`Response status: ${response.status}`);
                }

                const user_json = await response.json();
                setUser(user_json);
            } catch (error) {
                console.log(error);
                setUser(undefined);
            }
        }

        fetch_user();
    }, []);

    return (
        <>
            <h1 className="title">REGEX RANK</h1>
            <p className="subtitle">daily regex problems</p>
            <div className="outlet">
                <Outlet context={{ user }} />
            </div>
            <div className="nav">
                <Link className="icon" to={"/"}>
                    <FaHouse />
                </Link>
                <Link className="icon" to={"/login"}>
                    <FaUser />
                </Link>
            </div>
        </>
    )
}

export default App
