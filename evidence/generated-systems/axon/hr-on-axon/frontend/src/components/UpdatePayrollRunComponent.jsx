import React, { Component } from 'react'
import PayrollRunService from '../services/PayrollRunService';

class UpdatePayrollRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                runNumber: '',
                periodStart: '',
                periodEnd: '',
                paymentDate: '',
                status: ''
        }
        this.updatePayrollRun = this.updatePayrollRun.bind(this);

        this.changerunNumberHandler = this.changerunNumberHandler.bind(this);
        this.changeperiodStartHandler = this.changeperiodStartHandler.bind(this);
        this.changeperiodEndHandler = this.changeperiodEndHandler.bind(this);
        this.changepaymentDateHandler = this.changepaymentDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        PayrollRunService.getPayrollRunById(this.state.id).then( (res) =>{
            let payrollRun = res.data;
            this.setState({
                runNumber: payrollRun.runNumber,
                periodStart: payrollRun.periodStart,
                periodEnd: payrollRun.periodEnd,
                paymentDate: payrollRun.paymentDate,
                status: payrollRun.status
            });
        });
    }

    updatePayrollRun = (e) => {
        e.preventDefault();
        let payrollRun = {
            payrollRunId: this.state.id,
            runNumber: this.state.runNumber,
            periodStart: this.state.periodStart,
            periodEnd: this.state.periodEnd,
            paymentDate: this.state.paymentDate,
            status: this.state.status
        };
        console.log('payrollRun => ' + JSON.stringify(payrollRun));
        console.log('id => ' + JSON.stringify(this.state.id));
        PayrollRunService.updatePayrollRun(payrollRun).then( res => {
            this.props.history.push('/payrollRuns');
        });
    }

    changerunNumberHandler= (event) => {
        this.setState({runNumber: event.target.value});
    }
    changeperiodStartHandler= (event) => {
        this.setState({periodStart: event.target.value});
    }
    changeperiodEndHandler= (event) => {
        this.setState({periodEnd: event.target.value});
    }
    changepaymentDateHandler= (event) => {
        this.setState({paymentDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/payrollRuns');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PayrollRun</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> runNumber: </label>
                                                <input placeholder="runNumber" name="runNumber" className="form-control" value={this.state.runNumber} onChange={this.changerunNumberHandler}/>

                                            <label> periodStart: </label>
                                                <input type="date" placeholder="periodStart" name="periodStart" className="form-control" value={this.state.periodStart} onChange={this.changeperiodStartHandler}/>

                                            <label> periodEnd: </label>
                                                <input type="date" placeholder="periodEnd" name="periodEnd" className="form-control" value={this.state.periodEnd} onChange={this.changeperiodEndHandler}/>

                                            <label> paymentDate: </label>
                                                <input type="date" placeholder="paymentDate" name="paymentDate" className="form-control" value={this.state.paymentDate} onChange={this.changepaymentDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Scheduled
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Reversed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePayrollRun}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdatePayrollRunComponent
