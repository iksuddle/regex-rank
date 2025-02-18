import "./Case.css";

export default function Case({ literal, match, done }: any) {
    let matchColor = "#b8e0af";
    let ignoreColor = "#ebb7b7";
    let literalColor = "#707070";

    if (done) {
        matchColor = "#65ad55";
        ignoreColor = "#e34d4d";
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
