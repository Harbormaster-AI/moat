import React, { Component } from 'react'
import ProductionCertificateService from '../services/ProductionCertificateService'

class ViewProductionCertificateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            productionCertificate: {}
        }
    }

    componentDidMount(){
        ProductionCertificateService.getProductionCertificateById(this.state.id).then( res => {
            this.setState({productionCertificate: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ProductionCertificate Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> certificateNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionCertificate.certificateNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> authority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionCertificate.authority }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewProductionCertificateComponent
