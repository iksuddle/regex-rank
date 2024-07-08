import { MdBarChart, MdPerson } from "react-icons/md";
import RegexInput from "./components/RegexInput";
import "./App.css";

function App() {
    return (
        <div className="main">

            <h1>REGEX RANK</h1>
            <p>weekly regex problems</p>

            <div className="regex-input">
                <RegexInput />
            </div>

            <div className="buttons">
                <MdPerson size={40}/>
                <MdBarChart size={40}/>
            </div>
        </div>
    )
}

export default App
