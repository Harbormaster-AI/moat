import React, { Component } from 'react'
import EnterpriseService from '../services/EnterpriseService'

class ViewEnterpriseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            enterprise: {}
        }
    }

    componentDidMount(){
        EnterpriseService.getEnterpriseById(this.state.id).then( res => {
            this.setState({enterprise: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Enterprise Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.enterprise.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> legalName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.enterprise.legalName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> registrationCountry:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.enterprise.registrationCountry }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.enterprise.website }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.enterprise.taxId }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEnterpriseComponent
