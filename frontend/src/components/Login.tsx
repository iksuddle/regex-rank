export default function Login() {
    async function handleLogin() {
        const url = "http://localhost:3000/login"
        try {
            const response = await fetch(url);
            if (!response.ok) {
                throw new Error(`Response status: ${response.status}`);
            }
            const json = await response.json();
            console.log(json);
        }
        catch (error: any) {
            console.error(error.message);
        }
    }

    return <>
        <div style={{ display: "flex", justifyContent: "center", alignItems: "center", height: "80%" }}>
            <button onClick={handleLogin} className="action-button">Login with GitHub</button>
        </div>
    </>
}
