import React from 'react';
import {Greet} from "../wailsjs/go/main/App";
import {EventsOn} from "../wailsjs/runtime/runtime";


class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            resultText: null,
            inputText: null,
            finishTask:0,
            allTask:0,
            process:0,
            rate:"",
            size:"",
        };
    }

    greet() {
        Greet(this.state.inputText).then((result) => {
            this.setState({
                resultText: result
            })
        });
    }

    handleChange(e) {
        this.setState({
            inputText: e.target.value
        })
    }
    handleFileChange(e) {
        console.log(e)
    }

    componentDidMount() {
        var _this = this
        EventsOn("MESSAGE_TEST", function (m) {
            console.log(m.split("\t"))
            let data = m.split("\t")
            _this.setState({
                finishTask:data[0],
                allTask:data[1],
                process:data[2],
                rate:data[3],
                size:data[4],
            })
        });
    }
    render() {
        return (
            <div className="app">
                <div className="header">
                    <div className="header-container">
                        <a className="header-item" href="#" title="SVGDB Logo">
                        </a>
                        <div className="search-bar search-active">
                            <button type="submit" className="btn-icon">
                                <svg t="1655347290236" className="icon" viewBox="0 0 1024 1024" version="1.1"
                                    xmlns="http://www.w3.org/2000/svg" p-id="2958" width="22" height="22">
                                    <path d="M768 448a362.666667 362.666667 0 1 0-725.333333 0 362.666667 362.666667 0 0 0 725.333333 0z m-640 0a277.333333 277.333333 0 1 1 554.666667 0 277.333333 277.333333 0 0 1-554.666667 0z m739.925333 525.568l-194.304-196.949333 60.757334-59.904 194.304 196.949333-60.757334 59.904z"
                                        p-id="2959" fill="#777"></path>
                                </svg>
                            </button>
                            <input type="text" onChange={(e) => this.handleChange(e)} autoComplete="off" placeholder="请输入下载文件地址" required />
                        </div>
                        <a className="header-item saved" href="javascript:;"  onClick={() => this.greet()}>
                            <svg t="1661912744398" className="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2959" width="22" height="22"><path d="M505.7 661c3.2 4.1 9.4 4.1 12.6 0l112-141.7c4.1-5.2 0.4-12.9-6.3-12.9h-74.1V168c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v338.3H400c-6.7 0-10.4 7.7-6.3 12.9l112 141.8z" p-id="2960" fill="#8899a4"></path><path d="M878 626h-60c-4.4 0-8 3.6-8 8v154H214V634c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v198c0 17.7 14.3 32 32 32h684c17.7 0 32-14.3 32-32V634c0-4.4-3.6-8-8-8z" p-id="2961" fill="#8899a4"></path></svg>
                            <span>下载</span>
                        </a>
                    </div>
                </div>
                <div className="container">
                    <p>{this.state.resultText}</p>
                    <div className="rounded rounded-white">
                        <h1>进度 【<span>{this.state.finishTask}</span>/{this.state.allTask}】{this.state.process}% {this.state.rate} {this.state.size}</h1>
                    </div>
                </div>
            </div>
        )
    }
}

export default App
