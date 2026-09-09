import React, { Component } from 'react'
import CarrierServiceService from '../services/CarrierServiceService';

class CreateCarrierServiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                code: '',
                carrier: '',
                serviceLevel: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeCarrierHandler = this.changeCarrierHandler.bind(this);
        this.changeServiceLevelHandler = this.changeServiceLevelHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CarrierServiceService.getCarrierServiceById(this.state.id).then( (res) =>{
                let carrierService = res.data;
                this.setState({
                    name: carrierService.name,
                    code: carrierService.code,
                    carrier: carrierService.carrier,
                    serviceLevel: carrierService.serviceLevel
                });
            });
        }        
    }
    saveOrUpdateCarrierService = (e) => {
        e.preventDefault();
        let carrierService = {
                carrierServiceId: this.state.id,
                name: this.state.name,
                code: this.state.code,
                carrier: this.state.carrier,
                serviceLevel: this.state.serviceLevel
            };
        console.log('carrierService => ' + JSON.stringify(carrierService));

        // step 5
        if(this.state.id === '_add'){
            carrierService.carrierServiceId=''
            CarrierServiceService.createCarrierService(carrierService).then(res =>{
                this.props.history.push('/carrierServices');
            });
        }else{
            CarrierServiceService.updateCarrierService(carrierService).then( res => {
                this.props.history.push('/carrierServices');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changeCarrierHandler= (event) => {
        this.setState({carrier: event.target.value});
    }
    changeServiceLevelHandler= (event) => {
        this.setState({serviceLevel: event.target.value});
    }

    cancel(){
        this.props.history.push('/carrierServices');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CarrierService</h3>
        }else{
            return <h3 className="text-center">Update CarrierService</h3>
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

                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> Carrier:&emsp; </label>
                                                <select value={this.state.carrier} onChange={this.changeCarrierHandler}>
                      <option name="Carrier" className="form-control" >
                          UPS
                      </option>
                      <option name="Carrier" className="form-control" >
                          FedEx
                      </option>
                      <option name="Carrier" className="form-control" >
                          USPS
                      </option>
                      <option name="Carrier" className="form-control" >
                          DHL
                      </option>
                      <option name="Carrier" className="form-control" >
                          RoyalMail
                      </option>
                      <option name="Carrier" className="form-control" >
                          CanadaPost
                      </option>
                      <option name="Carrier" className="form-control" >
                          LocalCourier
                      </option>
                      <option name="Carrier" className="form-control" >
                          Other
                      </option>
                    </select>

                                            <label> ServiceLevel:&emsp; </label>
                                                <select value={this.state.serviceLevel} onChange={this.changeServiceLevelHandler}>
                      <option name="ServiceLevel" className="form-control" >
                          Economy
                      </option>
                      <option name="ServiceLevel" className="form-control" >
                          Standard
                      </option>
                      <option name="ServiceLevel" className="form-control" >
                          Express
                      </option>
                      <option name="ServiceLevel" className="form-control" >
                          Priority
                      </option>
                      <option name="ServiceLevel" className="form-control" >
                          NextDay
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCarrierService}>Save</button>
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

export default CreateCarrierServiceComponent
