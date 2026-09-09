import React, { Component } from 'react'
import BusinessUnitService from '../services/BusinessUnitService'

class ViewBusinessUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            businessUnit: {}
        }
    }

    componentDidMount(){
        BusinessUnitService.getBusinessUnitById(this.state.id).then( res => {
            this.setState({businessUnit: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BusinessUnit Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.businessUnit.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> leader:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.businessUnit.leader }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBusinessUnitComponent
