import React, { Component } from 'react'
import DischargeService from '../services/DischargeService'

class ViewDischargeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            discharge: {}
        }
    }

    componentDidMount(){
        DischargeService.getDischargeById(this.state.id).then( res => {
            this.setState({discharge: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Discharge Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dischargeDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.discharge.dischargeDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Disposition:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.discharge.disposition }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDischargeComponent
