import { useEffect, useState } from "react";
import RegexCases from "../RegexCases";
import RegexInput from "../RegexInput";
import "./RegexPage.css";

export default function RegexPage() {
    const [userInput, setUserInput] = useState(() => localStorage.getItem("user-input") || "");
    const [done, setDone] = useState(false);
    const [errorMsg, setErrorMsg] = useState("");

    useEffect(() => {
        localStorage.setItem("user-input", userInput);
    }, [userInput]);

    return <>
        <p className="error-message">{errorMsg}</p>
        <RegexInput userInput={userInput} setUserInput={setUserInput} done={done} />
        <RegexCases userInput={userInput} setDone={setDone} setErrorMsg={setErrorMsg} />
    </>;
}
