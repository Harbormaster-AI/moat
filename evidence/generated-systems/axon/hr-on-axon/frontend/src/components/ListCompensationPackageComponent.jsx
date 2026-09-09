import React, { Component } from 'react'
import CompensationPackageService from '../services/CompensationPackageService'

class ListCompensationPackageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                compensationPackages: []
        }
        this.addCompensationPackage = this.addCompensationPackage.bind(this);
        this.editCompensationPackage = this.editCompensationPackage.bind(this);
        this.deleteCompensationPackage = this.deleteCompensationPackage.bind(this);
    }

    deleteCompensationPackage(id){
        CompensationPackageService.deleteCompensationPackage(id).then( res => {
            this.setState({compensationPackages: this.state.compensationPackages.filter(compensationPackage => compensationPackage.compensationPackageId !== id)});
        });
    }
    viewCompensationPackage(id){
        this.props.history.push(`/view-compensationPackage/${id}`);
    }
    editCompensationPackage(id){
        this.props.history.push(`/add-compensationPackage/${id}`);
    }

    componentDidMount(){
        CompensationPackageService.getCompensationPackages().then((res) => {
            this.setState({ compensationPackages: res.data});
        });
    }

    addCompensationPackage(){
        this.props.history.push('/add-compensationPackage/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CompensationPackage List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCompensationPackage}> Add CompensationPackage</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EffectiveFrom </th>
                                    <th> EffectiveTo </th>
                                    <th> Currency </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.compensationPackages.map(
                                        compensationPackage => 
                                        <tr key = {compensationPackage.compensationPackageId}>
                                             <td> { compensationPackage.effectiveFrom } </td>
                                             <td> { compensationPackage.effectiveTo } </td>
                                             <td> { compensationPackage.currency } </td>
                                             <td>
                                                 <button onClick={ () => this.editCompensationPackage(compensationPackage.compensationPackageId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCompensationPackage(compensationPackage.compensationPackageId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCompensationPackage(compensationPackage.compensationPackageId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCompensationPackageComponent
