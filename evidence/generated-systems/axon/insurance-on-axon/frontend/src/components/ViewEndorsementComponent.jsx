import React, { Component } from 'react'
import EndorsementService from '../services/EndorsementService'

class ViewEndorsementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            endorsement: {}
        }
    }

    componentDidMount(){
        EndorsementService.getEndorsementById(this.state.id).then( res => {
            this.setState({endorsement: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Endorsement Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endorsementNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.endorsement.endorsementNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.endorsement.effectiveDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.endorsement.description }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEndorsementComponent
