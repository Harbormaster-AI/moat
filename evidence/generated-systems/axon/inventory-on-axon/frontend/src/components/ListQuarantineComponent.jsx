import React, { Component } from 'react'
import QuarantineService from '../services/QuarantineService'

class ListQuarantineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                quarantines: []
        }
        this.addQuarantine = this.addQuarantine.bind(this);
        this.editQuarantine = this.editQuarantine.bind(this);
        this.deleteQuarantine = this.deleteQuarantine.bind(this);
    }

    deleteQuarantine(id){
        QuarantineService.deleteQuarantine(id).then( res => {
            this.setState({quarantines: this.state.quarantines.filter(quarantine => quarantine.quarantineId !== id)});
        });
    }
    viewQuarantine(id){
        this.props.history.push(`/view-quarantine/${id}`);
    }
    editQuarantine(id){
        this.props.history.push(`/add-quarantine/${id}`);
    }

    componentDidMount(){
        QuarantineService.getQuarantines().then((res) => {
            this.setState({ quarantines: res.data});
        });
    }

    addQuarantine(){
        this.props.history.push('/add-quarantine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Quarantine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addQuarantine}> Add Quarantine</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Reason </th>
                                    <th> StartedAt </th>
                                    <th> ReleasedAt </th>
                                    <th> Disposition </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.quarantines.map(
                                        quarantine => 
                                        <tr key = {quarantine.quarantineId}>
                                             <td> { quarantine.reason } </td>
                                             <td> { quarantine.startedAt } </td>
                                             <td> { quarantine.releasedAt } </td>
                                             <td> { quarantine.disposition } </td>
                                             <td>
                                                 <button onClick={ () => this.editQuarantine(quarantine.quarantineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteQuarantine(quarantine.quarantineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewQuarantine(quarantine.quarantineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListQuarantineComponent
