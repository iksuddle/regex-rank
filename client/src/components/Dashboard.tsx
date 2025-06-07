import { useOutletContext } from "react-router";
import Login from "./Login";
import Profile from "./Profile";
import {User} from "../types/types";


type ContextType = {
    user: User | undefined;
};

export default function Dashboard() {
    const { user } = useOutletContext<ContextType>();

    return (
        <>
            {user ? (
                <Profile user={user} />
            ) : (
                <Login />
            )}
        </>
    );
}
