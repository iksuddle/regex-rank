import { useEffect, useState } from "react";
import RegexCases from "../RegexCases";
import RegexInput from "../RegexInput";

export default function RegexPage() {
  const [userInput, setUserInput] = useState(() => localStorage.getItem("user-input") || "");
  const [done, setDone] = useState(false);

  useEffect(() => {
    localStorage.setItem("user-input", userInput);
  }, [userInput]);

  return <>
    <RegexInput userInput={userInput} setUserInput={setUserInput} done={done} />
    <RegexCases userInput={userInput} setDone={setDone} />
  </>;
}