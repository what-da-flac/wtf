import { AiOutlinePlus } from "react-icons/ai"
import { TbPlaylist } from "react-icons/tb"
import {Link} from "react-router-dom";


const Library = () => {
    return (
        <div className="flex flex-col">
            <div className="flex items-center justify-between px-5 pt-4">
                <div className="inline-flex items-center gap-x-2">
                    <TbPlaylist size={26} className="text-neutral-400" />
                    <p className="text-neutral-400 font-medium">
                        Your Library
                    </p>
                </div>
                <Link to={`/library/new`}>
                    <AiOutlinePlus size={20} className="text-neutral-400 cursor-pointer hover:text-white transition"/>
                </Link>
            </div>
            <div className="flex flex-col mt-4 px-3">
                This is Library component
            </div>
        </div>
    )
}

export default Library