import React, { Component } from 'react'
import ProcedureService from '../services/ProcedureService'

class ListProcedureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                procedures: []
        }
        this.addProcedure = this.addProcedure.bind(this);
        this.editProcedure = this.editProcedure.bind(this);
        this.deleteProcedure = this.deleteProcedure.bind(this);
    }

    deleteProcedure(id){
        ProcedureService.deleteProcedure(id).then( res => {
            this.setState({procedures: this.state.procedures.filter(procedure => procedure.procedureId !== id)});
        });
    }
    viewProcedure(id){
        this.props.history.push(`/view-procedure/${id}`);
    }
    editProcedure(id){
        this.props.history.push(`/add-procedure/${id}`);
    }

    componentDidMount(){
        ProcedureService.getProcedures().then((res) => {
            this.setState({ procedures: res.data});
        });
    }

    addProcedure(){
        this.props.history.push('/add-procedure/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Procedure List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProcedure}> Add Procedure</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ProcedureCode </th>
                                    <th> StartDateTime </th>
                                    <th> EndDateTime </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.procedures.map(
                                        procedure => 
                                        <tr key = {procedure.procedureId}>
                                             <td> { procedure.procedureCode } </td>
                                             <td> { procedure.startDateTime } </td>
                                             <td> { procedure.endDateTime } </td>
                                             <td> { procedure.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editProcedure(procedure.procedureId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProcedure(procedure.procedureId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProcedure(procedure.procedureId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProcedureComponent
