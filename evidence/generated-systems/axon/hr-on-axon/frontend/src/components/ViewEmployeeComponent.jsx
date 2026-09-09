import React, { Component } from 'react'
import EmployeeService from '../services/EmployeeService'

class ViewEmployeeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            employee: {}
        }
    }

    componentDidMount(){
        EmployeeService.getEmployeeById(this.state.id).then( res => {
            this.setState({employee: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Employee Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> employeeNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employee.employeeNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employee.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> workEmail:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employee.workEmail }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> workPhone:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employee.workPhone }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dateOfHire:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employee.dateOfHire }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> nationalId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employee.nationalId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employee.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEmployeeComponent
