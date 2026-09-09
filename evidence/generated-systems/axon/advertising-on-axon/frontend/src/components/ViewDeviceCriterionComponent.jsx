import React, { Component } from 'react'
import DeviceCriterionService from '../services/DeviceCriterionService'

class ViewDeviceCriterionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            deviceCriterion: {}
        }
    }

    componentDidMount(){
        DeviceCriterionService.getDeviceCriterionById(this.state.id).then( res => {
            this.setState({deviceCriterion: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DeviceCriterion Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DeviceType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.deviceCriterion.deviceType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PlatformType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.deviceCriterion.platformType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Operator:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.deviceCriterion.operator }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDeviceCriterionComponent
