import React, { Component } from 'react'
import ControlService from '../services/ControlService'

class ViewControlComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            control: {}
        }
    }

    componentDidMount(){
        ControlService.getControlById(this.state.id).then( res => {
            this.setState({control: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Control Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.control.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> objective:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.control.objective }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ownerDepartment:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.control.ownerDepartment }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ControlType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.control.controlType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Frequency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.control.frequency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.control.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewControlComponent
