import { useState } from "react";
import Cases from "./Cases";

export default function RegexInput() {
    let [userInput, setUserInput] = useState("");

    function handleInputChange(e: any) {
        setUserInput(e.target.value);
    }

    return (
        <>
            <p style={{
                color: "#e66c6c",
                textAlign: "right",
            }}>error message here</p>
            <input onChange={handleInputChange} value={userInput}
                className="text-input"
                placeholder="enter regex here..." />
            <Cases userInput={userInput} />
        </>
    )
}
