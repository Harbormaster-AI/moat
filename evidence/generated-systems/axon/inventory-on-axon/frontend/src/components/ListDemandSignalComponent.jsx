import React, { Component } from 'react'
import DemandSignalService from '../services/DemandSignalService'

class ListDemandSignalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                demandSignals: []
        }
        this.addDemandSignal = this.addDemandSignal.bind(this);
        this.editDemandSignal = this.editDemandSignal.bind(this);
        this.deleteDemandSignal = this.deleteDemandSignal.bind(this);
    }

    deleteDemandSignal(id){
        DemandSignalService.deleteDemandSignal(id).then( res => {
            this.setState({demandSignals: this.state.demandSignals.filter(demandSignal => demandSignal.demandSignalId !== id)});
        });
    }
    viewDemandSignal(id){
        this.props.history.push(`/view-demandSignal/${id}`);
    }
    editDemandSignal(id){
        this.props.history.push(`/add-demandSignal/${id}`);
    }

    componentDidMount(){
        DemandSignalService.getDemandSignals().then((res) => {
            this.setState({ demandSignals: res.data});
        });
    }

    addDemandSignal(){
        this.props.history.push('/add-demandSignal/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DemandSignal List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDemandSignal}> Add DemandSignal</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ExternalReference </th>
                                    <th> RequestedDate </th>
                                    <th> Quantity </th>
                                    <th> DemandType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.demandSignals.map(
                                        demandSignal => 
                                        <tr key = {demandSignal.demandSignalId}>
                                             <td> { demandSignal.externalReference } </td>
                                             <td> { demandSignal.requestedDate } </td>
                                             <td> { demandSignal.quantity } </td>
                                             <td> { demandSignal.demandType } </td>
                                             <td>
                                                 <button onClick={ () => this.editDemandSignal(demandSignal.demandSignalId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDemandSignal(demandSignal.demandSignalId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDemandSignal(demandSignal.demandSignalId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDemandSignalComponent
