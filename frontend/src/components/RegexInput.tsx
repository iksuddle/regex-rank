import { useState } from "react";
import Cases from "./Cases";
import { MdArrowForwardIos } from "react-icons/md";

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
            <div className="input-div">
                <input onChange={handleInputChange} value={userInput}
                    className="text-input"
                    placeholder="enter regex here..." />
                <button className="submit"><MdArrowForwardIos size={20} color="white"/></button>
            </div>
            <Cases userInput={userInput} />
        </>
    )
}
