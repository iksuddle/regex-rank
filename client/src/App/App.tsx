import { Outlet, Link } from "react-router";
import { FaHouse, FaUser } from "react-icons/fa6";
import "./App.css";

export default function App() {
  return <>
    <h1 className="center">REGEX RANK</h1>
    {/* <p className="center">daily regex problems</p> */}
    <div className="content">
      <Outlet />
    </div>
    <nav className="center">
      <Link to="/"><FaHouse className="icon" /></Link>
      <Link to="/profile"><FaUser className="icon" /></Link>
    </nav>
  </>;
}