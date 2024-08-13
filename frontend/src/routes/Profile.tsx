import Login from "../components/Login";
import User from "../components/User";
import { useEffect, useState } from "react";
import Cookies from 'js-cookie'

export default function Profile() {
    let [user, setUser] = useState({});

    let loggedIn = Cookies.get("rgx_loggedin");

    if (loggedIn) {
        useEffect(() => {
            fetch("http://localhost:3000/user", {
                method: "get",
                credentials: "include"
            })
                .then((res) => {
                    res.json().then((data) => {
                        setUser(data);
                    })
                })
                .catch((_) => {
                    setUser({})
                })
        }, []);

    }
    return <>
        {loggedIn ? <User user={user} /> : <Login />}
    </>
}
