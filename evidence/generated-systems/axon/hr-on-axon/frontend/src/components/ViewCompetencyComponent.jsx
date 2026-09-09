import React, { Component } from 'react'
import CompetencyService from '../services/CompetencyService'

class ViewCompetencyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            competency: {}
        }
    }

    componentDidMount(){
        CompetencyService.getCompetencyById(this.state.id).then( res => {
            this.setState({competency: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Competency Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.competency.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> category:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.competency.category }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCompetencyComponent
