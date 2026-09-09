import React, { Component } from 'react'
import ThirdPartyService from '../services/ThirdPartyService'

class ViewThirdPartyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            thirdParty: {}
        }
    }

    componentDidMount(){
        ThirdPartyService.getThirdPartyById(this.state.id).then( res => {
            this.setState({thirdParty: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ThirdParty Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdParty.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> country:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdParty.country }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> contactEmail:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdParty.contactEmail }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ThirdPartyType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdParty.thirdPartyType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Criticality:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.thirdParty.criticality }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewThirdPartyComponent
