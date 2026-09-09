import React, { Component } from 'react'
import PayrollItemService from '../services/PayrollItemService'

class ViewPayrollItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            payrollItem: {}
        }
    }

    componentDidMount(){
        PayrollItemService.getPayrollItemById(this.state.id).then( res => {
            this.setState({payrollItem: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PayrollItem Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollItem.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxable:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollItem.taxable }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ItemType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollItem.itemType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPayrollItemComponent
