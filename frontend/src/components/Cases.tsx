import Case from "./Case";

export default function Cases({ userInput }: any) {
    const cases = [
        { literal: "foo", match: true },
        { literal: "bar", match: true },
        { literal: "hello", match: false },
        { literal: "world", match: false },
    ];

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

            return <Case literal={c.literal} done={m} match={c.match} />
        }
        catch (error) {
            // todo: let parent access the error (maybe using global state)
            console.log(error);
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
