import "./Login.css";

export default function Login() {
    function onLoginButtonClick() {
        window.location.href = "http://localhost:3000/login";
    }

    return (
        <div className="login">
            <h2>Login</h2>
            <div className="login-btn">
                <button onClick={onLoginButtonClick}>Login with GitHub</button>
            </div>
        </div>
    )
}
