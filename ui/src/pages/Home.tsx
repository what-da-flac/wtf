import {Link} from "react-router-dom";

export default function Home() {
    return (
        <>
            <p>I am home page</p>
            <Link to="/about">About</Link>
        </>
    )
}