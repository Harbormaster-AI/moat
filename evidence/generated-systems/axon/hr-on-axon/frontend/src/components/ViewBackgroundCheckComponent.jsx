import React, { Component } from 'react'
import BackgroundCheckService from '../services/BackgroundCheckService'

class ViewBackgroundCheckComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            backgroundCheck: {}
        }
    }

    componentDidMount(){
        BackgroundCheckService.getBackgroundCheckById(this.state.id).then( res => {
            this.setState({backgroundCheck: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BackgroundCheck Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> checkNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.backgroundCheck.checkNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> provider:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.backgroundCheck.provider }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> completedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.backgroundCheck.completedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.backgroundCheck.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBackgroundCheckComponent
