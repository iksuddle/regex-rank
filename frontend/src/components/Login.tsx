export default function Login() {
    async function handleLogin() {
        window.location.href = "http://localhost:3000/login"
    }

    return <>
        <div style={{ display: "flex", justifyContent: "center", alignItems: "center", height: "80%" }}>
            <button onClick={handleLogin} className="action-button">Login with GitHub</button>
        </div>
    </>
}
