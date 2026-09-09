import React, { Component } from 'react'
import ChargebackService from '../services/ChargebackService';

class UpdateChargebackComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                chargebackReference: '',
                amount: '',
                postedAt: '',
                stage: '',
                status: ''
        }
        this.updateChargeback = this.updateChargeback.bind(this);

        this.changechargebackReferenceHandler = this.changechargebackReferenceHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changepostedAtHandler = this.changepostedAtHandler.bind(this);
        this.changeStageHandler = this.changeStageHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ChargebackService.getChargebackById(this.state.id).then( (res) =>{
            let chargeback = res.data;
            this.setState({
                chargebackReference: chargeback.chargebackReference,
                amount: chargeback.amount,
                postedAt: chargeback.postedAt,
                stage: chargeback.stage,
                status: chargeback.status
            });
        });
    }

    updateChargeback = (e) => {
        e.preventDefault();
        let chargeback = {
            chargebackId: this.state.id,
            chargebackReference: this.state.chargebackReference,
            amount: this.state.amount,
            postedAt: this.state.postedAt,
            stage: this.state.stage,
            status: this.state.status
        };
        console.log('chargeback => ' + JSON.stringify(chargeback));
        console.log('id => ' + JSON.stringify(this.state.id));
        ChargebackService.updateChargeback(chargeback).then( res => {
            this.props.history.push('/chargebacks');
        });
    }

    changechargebackReferenceHandler= (event) => {
        this.setState({chargebackReference: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changepostedAtHandler= (event) => {
        this.setState({postedAt: event.target.value});
    }
    changeStageHandler= (event) => {
        this.setState({stage: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/chargebacks');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Chargeback</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> chargebackReference: </label>
                                                <input placeholder="chargebackReference" name="chargebackReference" className="form-control" value={this.state.chargebackReference} onChange={this.changechargebackReferenceHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> postedAt: </label>
                                                <input type="time" placeholder="postedAt" name="postedAt" className="form-control" value={this.state.postedAt} onChange={this.changepostedAtHandler}/>

                                            <label> Stage: </label>
                                                <select value={this.state.stage} onChange={this.changeStageHandler}>
                      <option name="Stage" className="form-control" >
                          FirstChargeback
                      </option>
                      <option name="Stage" className="form-control" >
                          SecondChargeback
                      </option>
                      <option name="Stage" className="form-control" >
                          Arbitration
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Accepted
                      </option>
                      <option name="Status" className="form-control" >
                          Reversed
                      </option>
                      <option name="Status" className="form-control" >
                          Lost
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateChargeback}>Save</button>
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

export default UpdateChargebackComponent
