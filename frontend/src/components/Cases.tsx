import Case from "./Case";

export default function Cases() {
    return (
        <>
            <ul className="cases">
                <Case literal="foo" done={true} match={true}/>
                <Case literal="bar" done={false} match={true}/>
                <Case literal="hello" done={true} match={false}/>
                <Case literal="world" done={false} match={false}/>
            </ul>
        </>
    )
}
