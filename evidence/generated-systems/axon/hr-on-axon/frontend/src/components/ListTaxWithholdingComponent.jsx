import React, { Component } from 'react'
import TaxWithholdingService from '../services/TaxWithholdingService'

class ListTaxWithholdingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                taxWithholdings: []
        }
        this.addTaxWithholding = this.addTaxWithholding.bind(this);
        this.editTaxWithholding = this.editTaxWithholding.bind(this);
        this.deleteTaxWithholding = this.deleteTaxWithholding.bind(this);
    }

    deleteTaxWithholding(id){
        TaxWithholdingService.deleteTaxWithholding(id).then( res => {
            this.setState({taxWithholdings: this.state.taxWithholdings.filter(taxWithholding => taxWithholding.taxWithholdingId !== id)});
        });
    }
    viewTaxWithholding(id){
        this.props.history.push(`/view-taxWithholding/${id}`);
    }
    editTaxWithholding(id){
        this.props.history.push(`/add-taxWithholding/${id}`);
    }

    componentDidMount(){
        TaxWithholdingService.getTaxWithholdings().then((res) => {
            this.setState({ taxWithholdings: res.data});
        });
    }

    addTaxWithholding(){
        this.props.history.push('/add-taxWithholding/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TaxWithholding List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTaxWithholding}> Add TaxWithholding</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TaxId </th>
                                    <th> Allowances </th>
                                    <th> AdditionalAmount </th>
                                    <th> FilingStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.taxWithholdings.map(
                                        taxWithholding => 
                                        <tr key = {taxWithholding.taxWithholdingId}>
                                             <td> { taxWithholding.taxId } </td>
                                             <td> { taxWithholding.allowances } </td>
                                             <td> { taxWithholding.additionalAmount } </td>
                                             <td> { taxWithholding.filingStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editTaxWithholding(taxWithholding.taxWithholdingId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTaxWithholding(taxWithholding.taxWithholdingId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTaxWithholding(taxWithholding.taxWithholdingId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTaxWithholdingComponent
