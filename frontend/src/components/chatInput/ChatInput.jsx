import React, { Component } from "react";

class ChatInput extends Component {
    render() {
        return (
            <div className="w-[90%] block m-auto">
                <input
                    type="text"
                    placeholder="Type a message..."
                    onKeyDown={this.props.send}
                    className="w-full p-4 border-t-2 border-b-2 border-l-2 rounded-lg border-gray-50"
                />
            </div>
        );
    }
}

export default ChatInput;
