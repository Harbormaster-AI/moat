import React, { Component } from 'react'
import PolicyCoverageService from '../services/PolicyCoverageService'

class ListPolicyCoverageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                policyCoverages: []
        }
        this.addPolicyCoverage = this.addPolicyCoverage.bind(this);
        this.editPolicyCoverage = this.editPolicyCoverage.bind(this);
        this.deletePolicyCoverage = this.deletePolicyCoverage.bind(this);
    }

    deletePolicyCoverage(id){
        PolicyCoverageService.deletePolicyCoverage(id).then( res => {
            this.setState({policyCoverages: this.state.policyCoverages.filter(policyCoverage => policyCoverage.policyCoverageId !== id)});
        });
    }
    viewPolicyCoverage(id){
        this.props.history.push(`/view-policyCoverage/${id}`);
    }
    editPolicyCoverage(id){
        this.props.history.push(`/add-policyCoverage/${id}`);
    }

    componentDidMount(){
        PolicyCoverageService.getPolicyCoverages().then((res) => {
            this.setState({ policyCoverages: res.data});
        });
    }

    addPolicyCoverage(){
        this.props.history.push('/add-policyCoverage/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PolicyCoverage List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPolicyCoverage}> Add PolicyCoverage</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Limit </th>
                                    <th> Deductible </th>
                                    <th> Premium </th>
                                    <th> CoverageType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.policyCoverages.map(
                                        policyCoverage => 
                                        <tr key = {policyCoverage.policyCoverageId}>
                                             <td> { policyCoverage.limit } </td>
                                             <td> { policyCoverage.deductible } </td>
                                             <td> { policyCoverage.premium } </td>
                                             <td> { policyCoverage.coverageType } </td>
                                             <td>
                                                 <button onClick={ () => this.editPolicyCoverage(policyCoverage.policyCoverageId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePolicyCoverage(policyCoverage.policyCoverageId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPolicyCoverage(policyCoverage.policyCoverageId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPolicyCoverageComponent
