import { useState } from "react";
import Output from "./Output.tsx";
import { MdArrowForwardIos } from "react-icons/md";
import "./RegexInput.css";

export default function RegexInput() {
    let [userInput, setUserInput] = useState("");
    let [errorMessage, setErrorMessage] = useState("");
    let [done, setDone] = useState(false);

    const flexCenter = {
        display: "flex",
        alignItems: "center",
        justifyContent: "center"
    }

    return (
        <div>
            <div className="error-message">{errorMessage}</div>
            <div className="input">
                <input onChange={(e) => { setUserInput(e.target.value); }}
                    value={userInput}
                    placeholder="enter regex here..."
                    autoComplete="off">
                </input>
                <button disabled={!done} style={flexCenter}><MdArrowForwardIos size={18} /></button>
            </div>
            <Output userInput={userInput} setErrorMessage={setErrorMessage} setDone={setDone} />
        </div>
    )
}
