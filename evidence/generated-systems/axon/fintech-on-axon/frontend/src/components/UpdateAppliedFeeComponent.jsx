import React, { Component } from 'react'
import AppliedFeeService from '../services/AppliedFeeService';

class UpdateAppliedFeeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                amount: '',
                description: '',
                feeType: ''
        }
        this.updateAppliedFee = this.updateAppliedFee.bind(this);

        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeFeeTypeHandler = this.changeFeeTypeHandler.bind(this);
    }

    componentDidMount(){
        AppliedFeeService.getAppliedFeeById(this.state.id).then( (res) =>{
            let appliedFee = res.data;
            this.setState({
                amount: appliedFee.amount,
                description: appliedFee.description,
                feeType: appliedFee.feeType
            });
        });
    }

    updateAppliedFee = (e) => {
        e.preventDefault();
        let appliedFee = {
            appliedFeeId: this.state.id,
            amount: this.state.amount,
            description: this.state.description,
            feeType: this.state.feeType
        };
        console.log('appliedFee => ' + JSON.stringify(appliedFee));
        console.log('id => ' + JSON.stringify(this.state.id));
        AppliedFeeService.updateAppliedFee(appliedFee).then( res => {
            this.props.history.push('/appliedFees');
        });
    }

    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeFeeTypeHandler= (event) => {
        this.setState({feeType: event.target.value});
    }

    cancel(){
        this.props.history.push('/appliedFees');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AppliedFee</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> FeeType: </label>
                                                <select value={this.state.feeType} onChange={this.changeFeeTypeHandler}>
                      <option name="FeeType" className="form-control" >
                          Fixed
                      </option>
                      <option name="FeeType" className="form-control" >
                          Percentage
                      </option>
                      <option name="FeeType" className="form-control" >
                          Tiered
                      </option>
                      <option name="FeeType" className="form-control" >
                          Interchange
                      </option>
                      <option name="FeeType" className="form-control" >
                          Network
                      </option>
                      <option name="FeeType" className="form-control" >
                          Chargeback
                      </option>
                      <option name="FeeType" className="form-control" >
                          ATM
                      </option>
                      <option name="FeeType" className="form-control" >
                          FX
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAppliedFee}>Save</button>
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

export default UpdateAppliedFeeComponent
