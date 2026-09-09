import React, { Component } from 'react'
import AppliedFeeService from '../services/AppliedFeeService';

class CreateAppliedFeeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                amount: '',
                description: '',
                feeType: ''
        }
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeFeeTypeHandler = this.changeFeeTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AppliedFeeService.getAppliedFeeById(this.state.id).then( (res) =>{
                let appliedFee = res.data;
                this.setState({
                    amount: appliedFee.amount,
                    description: appliedFee.description,
                    feeType: appliedFee.feeType
                });
            });
        }        
    }
    saveOrUpdateAppliedFee = (e) => {
        e.preventDefault();
        let appliedFee = {
                appliedFeeId: this.state.id,
                amount: this.state.amount,
                description: this.state.description,
                feeType: this.state.feeType
            };
        console.log('appliedFee => ' + JSON.stringify(appliedFee));

        // step 5
        if(this.state.id === '_add'){
            appliedFee.appliedFeeId=''
            AppliedFeeService.createAppliedFee(appliedFee).then(res =>{
                this.props.history.push('/appliedFees');
            });
        }else{
            AppliedFeeService.updateAppliedFee(appliedFee).then( res => {
                this.props.history.push('/appliedFees');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AppliedFee</h3>
        }else{
            return <h3 className="text-center">Update AppliedFee</h3>
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
                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> FeeType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAppliedFee}>Save</button>
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

export default CreateAppliedFeeComponent
