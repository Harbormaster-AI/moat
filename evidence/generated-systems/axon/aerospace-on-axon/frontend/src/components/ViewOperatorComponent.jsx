import React, { Component } from 'react'
import OperatorService from '../services/OperatorService'

class ViewOperatorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            operator: {}
        }
    }

    componentDidMount(){
        OperatorService.getOperatorById(this.state.id).then( res => {
            this.setState({operator: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Operator Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.operator.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> icaoDesignator:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.operator.icaoDesignator }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> OperatorType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.operator.operatorType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOperatorComponent
