import React, { Component } from 'react'
import NonconformanceService from '../services/NonconformanceService'

class ListNonconformanceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                nonconformances: []
        }
        this.addNonconformance = this.addNonconformance.bind(this);
        this.editNonconformance = this.editNonconformance.bind(this);
        this.deleteNonconformance = this.deleteNonconformance.bind(this);
    }

    deleteNonconformance(id){
        NonconformanceService.deleteNonconformance(id).then( res => {
            this.setState({nonconformances: this.state.nonconformances.filter(nonconformance => nonconformance.nonconformanceId !== id)});
        });
    }
    viewNonconformance(id){
        this.props.history.push(`/view-nonconformance/${id}`);
    }
    editNonconformance(id){
        this.props.history.push(`/add-nonconformance/${id}`);
    }

    componentDidMount(){
        NonconformanceService.getNonconformances().then((res) => {
            this.setState({ nonconformances: res.data});
        });
    }

    addNonconformance(){
        this.props.history.push('/add-nonconformance/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Nonconformance List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addNonconformance}> Add Nonconformance</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> NcNumber </th>
                                    <th> Description </th>
                                    <th> ContainmentAction </th>
                                    <th> NcType </th>
                                    <th> Severity </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.nonconformances.map(
                                        nonconformance => 
                                        <tr key = {nonconformance.nonconformanceId}>
                                             <td> { nonconformance.ncNumber } </td>
                                             <td> { nonconformance.description } </td>
                                             <td> { nonconformance.containmentAction } </td>
                                             <td> { nonconformance.ncType } </td>
                                             <td> { nonconformance.severity } </td>
                                             <td> { nonconformance.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editNonconformance(nonconformance.nonconformanceId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteNonconformance(nonconformance.nonconformanceId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewNonconformance(nonconformance.nonconformanceId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListNonconformanceComponent
