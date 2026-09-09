import React, { Component } from 'react'
import DistributorService from '../services/DistributorService'

class ListDistributorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                distributors: []
        }
        this.addDistributor = this.addDistributor.bind(this);
        this.editDistributor = this.editDistributor.bind(this);
        this.deleteDistributor = this.deleteDistributor.bind(this);
    }

    deleteDistributor(id){
        DistributorService.deleteDistributor(id).then( res => {
            this.setState({distributors: this.state.distributors.filter(distributor => distributor.distributorId !== id)});
        });
    }
    viewDistributor(id){
        this.props.history.push(`/view-distributor/${id}`);
    }
    editDistributor(id){
        this.props.history.push(`/add-distributor/${id}`);
    }

    componentDidMount(){
        DistributorService.getDistributors().then((res) => {
            this.setState({ distributors: res.data});
        });
    }

    addDistributor(){
        this.props.history.push('/add-distributor/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Distributor List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDistributor}> Add Distributor</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LicenseNumber </th>
                                    <th> Region </th>
                                    <th> DistributorType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.distributors.map(
                                        distributor => 
                                        <tr key = {distributor.distributorId}>
                                             <td> { distributor.name } </td>
                                             <td> { distributor.licenseNumber } </td>
                                             <td> { distributor.region } </td>
                                             <td> { distributor.distributorType } </td>
                                             <td>
                                                 <button onClick={ () => this.editDistributor(distributor.distributorId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDistributor(distributor.distributorId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDistributor(distributor.distributorId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDistributorComponent
