import React, { Component } from 'react'
import ServiceProviderService from '../services/ServiceProviderService';

class UpdateServiceProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                taxId: '',
                providerType: '',
                networkStatus: ''
        }
        this.updateServiceProvider = this.updateServiceProvider.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changeProviderTypeHandler = this.changeProviderTypeHandler.bind(this);
        this.changeNetworkStatusHandler = this.changeNetworkStatusHandler.bind(this);
    }

    componentDidMount(){
        ServiceProviderService.getServiceProviderById(this.state.id).then( (res) =>{
            let serviceProvider = res.data;
            this.setState({
                name: serviceProvider.name,
                taxId: serviceProvider.taxId,
                providerType: serviceProvider.providerType,
                networkStatus: serviceProvider.networkStatus
            });
        });
    }

    updateServiceProvider = (e) => {
        e.preventDefault();
        let serviceProvider = {
            serviceProviderId: this.state.id,
            name: this.state.name,
            taxId: this.state.taxId,
            providerType: this.state.providerType,
            networkStatus: this.state.networkStatus
        };
        console.log('serviceProvider => ' + JSON.stringify(serviceProvider));
        console.log('id => ' + JSON.stringify(this.state.id));
        ServiceProviderService.updateServiceProvider(serviceProvider).then( res => {
            this.props.history.push('/serviceProviders');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changetaxIdHandler= (event) => {
        this.setState({taxId: event.target.value});
    }
    changeProviderTypeHandler= (event) => {
        this.setState({providerType: event.target.value});
    }
    changeNetworkStatusHandler= (event) => {
        this.setState({networkStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/serviceProviders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ServiceProvider</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> taxId: </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> ProviderType: </label>
                                                <select value={this.state.providerType} onChange={this.changeProviderTypeHandler}>
                      <option name="ProviderType" className="form-control" >
                          RepairShop
                      </option>
                      <option name="ProviderType" className="form-control" >
                          Towing
                      </option>
                      <option name="ProviderType" className="form-control" >
                          MedicalProvider
                      </option>
                      <option name="ProviderType" className="form-control" >
                          Attorney
                      </option>
                      <option name="ProviderType" className="form-control" >
                          ForensicEngineer
                      </option>
                      <option name="ProviderType" className="form-control" >
                          RentalCar
                      </option>
                    </select>

                                            <label> NetworkStatus: </label>
                                                <select value={this.state.networkStatus} onChange={this.changeNetworkStatusHandler}>
                      <option name="NetworkStatus" className="form-control" >
                          InNetwork
                      </option>
                      <option name="NetworkStatus" className="form-control" >
                          OutOfNetwork
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateServiceProvider}>Save</button>
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

export default UpdateServiceProviderComponent
