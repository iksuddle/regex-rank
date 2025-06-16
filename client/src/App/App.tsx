import { Outlet, Link } from "react-router";
import { FaHouse, FaUser } from "react-icons/fa6";
import "./App.css";
import { useEffect, useState } from "react";
import type { User } from "../types";

export default function App() {
    const [user, setUser] = useState<User | undefined>(undefined);

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await fetch("http://localhost:3000/me", {
                    method: "GET",
                    credentials: "include"
                });

                if (!response.ok) {
                    throw new Error(`Response status: ${response.status}`);
                }

                const user_json = (await response.json() as User);
                setUser(user_json);
            } catch (error) {
                console.log(error);
                setUser(undefined);
            }
        };

        void fetchUser();
    }, []);


    return <>
        <h1 className="center">REGEX RANK</h1>
        <p className="center">daily regex problems</p>
        <div className="content">
            <Outlet context={{ user }} />
        </div>
        <nav className="center">
            <Link to="/"><FaHouse className="icon" /></Link>
            <Link to="/profile"><FaUser className="icon" /></Link>
        </nav>
    </>;
}
