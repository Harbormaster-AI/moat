import React, { Component } from 'react'
import PayrollRunService from '../services/PayrollRunService';

class CreatePayrollRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                runNumber: '',
                periodStart: '',
                periodEnd: '',
                paymentDate: '',
                status: ''
        }
        this.changerunNumberHandler = this.changerunNumberHandler.bind(this);
        this.changeperiodStartHandler = this.changeperiodStartHandler.bind(this);
        this.changeperiodEndHandler = this.changeperiodEndHandler.bind(this);
        this.changepaymentDateHandler = this.changepaymentDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdatePayrollRun = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            payrollRun.payrollRunId=''
            PayrollRunService.createPayrollRun(payrollRun).then(res =>{
                this.props.history.push('/payrollRuns');
            });
        }else{
            PayrollRunService.updatePayrollRun(payrollRun).then( res => {
                this.props.history.push('/payrollRuns');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PayrollRun</h3>
        }else{
            return <h3 className="text-center">Update PayrollRun</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> runNumber:&emsp; </label>
                                                <input placeholder="runNumber" name="runNumber" className="form-control" value={this.state.runNumber} onChange={this.changerunNumberHandler}/>

                                            <label> periodStart:&emsp; </label>
                                                <input type="date" placeholder="periodStart" name="periodStart" className="form-control" value={this.state.periodStart} onChange={this.changeperiodStartHandler}/>

                                            <label> periodEnd:&emsp; </label>
                                                <input type="date" placeholder="periodEnd" name="periodEnd" className="form-control" value={this.state.periodEnd} onChange={this.changeperiodEndHandler}/>

                                            <label> paymentDate:&emsp; </label>
                                                <input type="date" placeholder="paymentDate" name="paymentDate" className="form-control" value={this.state.paymentDate} onChange={this.changepaymentDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePayrollRun}>Save</button>
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

export default CreatePayrollRunComponent
