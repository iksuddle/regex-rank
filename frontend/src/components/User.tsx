export default function User({ user }: any) {
    let created = new Date(user.created_at).toLocaleDateString(undefined, {
        month: "long",
        year: "numeric",
        weekday: "long",
        day: "numeric"
    });

    async function handleUserLogout() {
        let req = new Request(
            "http://localhost:3000/logout",
            { method: "get", credentials: "include" } // include credentials so jwt cookie gets sent
        );
        try {
            const res = await fetch(req)
            if (!res.ok) {
                throw new Error("error logging out: " + res.status.toString());
            }

            window.location.reload();
            localStorage.removeItem("rgx_user");
        }
        catch (error: any) {
            console.log(error.message);
        }
    }

    return <>
        <div style={{ display: "flex" }}>
            <img src={user.avatar_url} style={{ width: "100px", height: "100px", borderRadius: "50%" }} />
            <div style={{ marginLeft: "1rem" }}>
                <h1>{user.username}</h1>
                <p>Account created on {created}</p>
                <button
                    onClick={handleUserLogout}
                    className="action-button" style={{ fontWeight: "400", marginTop: "0.5rem" }}>
                    logout
                </button>
            </div>
        </div>
    </>
}
