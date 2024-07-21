import Case from "./Case";

export default function Cases({ userInput }: any) {
    const cases = [
        { literal: "foo", match: true },
        { literal: "bar", match: true },
        { literal: "hello", match: false },
        { literal: "world", match: false },
    ];

    const listItems = cases.map((c) => {
        let re = new RegExp(userInput);

        let m = re.test(c.literal);
        if (userInput.trim().length === 0) {
            m = false;
        }

        return <Case literal={c.literal} done={m} match={c.match} />
    });

    return (
        <>
            <ul className="cases">
                {listItems}
            </ul>
        </>
    )
}
