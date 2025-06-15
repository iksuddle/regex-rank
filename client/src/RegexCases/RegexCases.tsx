import { useEffect, useMemo } from "react";
import Case from "../Case";
import "./RegexCases.css";

const cases = [
    { literal: "hello", match: true },
    { literal: "world", match: true },
    { literal: "foo", match: false },
    { literal: "bar", match: false },
];

type RegexCaseProp = {
    userInput: string;
    setDone: React.Dispatch<React.SetStateAction<boolean>>;
};

export default function RegexCases({ userInput, setDone }: RegexCaseProp) {
    const { caseComponents, allCorrect } = useMemo(() => {
        let allCorrect = true;

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
                console.log(error); // handle invalid regex
            }

            return <Case literal={c.literal} match={c.match} done={correct} key={i} />;
        });

        return { caseComponents, allCorrect };
    }, [userInput]);

    useEffect(() => {
        setDone(allCorrect);
    }, [allCorrect, setDone]);

    return <div className="cases">{caseComponents}</div>;
}
