import { MdArrowForwardIos } from "react-icons/md";

export default function Submit({ done }: any) {
    let classList = "submit";

    if (done) {
        classList = classList.concat(" submit-done");
    }

    return (
        <button className={classList} disabled={!done}>
            <MdArrowForwardIos size={20} color="white" />
        </button>
    )
}
