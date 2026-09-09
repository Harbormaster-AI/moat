import React, { Component } from 'react'
import TerminationService from '../services/TerminationService'

class ListTerminationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                terminations: []
        }
        this.addTermination = this.addTermination.bind(this);
        this.editTermination = this.editTermination.bind(this);
        this.deleteTermination = this.deleteTermination.bind(this);
    }

    deleteTermination(id){
        TerminationService.deleteTermination(id).then( res => {
            this.setState({terminations: this.state.terminations.filter(termination => termination.terminationId !== id)});
        });
    }
    viewTermination(id){
        this.props.history.push(`/view-termination/${id}`);
    }
    editTermination(id){
        this.props.history.push(`/add-termination/${id}`);
    }

    componentDidMount(){
        TerminationService.getTerminations().then((res) => {
            this.setState({ terminations: res.data});
        });
    }

    addTermination(){
        this.props.history.push('/add-termination/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Termination List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTermination}> Add Termination</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TerminationNumber </th>
                                    <th> TerminationDate </th>
                                    <th> Notes </th>
                                    <th> EligibleForRehire </th>
                                    <th> Reason </th>
                                    <th> Type </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.terminations.map(
                                        termination => 
                                        <tr key = {termination.terminationId}>
                                             <td> { termination.terminationNumber } </td>
                                             <td> { termination.terminationDate } </td>
                                             <td> { termination.notes } </td>
                                             <td> { termination.eligibleForRehire } </td>
                                             <td> { termination.reason } </td>
                                             <td> { termination.type } </td>
                                             <td>
                                                 <button onClick={ () => this.editTermination(termination.terminationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTermination(termination.terminationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTermination(termination.terminationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTerminationComponent
