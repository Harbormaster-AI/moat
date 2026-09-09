import React, { Component } from 'react'
import RegulationService from '../services/RegulationService'

class ViewRegulationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            regulation: {}
        }
    }

    componentDidMount(){
        RegulationService.getRegulationById(this.state.id).then( res => {
            this.setState({regulation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Regulation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.regulation.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> citation:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.regulation.citation }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> jurisdiction:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.regulation.jurisdiction }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> publicationUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.regulation.publicationUrl }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRegulationComponent
