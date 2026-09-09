import React, { Component } from 'react'
import AllergyService from '../services/AllergyService'

class ViewAllergyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            allergy: {}
        }
    }

    componentDidMount(){
        AllergyService.getAllergyById(this.state.id).then( res => {
            this.setState({allergy: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Allergy Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> substance:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.allergy.substance }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reaction:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.allergy.reaction }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Severity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.allergy.severity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.allergy.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAllergyComponent
