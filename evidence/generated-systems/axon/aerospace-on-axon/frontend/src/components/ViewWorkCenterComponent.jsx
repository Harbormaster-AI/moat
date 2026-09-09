import React, { Component } from 'react'
import WorkCenterService from '../services/WorkCenterService'

class ViewWorkCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            workCenter: {}
        }
    }

    componentDidMount(){
        WorkCenterService.getWorkCenterById(this.state.id).then( res => {
            this.setState({workCenter: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View WorkCenter Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workCenter.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> capability:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workCenter.capability }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewWorkCenterComponent
