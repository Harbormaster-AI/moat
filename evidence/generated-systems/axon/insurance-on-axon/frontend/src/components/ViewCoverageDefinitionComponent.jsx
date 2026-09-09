import React, { Component } from 'react'
import CoverageDefinitionService from '../services/CoverageDefinitionService'

class ViewCoverageDefinitionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            coverageDefinition: {}
        }
    }

    componentDidMount(){
        CoverageDefinitionService.getCoverageDefinitionById(this.state.id).then( res => {
            this.setState({coverageDefinition: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CoverageDefinition Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverageDefinition.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> defaultLimit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverageDefinition.defaultLimit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> defaultDeductible:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverageDefinition.defaultDeductible }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asMandatory:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverageDefinition.asMandatory }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> CoverageType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverageDefinition.coverageType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCoverageDefinitionComponent
