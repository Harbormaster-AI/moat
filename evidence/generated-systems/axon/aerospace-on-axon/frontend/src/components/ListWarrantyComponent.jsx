import React, { Component } from 'react'
import WarrantyService from '../services/WarrantyService'

class ListWarrantyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                warrantys: []
        }
        this.addWarranty = this.addWarranty.bind(this);
        this.editWarranty = this.editWarranty.bind(this);
        this.deleteWarranty = this.deleteWarranty.bind(this);
    }

    deleteWarranty(id){
        WarrantyService.deleteWarranty(id).then( res => {
            this.setState({warrantys: this.state.warrantys.filter(warranty => warranty.warrantyId !== id)});
        });
    }
    viewWarranty(id){
        this.props.history.push(`/view-warranty/${id}`);
    }
    editWarranty(id){
        this.props.history.push(`/add-warranty/${id}`);
    }

    componentDidMount(){
        WarrantyService.getWarrantys().then((res) => {
            this.setState({ warrantys: res.data});
        });
    }

    addWarranty(){
        this.props.history.push('/add-warranty/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Warranty List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWarranty}> Add Warranty</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CoverageMonths </th>
                                    <th> WarrantyType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.warrantys.map(
                                        warranty => 
                                        <tr key = {warranty.warrantyId}>
                                             <td> { warranty.coverageMonths } </td>
                                             <td> { warranty.warrantyType } </td>
                                             <td>
                                                 <button onClick={ () => this.editWarranty(warranty.warrantyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWarranty(warranty.warrantyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWarranty(warranty.warrantyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWarrantyComponent
