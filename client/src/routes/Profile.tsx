import { useEffect, useState } from "react"
import Login from "../components/Login";
import User from "../components/User";

// todo: move useEffect to utils
export default function Profile() {
    let [user, setUser] = useState(null);

    useEffect(() => {
        const getUser = async () => {
            let req = new Request(
                "http://localhost:3000/user",
                { method: "get", credentials: "include" } // include credentials so jwt cookie gets sent
            );

            try {
                const res = await fetch(req);
                if (!res.ok) {
                    throw new Error("error retrieving user: " + res.status.toString());
                }
                const userData = await res.json();
                setUser(userData)
                localStorage.setItem("rgx_user", JSON.stringify(userData));
            }
            catch (error: any) {
                setUser(null);
                localStorage.removeItem("rgx_user");
            }
        }

        let userDataString = localStorage.getItem("rgx_user")

        if (userDataString) {
            console.log("user not found");
            setUser(JSON.parse(userDataString));
        } else {
            console.log("user not found");
            getUser();
        }
    }, []);

    return <>
        {user ? <User user={user} /> : <Login />}
    </>
}
