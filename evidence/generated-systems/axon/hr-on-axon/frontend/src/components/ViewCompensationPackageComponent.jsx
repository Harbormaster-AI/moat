import React, { Component } from 'react'
import CompensationPackageService from '../services/CompensationPackageService'

class ViewCompensationPackageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            compensationPackage: {}
        }
    }

    componentDidMount(){
        CompensationPackageService.getCompensationPackageById(this.state.id).then( res => {
            this.setState({compensationPackage: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CompensationPackage Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveFrom:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.compensationPackage.effectiveFrom }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveTo:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.compensationPackage.effectiveTo }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> currency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.compensationPackage.currency }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCompensationPackageComponent
