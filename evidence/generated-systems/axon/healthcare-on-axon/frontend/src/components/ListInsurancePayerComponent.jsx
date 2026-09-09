import React, { Component } from 'react'
import InsurancePayerService from '../services/InsurancePayerService'

class ListInsurancePayerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                insurancePayers: []
        }
        this.addInsurancePayer = this.addInsurancePayer.bind(this);
        this.editInsurancePayer = this.editInsurancePayer.bind(this);
        this.deleteInsurancePayer = this.deleteInsurancePayer.bind(this);
    }

    deleteInsurancePayer(id){
        InsurancePayerService.deleteInsurancePayer(id).then( res => {
            this.setState({insurancePayers: this.state.insurancePayers.filter(insurancePayer => insurancePayer.insurancePayerId !== id)});
        });
    }
    viewInsurancePayer(id){
        this.props.history.push(`/view-insurancePayer/${id}`);
    }
    editInsurancePayer(id){
        this.props.history.push(`/add-insurancePayer/${id}`);
    }

    componentDidMount(){
        InsurancePayerService.getInsurancePayers().then((res) => {
            this.setState({ insurancePayers: res.data});
        });
    }

    addInsurancePayer(){
        this.props.history.push('/add-insurancePayer/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InsurancePayer List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInsurancePayer}> Add InsurancePayer</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Website </th>
                                    <th> PayerType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.insurancePayers.map(
                                        insurancePayer => 
                                        <tr key = {insurancePayer.insurancePayerId}>
                                             <td> { insurancePayer.name } </td>
                                             <td> { insurancePayer.website } </td>
                                             <td> { insurancePayer.payerType } </td>
                                             <td>
                                                 <button onClick={ () => this.editInsurancePayer(insurancePayer.insurancePayerId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInsurancePayer(insurancePayer.insurancePayerId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInsurancePayer(insurancePayer.insurancePayerId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInsurancePayerComponent
