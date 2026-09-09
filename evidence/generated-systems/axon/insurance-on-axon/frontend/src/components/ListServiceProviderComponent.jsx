import React, { Component } from 'react'
import ServiceProviderService from '../services/ServiceProviderService'

class ListServiceProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                serviceProviders: []
        }
        this.addServiceProvider = this.addServiceProvider.bind(this);
        this.editServiceProvider = this.editServiceProvider.bind(this);
        this.deleteServiceProvider = this.deleteServiceProvider.bind(this);
    }

    deleteServiceProvider(id){
        ServiceProviderService.deleteServiceProvider(id).then( res => {
            this.setState({serviceProviders: this.state.serviceProviders.filter(serviceProvider => serviceProvider.serviceProviderId !== id)});
        });
    }
    viewServiceProvider(id){
        this.props.history.push(`/view-serviceProvider/${id}`);
    }
    editServiceProvider(id){
        this.props.history.push(`/add-serviceProvider/${id}`);
    }

    componentDidMount(){
        ServiceProviderService.getServiceProviders().then((res) => {
            this.setState({ serviceProviders: res.data});
        });
    }

    addServiceProvider(){
        this.props.history.push('/add-serviceProvider/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ServiceProvider List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addServiceProvider}> Add ServiceProvider</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> TaxId </th>
                                    <th> ProviderType </th>
                                    <th> NetworkStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.serviceProviders.map(
                                        serviceProvider => 
                                        <tr key = {serviceProvider.serviceProviderId}>
                                             <td> { serviceProvider.name } </td>
                                             <td> { serviceProvider.taxId } </td>
                                             <td> { serviceProvider.providerType } </td>
                                             <td> { serviceProvider.networkStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editServiceProvider(serviceProvider.serviceProviderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteServiceProvider(serviceProvider.serviceProviderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewServiceProvider(serviceProvider.serviceProviderId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListServiceProviderComponent
