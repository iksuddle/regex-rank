import "./Case.css";

export default function Case({ literal, match, done }: any) {
    let matchColor = "#a4cf9b";
    let ignoreColor = "#e8a7a7";
    let literalColor = "#707070";

    if (done) {
        matchColor = "#65ad55";
        ignoreColor = "#e66c6c";
        literalColor = "#0a0a0a";
    }

    let color = (m: boolean) => {
        return m ? matchColor : ignoreColor;
    }

    return (
        <li className="case">
            <p className="match-text" style={{ color: color(match) }}>{match ? "match" : "ignore"}</p>
            <p className="literal" style={{ color: literalColor }}>{literal}</p>
            <span style={{ background: color(match) }}></span>
        </li>
    )
}
