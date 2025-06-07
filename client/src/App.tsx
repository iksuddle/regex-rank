import { Link, Outlet } from "react-router";
import "./App.css";
import { FaHouse, FaUser } from "react-icons/fa6";


function App() {
    return (
        <>
            <h1 className="title">REGEX RANK</h1>
            <p className="subtitle">daily regex problems</p>
            <div className="outlet">
                <Outlet />
            </div>
            <div className="nav">
                <Link className="icon" to={"/"}>
                    <FaHouse />
                </Link>
                <Link className="icon" to={"/login"}>
                    <FaUser />
                </Link>
            </div>
        </>
    )
}

export default App
