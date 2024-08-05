import { useState } from "react";
import Cases from "./Cases";
import Submit from "./Submit";

export default function RegexInput() {
    let [errorMessage, setErrorMessage] = useState("");
    let [allDone, setAllDone] = useState(true);
    let [userInput, setUserInput] = useState("");

    function handleInputChange(e: any) {
        setUserInput(e.target.value);
    }

    let submitButtonClassList = "submit";

    if (allDone) {
        submitButtonClassList = submitButtonClassList.concat(" done");
    }

    return (
        <>
            <p style={{
                color: "#e66c6c",
                textAlign: "right",
                height: "1.625rem",
            }}>{errorMessage}</p>
            <div className="input-div">
                <input onChange={handleInputChange} value={userInput}
                    className="text-input"
                    placeholder="enter regex here..." />
                <Submit done={allDone}/>
            </div>
            <Cases userInput={userInput} setErrorMessage={setErrorMessage} setAllDone={setAllDone} />
        </>
    )
}
