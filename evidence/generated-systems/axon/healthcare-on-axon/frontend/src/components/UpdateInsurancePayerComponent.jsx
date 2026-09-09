import React, { Component } from 'react'
import InsurancePayerService from '../services/InsurancePayerService';

class UpdateInsurancePayerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                website: '',
                payerType: ''
        }
        this.updateInsurancePayer = this.updateInsurancePayer.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changePayerTypeHandler = this.changePayerTypeHandler.bind(this);
    }

    componentDidMount(){
        InsurancePayerService.getInsurancePayerById(this.state.id).then( (res) =>{
            let insurancePayer = res.data;
            this.setState({
                name: insurancePayer.name,
                website: insurancePayer.website,
                payerType: insurancePayer.payerType
            });
        });
    }

    updateInsurancePayer = (e) => {
        e.preventDefault();
        let insurancePayer = {
            insurancePayerId: this.state.id,
            name: this.state.name,
            website: this.state.website,
            payerType: this.state.payerType
        };
        console.log('insurancePayer => ' + JSON.stringify(insurancePayer));
        console.log('id => ' + JSON.stringify(this.state.id));
        InsurancePayerService.updateInsurancePayer(insurancePayer).then( res => {
            this.props.history.push('/insurancePayers');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InsurancePayer</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> PayerType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateInsurancePayer}>Save</button>
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

export default UpdateInsurancePayerComponent
