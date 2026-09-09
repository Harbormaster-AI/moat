import React, { Component } from 'react'
import DischargeService from '../services/DischargeService'

class ListDischargeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                discharges: []
        }
        this.addDischarge = this.addDischarge.bind(this);
        this.editDischarge = this.editDischarge.bind(this);
        this.deleteDischarge = this.deleteDischarge.bind(this);
    }

    deleteDischarge(id){
        DischargeService.deleteDischarge(id).then( res => {
            this.setState({discharges: this.state.discharges.filter(discharge => discharge.dischargeId !== id)});
        });
    }
    viewDischarge(id){
        this.props.history.push(`/view-discharge/${id}`);
    }
    editDischarge(id){
        this.props.history.push(`/add-discharge/${id}`);
    }

    componentDidMount(){
        DischargeService.getDischarges().then((res) => {
            this.setState({ discharges: res.data});
        });
    }

    addDischarge(){
        this.props.history.push('/add-discharge/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Discharge List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDischarge}> Add Discharge</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> DischargeDateTime </th>
                                    <th> Disposition </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.discharges.map(
                                        discharge => 
                                        <tr key = {discharge.dischargeId}>
                                             <td> { discharge.dischargeDateTime } </td>
                                             <td> { discharge.disposition } </td>
                                             <td>
                                                 <button onClick={ () => this.editDischarge(discharge.dischargeId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDischarge(discharge.dischargeId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDischarge(discharge.dischargeId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDischargeComponent
