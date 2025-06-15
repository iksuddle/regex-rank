import { useEffect } from "react";
import Case from "../Case";
import "./RegexCases.css";

const cases = [
    { literal: "hello", match: true },
    { literal: "world", match: true },
    { literal: "foo", match: false },
    { literal: "bar", match: false },
];

interface RegexCaseProp {
    userInput: string;
    setDone: React.Dispatch<React.SetStateAction<boolean>>;
    setErrorMsg: React.Dispatch<React.SetStateAction<string>>;
};

export default function RegexCases({ userInput, setDone, setErrorMsg }: RegexCaseProp) {
    let allCorrect = true;
    let errorMsg = "";

    const caseComponents = cases.map((c, i) => {
        let correct = false;
        try {
            const re = RegExp(userInput);
            correct = re.test(c.literal);

            if (userInput.trim().length === 0) {
                correct = false;
            }

            if (!c.match) {
                correct = !correct;
            }

            if (!correct) {
                allCorrect = false;
            }
        } catch (error: unknown) {
            correct = false;
            allCorrect = false;
            if (error instanceof Error) {
                errorMsg = error.message;
            } else {
                errorMsg = "An unknown error occurred.";
            }
        }

        return <Case literal={c.literal} match={c.match} done={correct} key={i} />;
    });

    useEffect(() => {
        setErrorMsg(errorMsg);
    }, [errorMsg, setErrorMsg]);


    useEffect(() => {
        setDone(allCorrect);
    }, [allCorrect, setDone]);

    return <div className="cases">{caseComponents}</div>;
}
