import Cookies from "js-cookie";
import Login from "../components/Login";

export default function Profile() {
    let jwtCookie = Cookies.get("jwt");

    return <>
        {jwtCookie ? <Profile /> : <Login />}
    </>
}
