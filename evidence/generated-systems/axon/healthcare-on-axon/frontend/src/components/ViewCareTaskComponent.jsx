import React, { Component } from 'react'
import CareTaskService from '../services/CareTaskService'

class ViewCareTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            careTask: {}
        }
    }

    componentDidMount(){
        CareTaskService.getCareTaskById(this.state.id).then( res => {
            this.setState({careTask: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CareTask Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.careTask.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dueDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.careTask.dueDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.careTask.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.careTask.priority }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCareTaskComponent
