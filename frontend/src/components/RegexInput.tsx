import { useState } from "react";
import Cases from "./Cases";

export default function RegexInput() {
    let [userInput, setUserInput] = useState("");

    function handleInputChange(e: any) {
        setUserInput(e.target.value);
    }

    return (
        <>
            <input onChange={handleInputChange} value={userInput}
                className="text-input"
                placeholder="enter regex here..." />
            <Cases userInput={userInput} />
        </>
    )
}
