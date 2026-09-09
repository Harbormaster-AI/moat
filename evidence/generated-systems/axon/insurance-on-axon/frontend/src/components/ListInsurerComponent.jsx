import React, { Component } from 'react'
import InsurerService from '../services/InsurerService'

class ListInsurerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                insurers: []
        }
        this.addInsurer = this.addInsurer.bind(this);
        this.editInsurer = this.editInsurer.bind(this);
        this.deleteInsurer = this.deleteInsurer.bind(this);
    }

    deleteInsurer(id){
        InsurerService.deleteInsurer(id).then( res => {
            this.setState({insurers: this.state.insurers.filter(insurer => insurer.insurerId !== id)});
        });
    }
    viewInsurer(id){
        this.props.history.push(`/view-insurer/${id}`);
    }
    editInsurer(id){
        this.props.history.push(`/add-insurer/${id}`);
    }

    componentDidMount(){
        InsurerService.getInsurers().then((res) => {
            this.setState({ insurers: res.data});
        });
    }

    addInsurer(){
        this.props.history.push('/add-insurer/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Insurer List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInsurer}> Add Insurer</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LegalName </th>
                                    <th> DomicileCountry </th>
                                    <th> NaicNumber </th>
                                    <th> Website </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.insurers.map(
                                        insurer => 
                                        <tr key = {insurer.insurerId}>
                                             <td> { insurer.name } </td>
                                             <td> { insurer.legalName } </td>
                                             <td> { insurer.domicileCountry } </td>
                                             <td> { insurer.naicNumber } </td>
                                             <td> { insurer.website } </td>
                                             <td>
                                                 <button onClick={ () => this.editInsurer(insurer.insurerId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInsurer(insurer.insurerId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInsurer(insurer.insurerId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInsurerComponent
