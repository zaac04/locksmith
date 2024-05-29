import React, { useEffect, useState } from "react";
import { SiTicktick } from "react-icons/si";
import { DiffEditor, Editor } from "@monaco-editor/react";
import {
  IsStagePresent,
  ReadCipherFile,
  WriteCipherFile,
} from "../../../wailsjs/go/ui/App";

interface ModalProps {
  OnClose: () => void;
  refresh: () => void;
  stage: string;
  Location:string
}

const Modal = (props: ModalProps) => {
  const [text, setText] = useState("");
  const [changedValue, setChangedValue] = useState("");
  const [showEncryptionBox, setShowEncryptionBox] = useState(false);
  const [key, setKey] = useState("");
  const [createKey, setCreateKey] = useState(true);
  const [showCreateToggle, setShowCreateToggle] = useState(false);
  const [currentStage, setCurrentStage] = useState(0);
  const [isStageAvailable, setIsStageAvailable] = useState(false);
  const [newStage, setNewStage] = useState<string>("");

  const handleModalNextStage = () => setCurrentStage(currentStage + 1);

  useEffect(() => {
    if (props.stage !== "") {
      setNewStage(props.stage);
      IsStagePresent(props.stage).then((valid) => {
        setShowEncryptionBox(valid);
        setIsStageAvailable(!valid);
        setShowCreateToggle(!valid);
      });
    } else {
      resetState()
    }
  }, [props.stage]);

  const handleModalFetchEnv = () => {
    if (!key && createKey) {
      setText("");
      setChangedValue("");
    } else if (key) {
      ReadCipherFile(newStage, key).then((value) => {
        setText(value);
        setChangedValue(value);
      });
    }
    handleModalNextStage();
  };

  const handleWrite = () => {
    const temp = props.stage!="" && props.Location!="" ? props.Location : newStage
    console.log(temp)


    WriteCipherFile(temp, key, changedValue).then(() => {
      console.log("Success");
      handleCloseModal();
      props.OnClose();
      props.refresh();
    });
  };

  const handleStageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const stageName = event.target.value;
    setNewStage(stageName);

    if (!stageName) {
      resetState();
      return;
    }

    IsStagePresent(stageName).then((valid) => {
      setShowEncryptionBox(valid);
      setIsStageAvailable(!valid);
      setShowCreateToggle(!valid);
    });
  };

  const handleCreateKeyToggle = () => {
    setShowEncryptionBox(createKey);
    setCreateKey(!createKey);
  };

  const handleCloseModal = () => {
    setCurrentStage(0);
    resetState();
    props.OnClose();
  };

  const resetState = () => {
    setShowEncryptionBox(false);
    setNewStage("");
    setKey("");
    setIsStageAvailable(false);
    setShowCreateToggle(false);
  };

  const stages = [
    // Stage 0
    <div className="w-full h-full" key={0}>
      <h2 className="text-2xl font-bold mb-4">Enter Stage</h2>
      <div className="flex items-center">
        <label className="w-2/12 text-wrap pt-1">Stage Name :</label>
        <input
          type="text"
          className="p-2 m-1"
          placeholder="Stage name"
          value={newStage}
          autoFocus
          onChange={handleStageChange}
        />
        <span
          className={`p-1 m-1 flex ${isStageAvailable ? "visible" : "hidden"}`}
        >
          Available{" "}
          <SiTicktick
            className={`ml-2 text-xl pl-2 text-green-500 ${
              isStageAvailable ? "visible" : "hidden"
            }`}
          />
        </span>
      </div>
      <div
        className={showCreateToggle ? "flex items-center visible" : "hidden"}
        onClick={handleCreateKeyToggle}
      >
        <label className="w-2/12 text-wrap">Create Key:</label>
        <div className="hover:bg-gray-800">
          <span className={createKey ? "font-bold underline" : ""}>Yes</span>
          <span>/</span>
          <span className={!createKey ? "font-bold underline" : ""}>No</span>
        </div>
      </div>
      <div
        className={showEncryptionBox ? "flex items-center visible" : "hidden"}
      >
        <label className="w-2/12 text-wrap">Encryption Key:</label>
        <input
          type="password"
          className="p-2 m-1"
          value={key}
          autoFocus
          onChange={(e) => setKey(e.target.value)}
        />
      </div>
      <button
        className="mt-4 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
        onClick={handleModalFetchEnv}
      >
        Next
      </button>
    </div>,

    // Stage 1
    <div className="w-full h-full" key={1}>
      <h2 className="text-2xl font-bold mb-4">Enter Env's</h2>
      <div className="h-3/4">
        <Editor
          defaultLanguage="plaintext"
          value={changedValue}
          height="100%"
          defaultValue="// enter env as key=value pairs"
          theme="vs-dark"
          className="overflow-auto pt-1"
          onChange={(value) => setChangedValue(value || "")}
        />
      </div>
      <button
        className="mt-4 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
        onClick={() => setCurrentStage(currentStage - 1)}
      >
        Previous
      </button>
      <button
        className="ml-4 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
        onClick={handleModalNextStage}
      >
        Next
      </button>
    </div>,

    // Stage 2
    <div className="w-full h-full" key={2}>
      <h2 className="text-2xl font-bold mb-4">Review</h2>
      <div className="h-3/4">
        <DiffEditor
          theme="vs-dark"
          height="100%"
          original={text}
          modified={changedValue}
          options={{ readOnly: true }}
        />
      </div>
      <button
        className="mt-4 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
        onClick={() => setCurrentStage(currentStage - 1)}
      >
        Previous
      </button>
      <button
        className="ml-4 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
        onClick={handleWrite}
      >
        Write
      </button>
    </div>,
  ];

  return stages[currentStage];
};

export default Modal;
