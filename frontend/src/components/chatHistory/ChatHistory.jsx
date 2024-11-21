import React, { Component } from "react";
import Message from "../message/Message";

class ChatHistory extends Component {
    render() {
        const messages = this.props.chatHistory.map((msg) => (
            <Message key={msg.timeStamp} message={msg.data} />
        ));
        console.log(messages);
        return (
            <div className="flex flex-col justify-center text-center items-center bg-lime-400">
                <h2 className=" w-full m-auto p-8 text-xl">Chat History</h2>
                {messages}
            </div>
        );
    }
}

export default ChatHistory;
