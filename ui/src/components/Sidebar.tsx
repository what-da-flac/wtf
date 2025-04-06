import {Link} from "react-router-dom";

const Sidebar = () => {
    return (
        <div>
            <Link to="/">Home SB</Link>
            <Link to="/about">About SB</Link>
        </div>
    );
};

export default Sidebar;