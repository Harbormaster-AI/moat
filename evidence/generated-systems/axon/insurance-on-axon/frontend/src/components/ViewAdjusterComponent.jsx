import React, { Component } from 'react'
import AdjusterService from '../services/AdjusterService'

class ViewAdjusterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            adjuster: {}
        }
    }

    componentDidMount(){
        AdjusterService.getAdjusterById(this.state.id).then( res => {
            this.setState({adjuster: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Adjuster Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> firstName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adjuster.firstName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adjuster.lastName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> licenseNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adjuster.licenseNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AdjusterType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adjuster.adjusterType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAdjusterComponent
