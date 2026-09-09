import React, { Component } from 'react'
import CareTeamService from '../services/CareTeamService'

class ViewCareTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            careTeam: {}
        }
    }

    componentDidMount(){
        CareTeamService.getCareTeamById(this.state.id).then( res => {
            this.setState({careTeam: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CareTeam Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.careTeam.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> CareSetting:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.careTeam.careSetting }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCareTeamComponent
