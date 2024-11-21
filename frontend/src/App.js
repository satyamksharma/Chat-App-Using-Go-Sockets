import "./App.css";
import Header from "./components/header";
import ChatInput from "./components/chatInput";
import React, { Component } from "react";
import { connect, sendMsg } from "./api";
import ChatHistory from "./components/chatHistory";

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            chatHistory: [],
        };

        this.send = this.send.bind(this);
    }

    componentDidMount() {
        connect((msg) => {
            console.log("New Messages from user");

            this.setState((prevState) => ({
                chatHistory: [...prevState.chatHistory, msg],
            }));
            console.log(this.state);
        });
    }

    send(event) {
        if (event.keyCode === 13) {
            sendMsg(event.target.value);
            event.target.value = "";
        }
    }

    render() {
        return (
            <div className="App">
                <Header />
                <ChatHistory chatHistory={this.state.chatHistory} />
                <ChatInput send={this.send} />
            </div>
        );
    }
}

export default App;
