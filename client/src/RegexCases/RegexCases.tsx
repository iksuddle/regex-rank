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
}

export default function RegexCases({ userInput, setDone }: RegexCaseProp) {
    let done = true;
    const caseComponents = cases.map((c, i) => {
        let correct = false;
        try {
            const re = RegExp(userInput);
            correct = re.test(c.literal);

            // if userInput is empty, it is not correct
            if (userInput.trim().length == 0) {
                correct = false;
            }

            // invert correctness for ignore cases
            if (!c.match) {
                correct = !correct;
            }

            if (!correct) {
                done = false;
            }
        } catch (error: unknown) {
            correct = false;
            setDone(false);
            // todo: send error message
            console.log(error);
        }

        setDone(done);

        return <Case literal={c.literal} match={c.match} done={correct} key={i} />;
    });

    return <div className="cases">
        {caseComponents}
    </div>;
}