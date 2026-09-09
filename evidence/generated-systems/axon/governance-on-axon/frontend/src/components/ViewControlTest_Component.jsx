import React, { Component } from 'react'
import ControlTest_Service from '../services/ControlTest_Service'

class ViewControlTest_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            controlTest_: {}
        }
    }

    componentDidMount(){
        ControlTest_Service.getControlTest_ById(this.state.id).then( res => {
            this.setState({controlTest_: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ControlTest_ Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.controlTest_.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> testPeriodStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.controlTest_.testPeriodStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> testPeriodEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.controlTest_.testPeriodEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> sampleSize:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.controlTest_.sampleSize }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> TestType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.controlTest_.testType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Effectiveness:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.controlTest_.effectiveness }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.controlTest_.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewControlTest_Component
