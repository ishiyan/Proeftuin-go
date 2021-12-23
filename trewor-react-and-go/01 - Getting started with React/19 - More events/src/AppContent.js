import React, { Component } from 'react';

export default class AppContent extends Component {

    anotherFunction = () => {
        console.log("another function");
    }

    leftParagraph = () => {
        console.log("left the paragraph");
    }

    fetchList = () => {
        fetch('https://jsonplaceholder.typicode.com/posts')
            .then((response) => response.json())
            .then(json => {
                console.log(json);
                this.anotherFunction();
                let posts = document.getElementById("post-list");

                json.forEach(function(obj){
                    let li = document.createElement("li");
                    li.appendChild(document.createTextNode(obj.title));
                    posts.appendChild(li);
                })
            })
    }

    render(){
        return (
            <p>
                This is the content.

                <br />
                <hr />
                <p onMouseEnter={this.anotherFunction} onMouseLeave={this.leftParagraph}>This is some text</p>


                <button onClick={this.fetchList} className="btn btn-primary">Fetch Data</button>

                <hr />

                <ul id="post-list"></ul>
            </p>
        );
    }
}