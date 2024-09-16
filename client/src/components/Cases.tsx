import { useEffect } from "react";
import Case from "./Case";

export default function Cases({ userInput, setErrorMessage, setAllDone }: any) {
    const casesData = [
        { literal: "foo", ignore: false },
        { literal: "bar", ignore: false },
        { literal: "hello", ignore: true },
        { literal: "world", ignore: true },
    ];

    let errorMessage = "";
    let done = true;

    const listItems = casesData.map((c, index) => {
        try {
            const re = new RegExp(userInput);

            let patternFound = re.test(c.literal);

            // if the user didn't enter anything then the pattern shouldn't match
            if (userInput.trim().length === 0) {
                patternFound = false;
            }

            // invert status of cases that should be ignored
            if (c.ignore) {
                patternFound = !patternFound;
            }

            // if any aren't matching, we are not done
            if (!patternFound) {
                done = false;
            }

            return <Case literal={c.literal} ignore={c.ignore} done={patternFound} key={index} />
        }
        // if there's any errors, no pattern should be matched
        catch (error: any) {
            done = false;
            errorMessage = error.message;
            return <Case literal={c.literal} ignore={c.ignore} done={false} key={index} />
        }
    });

    useEffect(() => {
        setErrorMessage(errorMessage);
        setAllDone(done);
    });

    return (
        <ul className="cases">
            {listItems}
        </ul>
    )
}
