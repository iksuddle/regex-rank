import { MdBarChart, MdPerson } from "react-icons/md";
import { Outlet, Link } from "react-router-dom";

export default function App() {
    return (
        <div className="main">

            <Link style={{ textDecoration: "none", color: "#333333" }} to={``}><h1>REGEX RANK</h1></Link>
            <p>weekly regex problems</p>

            <div style={{ height: "15rem" }} className="regex-input">
                <Outlet />
            </div>

            <div className="buttons">
                <Link to={`login`}>
                    <MdPerson size={40} />
                </Link>
                <MdBarChart size={40} />
            </div>
        </div>
    )
}
