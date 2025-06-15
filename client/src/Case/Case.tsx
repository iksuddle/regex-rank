import "./Case.css";

type CaseProps = {
    literal: string;
    match: boolean;
    done: boolean;
}

export default function Case({ literal, match, done }: CaseProps) {
    function getColor(match: boolean) {
        let color = match ? "#2c4822" : "#552121";
        if (done) {
            color = match ? "#259c25" : "#bf1717";
        }
        return color;
    };

    return <div className="case">
        <p className="status" style={{ color: getColor(match) }}>{match ? "match" : "ignore"}</p>
        <p className="literal">{literal}</p>
        <span style={{ background: getColor(match) }}></span>
    </div>;
}
