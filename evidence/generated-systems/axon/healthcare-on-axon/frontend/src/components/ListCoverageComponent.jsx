import React, { Component } from 'react'
import CoverageService from '../services/CoverageService'

class ListCoverageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                coverages: []
        }
        this.addCoverage = this.addCoverage.bind(this);
        this.editCoverage = this.editCoverage.bind(this);
        this.deleteCoverage = this.deleteCoverage.bind(this);
    }

    deleteCoverage(id){
        CoverageService.deleteCoverage(id).then( res => {
            this.setState({coverages: this.state.coverages.filter(coverage => coverage.coverageId !== id)});
        });
    }
    viewCoverage(id){
        this.props.history.push(`/view-coverage/${id}`);
    }
    editCoverage(id){
        this.props.history.push(`/add-coverage/${id}`);
    }

    componentDidMount(){
        CoverageService.getCoverages().then((res) => {
            this.setState({ coverages: res.data});
        });
    }

    addCoverage(){
        this.props.history.push('/add-coverage/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Coverage List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCoverage}> Add Coverage</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> MemberId </th>
                                    <th> GroupNumber </th>
                                    <th> EffectiveDate </th>
                                    <th> EndDate </th>
                                    <th> CoverageType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.coverages.map(
                                        coverage => 
                                        <tr key = {coverage.coverageId}>
                                             <td> { coverage.memberId } </td>
                                             <td> { coverage.groupNumber } </td>
                                             <td> { coverage.effectiveDate } </td>
                                             <td> { coverage.endDate } </td>
                                             <td> { coverage.coverageType } </td>
                                             <td>
                                                 <button onClick={ () => this.editCoverage(coverage.coverageId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCoverage(coverage.coverageId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCoverage(coverage.coverageId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCoverageComponent
