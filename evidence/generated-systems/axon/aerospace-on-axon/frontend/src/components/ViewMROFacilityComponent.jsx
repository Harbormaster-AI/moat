import React, { Component } from 'react'
import MROFacilityService from '../services/MROFacilityService'

class ViewMROFacilityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            mROFacility: {}
        }
    }

    componentDidMount(){
        MROFacilityService.getMROFacilityById(this.state.id).then( res => {
            this.setState({mROFacility: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MROFacility Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mROFacility.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> approvalScope:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mROFacility.approvalScope }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mROFacility.address }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMROFacilityComponent
