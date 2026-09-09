import React, { Component } from 'react'
import OperationService from '../services/OperationService'

class ViewOperationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            operation: {}
        }
    }

    componentDidMount(){
        OperationService.getOperationById(this.state.id).then( res => {
            this.setState({operation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Operation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> operationNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.operation.operationNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.operation.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> setupTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.operation.setupTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> standardCycleTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.operation.standardCycleTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> OperationType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.operation.operationType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOperationComponent
