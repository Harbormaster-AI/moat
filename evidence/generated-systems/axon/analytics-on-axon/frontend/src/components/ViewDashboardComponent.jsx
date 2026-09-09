import React, { Component } from 'react'
import DashboardService from '../services/DashboardService'

class ViewDashboardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dashboard: {}
        }
    }

    componentDidMount(){
        DashboardService.getDashboardById(this.state.id).then( res => {
            this.setState({dashboard: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Dashboard Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dashboard.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> theme:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dashboard.theme }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dashboard.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDashboardComponent
