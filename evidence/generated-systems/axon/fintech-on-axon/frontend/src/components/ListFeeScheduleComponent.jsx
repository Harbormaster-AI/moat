import React, { Component } from 'react'
import FeeScheduleService from '../services/FeeScheduleService'

class ListFeeScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                feeSchedules: []
        }
        this.addFeeSchedule = this.addFeeSchedule.bind(this);
        this.editFeeSchedule = this.editFeeSchedule.bind(this);
        this.deleteFeeSchedule = this.deleteFeeSchedule.bind(this);
    }

    deleteFeeSchedule(id){
        FeeScheduleService.deleteFeeSchedule(id).then( res => {
            this.setState({feeSchedules: this.state.feeSchedules.filter(feeSchedule => feeSchedule.feeScheduleId !== id)});
        });
    }
    viewFeeSchedule(id){
        this.props.history.push(`/view-feeSchedule/${id}`);
    }
    editFeeSchedule(id){
        this.props.history.push(`/add-feeSchedule/${id}`);
    }

    componentDidMount(){
        FeeScheduleService.getFeeSchedules().then((res) => {
            this.setState({ feeSchedules: res.data});
        });
    }

    addFeeSchedule(){
        this.props.history.push('/add-feeSchedule/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FeeSchedule List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFeeSchedule}> Add FeeSchedule</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Amount </th>
                                    <th> Percentage </th>
                                    <th> Minimum </th>
                                    <th> Maximum </th>
                                    <th> FeeType </th>
                                    <th> CalculationMethod </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.feeSchedules.map(
                                        feeSchedule => 
                                        <tr key = {feeSchedule.feeScheduleId}>
                                             <td> { feeSchedule.name } </td>
                                             <td> { feeSchedule.amount } </td>
                                             <td> { feeSchedule.percentage } </td>
                                             <td> { feeSchedule.minimum } </td>
                                             <td> { feeSchedule.maximum } </td>
                                             <td> { feeSchedule.feeType } </td>
                                             <td> { feeSchedule.calculationMethod } </td>
                                             <td>
                                                 <button onClick={ () => this.editFeeSchedule(feeSchedule.feeScheduleId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFeeSchedule(feeSchedule.feeScheduleId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFeeSchedule(feeSchedule.feeScheduleId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFeeScheduleComponent
