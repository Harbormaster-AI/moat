import React, { Component } from 'react'
import DistributorService from '../services/DistributorService'

class ViewDistributorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            distributor: {}
        }
    }

    componentDidMount(){
        DistributorService.getDistributorById(this.state.id).then( res => {
            this.setState({distributor: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Distributor Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.distributor.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> licenseNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.distributor.licenseNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> region:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.distributor.region }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DistributorType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.distributor.distributorType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDistributorComponent
