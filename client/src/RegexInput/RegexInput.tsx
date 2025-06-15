import { FaAngleRight } from "react-icons/fa6";

import "./RegexInput.css";

interface RegexInputProps {
    userInput: string;
    setUserInput: React.Dispatch<React.SetStateAction<string>>;
    done: boolean;
}

export default function RegexInput({ userInput, setUserInput, done }: RegexInputProps) {
    return <>
        <div className="input">
            <input onChange={e => setUserInput(e.target.value)}
                value={userInput}
                placeholder="enter regex here..."
                type="text" />
            <button disabled={!done}><FaAngleRight /></button>
        </div>
    </>;
}
