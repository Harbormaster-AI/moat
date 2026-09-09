import React, { Component } from 'react'
import ContractService from '../services/ContractService';

class UpdateContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                contractNumber: '',
                startDate: '',
                endDate: '',
                renewalTermMonths: '',
                autoRenew: '',
                status: ''
        }
        this.updateContract = this.updateContract.bind(this);

        this.changecontractNumberHandler = this.changecontractNumberHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changerenewalTermMonthsHandler = this.changerenewalTermMonthsHandler.bind(this);
        this.changeautoRenewHandler = this.changeautoRenewHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ContractService.getContractById(this.state.id).then( (res) =>{
            let contract = res.data;
            this.setState({
                contractNumber: contract.contractNumber,
                startDate: contract.startDate,
                endDate: contract.endDate,
                renewalTermMonths: contract.renewalTermMonths,
                autoRenew: contract.autoRenew,
                status: contract.status
            });
        });
    }

    updateContract = (e) => {
        e.preventDefault();
        let contract = {
            contractId: this.state.id,
            contractNumber: this.state.contractNumber,
            startDate: this.state.startDate,
            endDate: this.state.endDate,
            renewalTermMonths: this.state.renewalTermMonths,
            autoRenew: this.state.autoRenew,
            status: this.state.status
        };
        console.log('contract => ' + JSON.stringify(contract));
        console.log('id => ' + JSON.stringify(this.state.id));
        ContractService.updateContract(contract).then( res => {
            this.props.history.push('/contracts');
        });
    }

    changecontractNumberHandler= (event) => {
        this.setState({contractNumber: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changerenewalTermMonthsHandler= (event) => {
        this.setState({renewalTermMonths: event.target.value});
    }
    changeautoRenewHandler= (event) => {
        this.setState({autoRenew: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/contracts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Contract</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> contractNumber: </label>
                                                <input placeholder="contractNumber" name="contractNumber" className="form-control" value={this.state.contractNumber} onChange={this.changecontractNumberHandler}/>

                                            <label> startDate: </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate: </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> renewalTermMonths: </label>
                                                <input type="number" placeholder="renewalTermMonths" name="renewalTermMonths" className="form-control" value={this.state.renewalTermMonths} onChange={this.changerenewalTermMonthsHandler}/>

                                            <label> autoRenew: </label>
                                                <input type="checkbox" placeholder="autoRenew" name="autoRenew" className="form-control" value={this.state.autoRenew} onChange={this.changeautoRenewHandler}/>


                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                      <option name="Status" className="form-control" >
                          Terminated
                      </option>
                      <option name="Status" className="form-control" >
                          Renewed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateContract}>Save</button>
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

export default UpdateContractComponent
