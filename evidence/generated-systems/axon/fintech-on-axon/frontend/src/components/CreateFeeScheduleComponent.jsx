import React, { Component } from 'react'
import FeeScheduleService from '../services/FeeScheduleService';

class CreateFeeScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                amount: '',
                percentage: '',
                minimum: '',
                maximum: '',
                feeType: '',
                calculationMethod: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changepercentageHandler = this.changepercentageHandler.bind(this);
        this.changeminimumHandler = this.changeminimumHandler.bind(this);
        this.changemaximumHandler = this.changemaximumHandler.bind(this);
        this.changeFeeTypeHandler = this.changeFeeTypeHandler.bind(this);
        this.changeCalculationMethodHandler = this.changeCalculationMethodHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            FeeScheduleService.getFeeScheduleById(this.state.id).then( (res) =>{
                let feeSchedule = res.data;
                this.setState({
                    name: feeSchedule.name,
                    amount: feeSchedule.amount,
                    percentage: feeSchedule.percentage,
                    minimum: feeSchedule.minimum,
                    maximum: feeSchedule.maximum,
                    feeType: feeSchedule.feeType,
                    calculationMethod: feeSchedule.calculationMethod
                });
            });
        }        
    }
    saveOrUpdateFeeSchedule = (e) => {
        e.preventDefault();
        let feeSchedule = {
                feeScheduleId: this.state.id,
                name: this.state.name,
                amount: this.state.amount,
                percentage: this.state.percentage,
                minimum: this.state.minimum,
                maximum: this.state.maximum,
                feeType: this.state.feeType,
                calculationMethod: this.state.calculationMethod
            };
        console.log('feeSchedule => ' + JSON.stringify(feeSchedule));

        // step 5
        if(this.state.id === '_add'){
            feeSchedule.feeScheduleId=''
            FeeScheduleService.createFeeSchedule(feeSchedule).then(res =>{
                this.props.history.push('/feeSchedules');
            });
        }else{
            FeeScheduleService.updateFeeSchedule(feeSchedule).then( res => {
                this.props.history.push('/feeSchedules');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changepercentageHandler= (event) => {
        this.setState({percentage: event.target.value});
    }
    changeminimumHandler= (event) => {
        this.setState({minimum: event.target.value});
    }
    changemaximumHandler= (event) => {
        this.setState({maximum: event.target.value});
    }
    changeFeeTypeHandler= (event) => {
        this.setState({feeType: event.target.value});
    }
    changeCalculationMethodHandler= (event) => {
        this.setState({calculationMethod: event.target.value});
    }

    cancel(){
        this.props.history.push('/feeSchedules');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add FeeSchedule</h3>
        }else{
            return <h3 className="text-center">Update FeeSchedule</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> percentage:&emsp; </label>
                                                <input placeholder="percentage" name="percentage" className="form-control" value={this.state.percentage} onChange={this.changepercentageHandler}/>

                                            <label> minimum:&emsp; </label>
                                                <input placeholder="minimum" name="minimum" className="form-control" value={this.state.minimum} onChange={this.changeminimumHandler}/>

                                            <label> maximum:&emsp; </label>
                                                <input placeholder="maximum" name="maximum" className="form-control" value={this.state.maximum} onChange={this.changemaximumHandler}/>

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

                                            <label> CalculationMethod:&emsp; </label>
                                                <select value={this.state.calculationMethod} onChange={this.changeCalculationMethodHandler}>
                      <option name="CalculationMethod" className="form-control" >
                          PerTransaction
                      </option>
                      <option name="CalculationMethod" className="form-control" >
                          PerMonth
                      </option>
                      <option name="CalculationMethod" className="form-control" >
                          PerAnnum
                      </option>
                      <option name="CalculationMethod" className="form-control" >
                          Slab
                      </option>
                      <option name="CalculationMethod" className="form-control" >
                          Tiered
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateFeeSchedule}>Save</button>
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

export default CreateFeeScheduleComponent
