import { FaAngleRight } from "react-icons/fa6";

import "./RegexInput.css";

export default function RegexInput() {
    return <>
        <div className="input">
            <input type="text" placeholder="enter regex here..." />
            <button><FaAngleRight /></button>
        </div>
    </>;
}