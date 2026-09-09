import React, { Component } from 'react'
import FraudSignalService from '../services/FraudSignalService'

class ListFraudSignalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                fraudSignals: []
        }
        this.addFraudSignal = this.addFraudSignal.bind(this);
        this.editFraudSignal = this.editFraudSignal.bind(this);
        this.deleteFraudSignal = this.deleteFraudSignal.bind(this);
    }

    deleteFraudSignal(id){
        FraudSignalService.deleteFraudSignal(id).then( res => {
            this.setState({fraudSignals: this.state.fraudSignals.filter(fraudSignal => fraudSignal.fraudSignalId !== id)});
        });
    }
    viewFraudSignal(id){
        this.props.history.push(`/view-fraudSignal/${id}`);
    }
    editFraudSignal(id){
        this.props.history.push(`/add-fraudSignal/${id}`);
    }

    componentDidMount(){
        FraudSignalService.getFraudSignals().then((res) => {
            this.setState({ fraudSignals: res.data});
        });
    }

    addFraudSignal(){
        this.props.history.push('/add-fraudSignal/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FraudSignal List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFraudSignal}> Add FraudSignal</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> RuleLogic </th>
                                    <th> SignalType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.fraudSignals.map(
                                        fraudSignal => 
                                        <tr key = {fraudSignal.fraudSignalId}>
                                             <td> { fraudSignal.name } </td>
                                             <td> { fraudSignal.ruleLogic } </td>
                                             <td> { fraudSignal.signalType } </td>
                                             <td>
                                                 <button onClick={ () => this.editFraudSignal(fraudSignal.fraudSignalId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFraudSignal(fraudSignal.fraudSignalId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFraudSignal(fraudSignal.fraudSignalId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFraudSignalComponent
