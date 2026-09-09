import React, { Component } from 'react'
import CarrierServiceService from '../services/CarrierServiceService';

class UpdateCarrierServiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                code: '',
                carrier: '',
                serviceLevel: ''
        }
        this.updateCarrierService = this.updateCarrierService.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeCarrierHandler = this.changeCarrierHandler.bind(this);
        this.changeServiceLevelHandler = this.changeServiceLevelHandler.bind(this);
    }

    componentDidMount(){
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

    updateCarrierService = (e) => {
        e.preventDefault();
        let carrierService = {
            carrierServiceId: this.state.id,
            name: this.state.name,
            code: this.state.code,
            carrier: this.state.carrier,
            serviceLevel: this.state.serviceLevel
        };
        console.log('carrierService => ' + JSON.stringify(carrierService));
        console.log('id => ' + JSON.stringify(this.state.id));
        CarrierServiceService.updateCarrierService(carrierService).then( res => {
            this.props.history.push('/carrierServices');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CarrierService</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> Carrier: </label>
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

                                            <label> ServiceLevel: </label>
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
                                        <button className="btn btn-success" onClick={this.updateCarrierService}>Save</button>
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

export default UpdateCarrierServiceComponent
