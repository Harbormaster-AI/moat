import React, { Component } from 'react'
import PayrollRunService from '../services/PayrollRunService'

class ListPayrollRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                payrollRuns: []
        }
        this.addPayrollRun = this.addPayrollRun.bind(this);
        this.editPayrollRun = this.editPayrollRun.bind(this);
        this.deletePayrollRun = this.deletePayrollRun.bind(this);
    }

    deletePayrollRun(id){
        PayrollRunService.deletePayrollRun(id).then( res => {
            this.setState({payrollRuns: this.state.payrollRuns.filter(payrollRun => payrollRun.payrollRunId !== id)});
        });
    }
    viewPayrollRun(id){
        this.props.history.push(`/view-payrollRun/${id}`);
    }
    editPayrollRun(id){
        this.props.history.push(`/add-payrollRun/${id}`);
    }

    componentDidMount(){
        PayrollRunService.getPayrollRuns().then((res) => {
            this.setState({ payrollRuns: res.data});
        });
    }

    addPayrollRun(){
        this.props.history.push('/add-payrollRun/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PayrollRun List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPayrollRun}> Add PayrollRun</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RunNumber </th>
                                    <th> PeriodStart </th>
                                    <th> PeriodEnd </th>
                                    <th> PaymentDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.payrollRuns.map(
                                        payrollRun => 
                                        <tr key = {payrollRun.payrollRunId}>
                                             <td> { payrollRun.runNumber } </td>
                                             <td> { payrollRun.periodStart } </td>
                                             <td> { payrollRun.periodEnd } </td>
                                             <td> { payrollRun.paymentDate } </td>
                                             <td> { payrollRun.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPayrollRun(payrollRun.payrollRunId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePayrollRun(payrollRun.payrollRunId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPayrollRun(payrollRun.payrollRunId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPayrollRunComponent
