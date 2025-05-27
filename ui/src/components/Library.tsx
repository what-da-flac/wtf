import { AiOutlinePlus } from "react-icons/ai"
import { TbPlaylist } from "react-icons/tb"
import {Link} from "react-router-dom";


const Library = () => {
    return (
        <div className="box-red">
            <div>
                <div className="box-red">
                    <TbPlaylist size={26} />
                    <p>
                        Your Library
                    </p>
                </div>
                <Link to={`/library/new`}>
                    <AiOutlinePlus size={20}/>
                </Link>
            </div>
            <div>
                This is Library component
            </div>
        </div>
    )
}

export default Library