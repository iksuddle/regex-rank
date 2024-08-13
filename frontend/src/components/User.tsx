export default function User({ user }: any) {
    let created = new Date(user.created_at).toLocaleDateString();

    return <>
        <div style={{ display: "flex" }}>
            <img src={user.avatar_url} style={{ width: "100px", height: "100px", borderRadius: "50%" }} />
            <div style={{ marginLeft: "1rem" }}>
                <h1>{user.username}</h1>
                <p>Account created on {created}</p>
                <button
                    onClick={() => {
                        fetch("http://localhost:3000/logout", {
                            method: "get",
                            credentials: "include"
                        }).finally(() => {
                            setTimeout(() => {
                                window.location.reload();
                            })
                        });
                    }}
                    className="action-button" style={{ fontWeight: "400", marginTop: "0.5rem" }}>
                    logout
                </button>
            </div>
        </div>
    </>
}
