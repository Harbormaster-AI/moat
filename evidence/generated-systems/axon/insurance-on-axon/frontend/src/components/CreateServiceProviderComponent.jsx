import React, { Component } from 'react'
import ServiceProviderService from '../services/ServiceProviderService';

class CreateServiceProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                taxId: '',
                providerType: '',
                networkStatus: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changeProviderTypeHandler = this.changeProviderTypeHandler.bind(this);
        this.changeNetworkStatusHandler = this.changeNetworkStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateServiceProvider = (e) => {
        e.preventDefault();
        let serviceProvider = {
                serviceProviderId: this.state.id,
                name: this.state.name,
                taxId: this.state.taxId,
                providerType: this.state.providerType,
                networkStatus: this.state.networkStatus
            };
        console.log('serviceProvider => ' + JSON.stringify(serviceProvider));

        // step 5
        if(this.state.id === '_add'){
            serviceProvider.serviceProviderId=''
            ServiceProviderService.createServiceProvider(serviceProvider).then(res =>{
                this.props.history.push('/serviceProviders');
            });
        }else{
            ServiceProviderService.updateServiceProvider(serviceProvider).then( res => {
                this.props.history.push('/serviceProviders');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ServiceProvider</h3>
        }else{
            return <h3 className="text-center">Update ServiceProvider</h3>
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

                                            <label> taxId:&emsp; </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> ProviderType:&emsp; </label>
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

                                            <label> NetworkStatus:&emsp; </label>
                                                <select value={this.state.networkStatus} onChange={this.changeNetworkStatusHandler}>
                      <option name="NetworkStatus" className="form-control" >
                          InNetwork
                      </option>
                      <option name="NetworkStatus" className="form-control" >
                          OutOfNetwork
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateServiceProvider}>Save</button>
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

export default CreateServiceProviderComponent
