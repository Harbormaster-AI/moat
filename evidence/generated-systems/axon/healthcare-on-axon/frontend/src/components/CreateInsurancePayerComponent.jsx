import React, { Component } from 'react'
import InsurancePayerService from '../services/InsurancePayerService';

class CreateInsurancePayerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                website: '',
                payerType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changePayerTypeHandler = this.changePayerTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InsurancePayerService.getInsurancePayerById(this.state.id).then( (res) =>{
                let insurancePayer = res.data;
                this.setState({
                    name: insurancePayer.name,
                    website: insurancePayer.website,
                    payerType: insurancePayer.payerType
                });
            });
        }        
    }
    saveOrUpdateInsurancePayer = (e) => {
        e.preventDefault();
        let insurancePayer = {
                insurancePayerId: this.state.id,
                name: this.state.name,
                website: this.state.website,
                payerType: this.state.payerType
            };
        console.log('insurancePayer => ' + JSON.stringify(insurancePayer));

        // step 5
        if(this.state.id === '_add'){
            insurancePayer.insurancePayerId=''
            InsurancePayerService.createInsurancePayer(insurancePayer).then(res =>{
                this.props.history.push('/insurancePayers');
            });
        }else{
            InsurancePayerService.updateInsurancePayer(insurancePayer).then( res => {
                this.props.history.push('/insurancePayers');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changePayerTypeHandler= (event) => {
        this.setState({payerType: event.target.value});
    }

    cancel(){
        this.props.history.push('/insurancePayers');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InsurancePayer</h3>
        }else{
            return <h3 className="text-center">Update InsurancePayer</h3>
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

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> PayerType:&emsp; </label>
                                                <select value={this.state.payerType} onChange={this.changePayerTypeHandler}>
                      <option name="PayerType" className="form-control" >
                          Commercial
                      </option>
                      <option name="PayerType" className="form-control" >
                          Government
                      </option>
                      <option name="PayerType" className="form-control" >
                          SelfInsured
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInsurancePayer}>Save</button>
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

export default CreateInsurancePayerComponent
