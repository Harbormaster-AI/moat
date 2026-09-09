import React, { Component } from 'react'
import CarrierServiceService from '../services/CarrierServiceService'

class ListCarrierServiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                carrierServices: []
        }
        this.addCarrierService = this.addCarrierService.bind(this);
        this.editCarrierService = this.editCarrierService.bind(this);
        this.deleteCarrierService = this.deleteCarrierService.bind(this);
    }

    deleteCarrierService(id){
        CarrierServiceService.deleteCarrierService(id).then( res => {
            this.setState({carrierServices: this.state.carrierServices.filter(carrierService => carrierService.carrierServiceId !== id)});
        });
    }
    viewCarrierService(id){
        this.props.history.push(`/view-carrierService/${id}`);
    }
    editCarrierService(id){
        this.props.history.push(`/add-carrierService/${id}`);
    }

    componentDidMount(){
        CarrierServiceService.getCarrierServices().then((res) => {
            this.setState({ carrierServices: res.data});
        });
    }

    addCarrierService(){
        this.props.history.push('/add-carrierService/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CarrierService List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCarrierService}> Add CarrierService</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Code </th>
                                    <th> Carrier </th>
                                    <th> ServiceLevel </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.carrierServices.map(
                                        carrierService => 
                                        <tr key = {carrierService.carrierServiceId}>
                                             <td> { carrierService.name } </td>
                                             <td> { carrierService.code } </td>
                                             <td> { carrierService.carrier } </td>
                                             <td> { carrierService.serviceLevel } </td>
                                             <td>
                                                 <button onClick={ () => this.editCarrierService(carrierService.carrierServiceId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCarrierService(carrierService.carrierServiceId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCarrierService(carrierService.carrierServiceId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCarrierServiceComponent
