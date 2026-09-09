import React, { Component } from 'react'
import AppliedFeeService from '../services/AppliedFeeService'

class ListAppliedFeeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                appliedFees: []
        }
        this.addAppliedFee = this.addAppliedFee.bind(this);
        this.editAppliedFee = this.editAppliedFee.bind(this);
        this.deleteAppliedFee = this.deleteAppliedFee.bind(this);
    }

    deleteAppliedFee(id){
        AppliedFeeService.deleteAppliedFee(id).then( res => {
            this.setState({appliedFees: this.state.appliedFees.filter(appliedFee => appliedFee.appliedFeeId !== id)});
        });
    }
    viewAppliedFee(id){
        this.props.history.push(`/view-appliedFee/${id}`);
    }
    editAppliedFee(id){
        this.props.history.push(`/add-appliedFee/${id}`);
    }

    componentDidMount(){
        AppliedFeeService.getAppliedFees().then((res) => {
            this.setState({ appliedFees: res.data});
        });
    }

    addAppliedFee(){
        this.props.history.push('/add-appliedFee/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AppliedFee List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAppliedFee}> Add AppliedFee</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Amount </th>
                                    <th> Description </th>
                                    <th> FeeType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.appliedFees.map(
                                        appliedFee => 
                                        <tr key = {appliedFee.appliedFeeId}>
                                             <td> { appliedFee.amount } </td>
                                             <td> { appliedFee.description } </td>
                                             <td> { appliedFee.feeType } </td>
                                             <td>
                                                 <button onClick={ () => this.editAppliedFee(appliedFee.appliedFeeId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAppliedFee(appliedFee.appliedFeeId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAppliedFee(appliedFee.appliedFeeId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAppliedFeeComponent
