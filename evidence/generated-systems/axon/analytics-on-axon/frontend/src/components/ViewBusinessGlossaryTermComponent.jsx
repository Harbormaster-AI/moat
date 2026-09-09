import React, { Component } from 'react'
import BusinessGlossaryTermService from '../services/BusinessGlossaryTermService'

class ViewBusinessGlossaryTermComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            businessGlossaryTerm: {}
        }
    }

    componentDidMount(){
        BusinessGlossaryTermService.getBusinessGlossaryTermById(this.state.id).then( res => {
            this.setState({businessGlossaryTerm: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BusinessGlossaryTerm Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> term:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.businessGlossaryTerm.term }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> definition:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.businessGlossaryTerm.definition }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> steward:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.businessGlossaryTerm.steward }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBusinessGlossaryTermComponent
