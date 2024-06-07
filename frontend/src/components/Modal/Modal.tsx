import React, { useEffect, useRef, useState } from "react";
import { SiTicktick } from "react-icons/si";
import { VscLoading } from "react-icons/vsc";
import { DiffEditor, Editor, loader } from "@monaco-editor/react";
import * as monaco from 'monaco-editor';
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker';
import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker';
import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker';
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker';

self.MonacoEnvironment = {
  getWorker(_, label) {
    if (label === 'json') {
      return new jsonWorker();
    }
    if (label === 'css' || label === 'scss' || label === 'less') {
      return new cssWorker();
    }
    if (label === 'html' || label === 'handlebars' || label === 'razor') {
      return new htmlWorker();
    }
    if (label === 'typescript' || label === 'javascript') {
      return new tsWorker();
    }
    return new editorWorker();
  },
};

loader.config({ monaco });
import {
  IsStagePresent,
  ReadCipherFile,
  WriteCipherFile,
} from "../../../wailsjs/go/ui/App";
import { debounce } from "lodash";
import { Bounce, toast } from 'react-toastify';
import "react-toastify/dist/ReactToastify.css";
interface ModalProps {
  OnClose: () => void;
  refresh: () => void;
  stage: string;
  Location: string;
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
  const [IsLoading, setIsLoading] = useState(false);

  const handleModalNextStage = () => setCurrentStage(currentStage + 1);

  useEffect(() => {
    if (props.stage !== "") {
      setNewStage(props.stage);
      IsStagePresent(props.stage)
        .then((valid) => {
          setShowEncryptionBox(valid);
          setIsStageAvailable(!valid);
          setShowCreateToggle(!valid);
        })
        .catch((err) => {
          toast(err)
          throw err;
        });
    } else {
      resetState();
    }
  }, [props.stage]);

  const handleModalFetchEnv = () => {
    if (!key && createKey) {
      setText("");
      setChangedValue("");
      handleModalNextStage();
    } else if (key) {
      ReadCipherFile(newStage, key)
        .then((value) => {
          setText(value);
          setChangedValue(value);
          handleModalNextStage();
        })
        .catch(() => {
          toast.error("invalid key entered", {
            position: "top-right",
            autoClose: 3000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "dark",
            transition: Bounce,
            });
        });
    }

  };

  const handleWrite = () => {
    const temp =
      props.stage != "" && props.Location != "" ? props.Location : newStage;

    WriteCipherFile(temp, key, changedValue)
      .then(() => {
        console.log("Success");
        handleCloseModal();
        props.OnClose();
        props.refresh();
      })
      .catch((err) => {
        toast.error(err, {
          position: "top-right",
          autoClose: 3000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "dark",
          transition: Bounce,
          });
      });
  };

  const handleStageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const stageName = event.target.value;
    setIsLoading(true);
    setIsStageAvailable(false);
    setNewStage(stageName);
    debounceSearch(stageName);
  };

  const debounceSearch = useRef(
    debounce(async (stageName: string) => {
      if (stageName) {
        console.log(stageName);
        IsStagePresent(stageName).then((valid) => {
          setShowEncryptionBox(valid);
          setIsStageAvailable(!valid);
          setShowCreateToggle(!valid);
        });
      } else {
        resetState();
      }
      setIsLoading(false);
    }, 600)
  ).current;

  React.useEffect(() => {
    return () => {
      debounceSearch.cancel();
    };
  }, [debounceSearch]);

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
          className="p-2 m-1 bg-gray-950 text-gray-300"
          placeholder="Stage name"
          value={newStage}
          autoFocus
          onChange={handleStageChange}
        />
        <VscLoading
          className={`m-1 animate-spin ${IsLoading ? "visible" : "hidden"} `}
        ></VscLoading>
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
          className="p-2 m-1  bg-gray-950 text-gray-300 border-emerald-50"
          value={key}
          autoFocus
          onChange={(e) => setKey(e.target.value)}
        />
      </div>
      <button
        className={`mt-4 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded ${
          isStageAvailable || showEncryptionBox ? "visible" : "hidden"
        }`}
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

    <div className="w-full h-full" key={3}>
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
