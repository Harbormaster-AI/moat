import React, { Component } from 'react'
import TaxRuleService from '../services/TaxRuleService'

class ListTaxRuleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                taxRules: []
        }
        this.addTaxRule = this.addTaxRule.bind(this);
        this.editTaxRule = this.editTaxRule.bind(this);
        this.deleteTaxRule = this.deleteTaxRule.bind(this);
    }

    deleteTaxRule(id){
        TaxRuleService.deleteTaxRule(id).then( res => {
            this.setState({taxRules: this.state.taxRules.filter(taxRule => taxRule.taxRuleId !== id)});
        });
    }
    viewTaxRule(id){
        this.props.history.push(`/view-taxRule/${id}`);
    }
    editTaxRule(id){
        this.props.history.push(`/add-taxRule/${id}`);
    }

    componentDidMount(){
        TaxRuleService.getTaxRules().then((res) => {
            this.setState({ taxRules: res.data});
        });
    }

    addTaxRule(){
        this.props.history.push('/add-taxRule/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TaxRule List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTaxRule}> Add TaxRule</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Country </th>
                                    <th> Region </th>
                                    <th> Rate </th>
                                    <th> TaxInclusive </th>
                                    <th> TaxClass </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.taxRules.map(
                                        taxRule => 
                                        <tr key = {taxRule.taxRuleId}>
                                             <td> { taxRule.name } </td>
                                             <td> { taxRule.country } </td>
                                             <td> { taxRule.region } </td>
                                             <td> { taxRule.rate } </td>
                                             <td> { taxRule.taxInclusive } </td>
                                             <td> { taxRule.taxClass } </td>
                                             <td>
                                                 <button onClick={ () => this.editTaxRule(taxRule.taxRuleId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTaxRule(taxRule.taxRuleId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTaxRule(taxRule.taxRuleId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTaxRuleComponent
