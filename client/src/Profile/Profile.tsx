import { useOutletContext } from "react-router";
import type { User } from "../types";

type ContextType = {
    user: User | undefined;
};


export default function Profile() {
    const { user } = useOutletContext<ContextType>();

    function onLoginButtonClick() {
        window.location.href = "http://localhost:3000/login";
    }
    function onLogoutButtonClick() {
        window.location.href = "http://localhost:3000/logout";
    }

    return <>
        {user ? (
            <>
                <h1>{user.login}</h1>
                <p>{user.id}</p>
                <img src={user.avatar_url} width={100} height={100} />
                <button onClick={onLogoutButtonClick}>Logout</button>
            </>
        ) : (
            <>
                <h2>Login</h2>
                <button onClick={onLoginButtonClick}>Login with GitHub</button>
            </>
        )}
    </>;

}
