import React, { Component } from 'react'
import DataTaskService from '../services/DataTaskService'

class ViewDataTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dataTask: {}
        }
    }

    componentDidMount(){
        DataTaskService.getDataTaskById(this.state.id).then( res => {
            this.setState({dataTask: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DataTask Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataTask.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> command:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataTask.command }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> retries:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataTask.retries }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> TaskType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataTask.taskType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDataTaskComponent
