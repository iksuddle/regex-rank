import { useRouteError } from "react-router-dom";

export default function ErrorPage() {
    const error: any = useRouteError();
    console.error(error);

    return (
        <div style={{ display: "flex", flexDirection: "column", alignItems: "center", marginTop: "8rem" }}>
            <h1>Uh-oh..</h1>
            <p style={{ marginTop: "1rem" }}>smth went wrong mb</p>
            <p style={{ marginTop: "1rem" }}>
                <i>{error.statusText || error.message}</i>
            </p>
        </div>
    )
}

