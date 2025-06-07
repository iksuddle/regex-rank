import { User } from "../types/types";

export default function Profile({ user }: { user: User }) {
    return (
        <>
            <h1>{user.login}</h1>
            <p>{user.id}</p>
            <img src={user.avatar_url} width={100} height={100} />
        </>
    )
}