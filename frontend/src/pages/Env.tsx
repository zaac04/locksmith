import { GetEnvs } from "../../wailsjs/go/ui/App";
import { IoMdAddCircle } from "react-icons/io";
import { FaSyncAlt } from "react-icons/fa";
import { useEffect, useState, useRef } from "react";
import Card from "../components/Card/Card";
import Modal from "../components/Modal/Modal";
import { Bounce, ToastContainer } from "react-toastify";

interface StageInfo {
  Algorithm: string;
  CipherSize: number;
  FileName: string;
  LastModified: string;
  StageName: string;
  Id: string;
}

function Env() {
  const modalEl = useRef<HTMLDivElement>(null);
  const [isOpen, setIsOpen] = useState(false);
  const [stages, setStages] = useState<StageInfo[]>();
  const [stage, setStage] = useState("");
  const [location, setLocation] = useState("");
  const getStages = () => {
    GetEnvs()
      .then((value) => setStages(value))
      .catch((err) => {
        throw err;
      });
  };
  ``;
  const handleOpenModal = () => {
    setStage("");
    setLocation("");
    setIsOpen(!isOpen);
  };

  const handleClose = (e: React.MouseEvent) => {
    if (modalEl.current && !modalEl.current.contains(e.target as Node)) {
      setIsOpen(!isOpen);
      setStage("");
      setLocation("");
    }
  };

  const handleEditFile = (stage: string, location: string) => {
    setStage(stage);
    setLocation(location);
    setIsOpen((prev) => !prev);
  };

  const refresh = () => getStages();

  useEffect(() => {
    getStages();
  }, []);

  return (
    <>
      <div className="p-3 m-6 h-screen" onClick={handleClose}>
        <ToastContainer
          position="top-right"
          autoClose={3000}
          hideProgressBar={false}
          newestOnTop={false}
          closeOnClick
          rtl={false}
          pauseOnFocusLoss
          draggable
          pauseOnHover
          theme="light"
          transition={Bounce}
        />
        {/* Same as */}
        <ToastContainer />
        <div className="flex justify-between">
          <p className=" text-xl font-semibold p-1 m-1">Stages</p>

          <div className="flex justify-between items-center bg-slate-900 p-1">
            <div
              className="flex bg-slate-800 p-2 hover:bg-slate-900"
              onClick={refresh}
            >
              <FaSyncAlt className="m-1" />
              Sync
            </div>
            <button
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
              onClick={handleOpenModal}
            >
              <div className="flex justify-between">
                <IoMdAddCircle className="m-1" />
                <p>Add Stage</p>
              </div>
            </button>
          </div>
        </div>

        <div className="h-full w-full overflow-auto">
          {stages && stages.length > 0 ? (
            <div className="flex flex-wrap justify-start">
              {stages?.map((stage) => (
                <Card
                  stage={stage.StageName}
                  LastModified={stage.LastModified}
                  Algorithm={stage.Algorithm}
                  FileName={stage.FileName}
                  CipherBytes={stage.CipherSize}
                  OnView={() => handleEditFile(stage.StageName, stage.Id)}
                />
              ))}
            </div>
          ) : (
            <div className="flex justify-center items-center align-middle h-full">
              No Stages Found
            </div>
          )}
        </div>

        {isOpen && (
          <div className="fixed inset-0 flex items-center justify-center">
            <div className="fixed inset-0 bg-gray-700 bg-opacity-95"></div>
            <div
              className="bg-slate-800 w-3/4 h-2/3 p-8 rounded-lg z-10 overflow-auto"
              ref={modalEl}
            >
              <Modal
                OnClose={() => setIsOpen(!isOpen)}
                refresh={() => refresh()}
                stage={stage}
                Location={location}
              />
            </div>
          </div>
        )}
      </div>
    </>
  );
}

export default Env;
