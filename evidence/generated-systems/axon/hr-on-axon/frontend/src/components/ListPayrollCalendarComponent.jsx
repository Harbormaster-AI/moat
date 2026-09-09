import React, { Component } from 'react'
import PayrollCalendarService from '../services/PayrollCalendarService'

class ListPayrollCalendarComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                payrollCalendars: []
        }
        this.addPayrollCalendar = this.addPayrollCalendar.bind(this);
        this.editPayrollCalendar = this.editPayrollCalendar.bind(this);
        this.deletePayrollCalendar = this.deletePayrollCalendar.bind(this);
    }

    deletePayrollCalendar(id){
        PayrollCalendarService.deletePayrollCalendar(id).then( res => {
            this.setState({payrollCalendars: this.state.payrollCalendars.filter(payrollCalendar => payrollCalendar.payrollCalendarId !== id)});
        });
    }
    viewPayrollCalendar(id){
        this.props.history.push(`/view-payrollCalendar/${id}`);
    }
    editPayrollCalendar(id){
        this.props.history.push(`/add-payrollCalendar/${id}`);
    }

    componentDidMount(){
        PayrollCalendarService.getPayrollCalendars().then((res) => {
            this.setState({ payrollCalendars: res.data});
        });
    }

    addPayrollCalendar(){
        this.props.history.push('/add-payrollCalendar/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PayrollCalendar List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPayrollCalendar}> Add PayrollCalendar</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Country </th>
                                    <th> PayFrequency </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.payrollCalendars.map(
                                        payrollCalendar => 
                                        <tr key = {payrollCalendar.payrollCalendarId}>
                                             <td> { payrollCalendar.name } </td>
                                             <td> { payrollCalendar.country } </td>
                                             <td> { payrollCalendar.payFrequency } </td>
                                             <td>
                                                 <button onClick={ () => this.editPayrollCalendar(payrollCalendar.payrollCalendarId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePayrollCalendar(payrollCalendar.payrollCalendarId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPayrollCalendar(payrollCalendar.payrollCalendarId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListPayrollCalendarComponent
