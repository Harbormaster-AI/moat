import React, { Component } from 'react'
import TeamService from '../services/TeamService'

class ViewTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            team: {}
        }
    }

    componentDidMount(){
        TeamService.getTeamById(this.state.id).then( res => {
            this.setState({team: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Team Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.team.name }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTeamComponent
