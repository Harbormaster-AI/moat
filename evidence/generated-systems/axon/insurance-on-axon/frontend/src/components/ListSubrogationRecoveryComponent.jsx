import React, { Component } from 'react'
import SubrogationRecoveryService from '../services/SubrogationRecoveryService'

class ListSubrogationRecoveryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                subrogationRecoverys: []
        }
        this.addSubrogationRecovery = this.addSubrogationRecovery.bind(this);
        this.editSubrogationRecovery = this.editSubrogationRecovery.bind(this);
        this.deleteSubrogationRecovery = this.deleteSubrogationRecovery.bind(this);
    }

    deleteSubrogationRecovery(id){
        SubrogationRecoveryService.deleteSubrogationRecovery(id).then( res => {
            this.setState({subrogationRecoverys: this.state.subrogationRecoverys.filter(subrogationRecovery => subrogationRecovery.subrogationRecoveryId !== id)});
        });
    }
    viewSubrogationRecovery(id){
        this.props.history.push(`/view-subrogationRecovery/${id}`);
    }
    editSubrogationRecovery(id){
        this.props.history.push(`/add-subrogationRecovery/${id}`);
    }

    componentDidMount(){
        SubrogationRecoveryService.getSubrogationRecoverys().then((res) => {
            this.setState({ subrogationRecoverys: res.data});
        });
    }

    addSubrogationRecovery(){
        this.props.history.push('/add-subrogationRecovery/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SubrogationRecovery List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSubrogationRecovery}> Add SubrogationRecovery</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RecoveryReference </th>
                                    <th> Amount </th>
                                    <th> RecoveryDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.subrogationRecoverys.map(
                                        subrogationRecovery => 
                                        <tr key = {subrogationRecovery.subrogationRecoveryId}>
                                             <td> { subrogationRecovery.recoveryReference } </td>
                                             <td> { subrogationRecovery.amount } </td>
                                             <td> { subrogationRecovery.recoveryDate } </td>
                                             <td> { subrogationRecovery.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editSubrogationRecovery(subrogationRecovery.subrogationRecoveryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSubrogationRecovery(subrogationRecovery.subrogationRecoveryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSubrogationRecovery(subrogationRecovery.subrogationRecoveryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSubrogationRecoveryComponent
