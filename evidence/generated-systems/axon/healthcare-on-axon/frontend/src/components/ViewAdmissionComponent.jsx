import React, { Component } from 'react'
import AdmissionService from '../services/AdmissionService'

class ViewAdmissionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            admission: {}
        }
    }

    componentDidMount(){
        AdmissionService.getAdmissionById(this.state.id).then( res => {
            this.setState({admission: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Admission Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> admitDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.admission.admitDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bed:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.admission.bed }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AdmissionType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.admission.admissionType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAdmissionComponent
