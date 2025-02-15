import Case from "./Case";
import "./Output.css";

export default function Output({ userInput, setErrorMessage, setDone }: any) {
    const casesData = [
        { literal: "hello", match: true },
        { literal: "world", match: true },
        { literal: "foo", match: false },
        { literal: "bar", match: false }
    ]

    let done = true;
    let errorMsg = "";
    let cases = casesData.map((c, index) => {
        let correct = false;
        try {
            const re = new RegExp(userInput);
            correct = re.test(c.literal);

            // if user didn't input anything then pattern shouldn't match
            if (userInput.trim().length === 0) {
                correct = false;
            }

            // invert status for ignore cases
            if (!c.match) {
                correct = !correct;
            }

            if (!correct) {
                done = false;
            }
        }
        // if there's any errors, no patterns should be matched
        catch (error: any) {
            correct = false;
            done = false;
            errorMsg = error.message;
        }

        setErrorMessage(errorMsg);
        setDone(done);
        return <Case literal={c.literal} match={c.match} done={correct} key={index} />
    });

    return (
        <div className="output">
            {cases}
        </div>
    )
}
