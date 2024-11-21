import React, { Component } from "react";

class Message extends Component {
    constructor(props) {
        super(props);
        let temp = JSON.parse(this.props.message);
        this.state = {
            message: temp,
        };
    }

    render() {
        return (
            <div className="w-[90%] block m-auto bg-sky-300 p-2 rounded-lg shadow-md shadow-black/25">
                <div className="flex justify-end">
                    {this.state.message.body}
                </div>
            </div>
        );
    }
}

export default Message;
