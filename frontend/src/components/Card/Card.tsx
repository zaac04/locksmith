import { FaEdit } from "react-icons/fa";
interface CardProps {
  stage: string;
  LastModified: string;
  Algorithm: string;
  FileName: string;
  CipherBytes: number;
  OnView: () => void;
  // Edit : (a:string,b:string)=>void
}

const Card = (Card: CardProps) => {
  return (
    <div className=" w-72 h-64 border  border-gray-300 m-2 p-3 rounded-lg shadow-md bg-gray-800 hover:bg-gray-0 transition duration-300">
      <h2 className="text-xl font-semibold mb-2">{Card.stage}</h2>
      <div className="text-gray-600">
        <ul className=" list-none min-h-40 ">
          <li className=" text-base  text-gray-500  hover:text-slate-200 deo text-wrap">
            LastModified: {Card.LastModified}
          </li>
          <li className="text-base text-gray-500 hover:text-slate-200 deo text-wrap">
            Algorithm: {Card.Algorithm}
          </li>
          <li className="text-base text-gray-500 hover:text-slate-200 deo text-wrap">
            FileName: {Card.FileName}
          </li>
          <li className="text-base text-gray-500 hover:text-slate-200 deo text-wrap">
            CipherBytes: {Card.CipherBytes} bytes
          </li>
        </ul>
      </div>
      <div className="flex w-full h-full border-t-2 p-3 justify-center  mb-1">
        <FaEdit
          className=" text-lg m-1 text-gray-500 hover:text-slate-200"
          onClick={Card.OnView}
        ></FaEdit>
      </div>
    </div>
  );
};

export default Card;
