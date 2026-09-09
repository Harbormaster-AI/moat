import React, { Component } from 'react'
import NotebookService from '../services/NotebookService'

class ViewNotebookComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            notebook: {}
        }
    }

    componentDidMount(){
        NotebookService.getNotebookById(this.state.id).then( res => {
            this.setState({notebook: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Notebook Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.notebook.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> repository:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.notebook.repository }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Language:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.notebook.language }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewNotebookComponent
