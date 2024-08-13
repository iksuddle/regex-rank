import Cookies from "js-cookie";
import Login from "../components/Login";
import User from "../components/User";
import { useSearchParams } from "react-router-dom";

export default function Profile() {
    let [params, _setParams] = useSearchParams();

    if (params.size > 0) {
        let _id = params.get("id")
        // todo: get user by id
    }

    let loggedInCookie = Cookies.get("rgx_loggedin");
    console.log(loggedInCookie)

    return <>
        {loggedInCookie ? <User /> : <Login />}
    </>
}
