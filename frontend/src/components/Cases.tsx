import Case from "./Case";

export default function Cases({ userInput, setErrorMessage, setAllDone }: any) {
    const cases = [
        { literal: "foo", match: true },
        { literal: "bar", match: true },
        { literal: "hello", match: false },
        { literal: "world", match: false },
    ];

    setAllDone(true);

    const listItems = cases.map((c) => {
        try {
            let re = new RegExp(userInput);

            let m = re.test(c.literal);

            if (userInput.trim().length === 0) {
                m = false;
            }

            // invert status of cases that should not be matched
            if (!c.match) {
                m = !m;
            }

            setErrorMessage("");

            // if any cases aren't finsihed, we are not done
            if (!m) {
                setAllDone(false);
            }

            return <Case literal={c.literal} done={m} match={c.match} />
        }
        catch (error: any) {
            setAllDone(false);
            setErrorMessage(error.message);
            return <Case literal={c.literal} done={false} match={c.match} />
        }
    });

    return (
        <>
            <ul className="cases">
                {listItems}
            </ul>
        </>
    )
}
