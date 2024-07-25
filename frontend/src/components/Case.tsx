export default function Case({ match, literal, done }: any) {
    let matchColor = "#a4cf9b";
    let ignoreColor = "#e8a7a7";

    if (done) {
        matchColor = "#65ad55";
        ignoreColor = "#e66c6c";
    }

    return (
        <li className="case">
            <span className="match-text" style={match ? { color: matchColor } : { color: ignoreColor }}>
                {match ? "match" : "ignore"}
            </span>
            <p style={done ? { color: "black" } : { color: "grey" }}>
                {literal}
            </p>
            <span className="match-color" style={match ? { background: matchColor } : { background: ignoreColor }} />
        </li>
    )
}
