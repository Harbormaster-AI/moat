import React, { Component } from 'react'
import PharmacyService from '../services/PharmacyService'

class ViewPharmacyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            pharmacy: {}
        }
    }

    componentDidMount(){
        PharmacyService.getPharmacyById(this.state.id).then( res => {
            this.setState({pharmacy: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Pharmacy Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.pharmacy.name }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPharmacyComponent
