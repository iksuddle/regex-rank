export default function Case({ ignore, literal, done }: any) {
    let matchColor = "#a4cf9b";
    let ignoreColor = "#e8a7a7";

    if (done) {
        matchColor = "#65ad55";
        ignoreColor = "#e66c6c";
    }

    return (
        <li className="case">
            <span className="match-text" style={ignore ? { color: ignoreColor } : { color: matchColor }}>
                {ignore ? "ignore" : "match"}
            </span>
            <p style={done ? { color: "black" } : { color: "grey" }}>
                {literal}
            </p>
            <span className="match-color" style={ignore ? { background: ignoreColor } : { background: matchColor }} />
        </li>
    )
}
