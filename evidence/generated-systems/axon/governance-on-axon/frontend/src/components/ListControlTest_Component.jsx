import React, { Component } from 'react'
import ControlTest_Service from '../services/ControlTest_Service'

class ListControlTest_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
                controlTest_s: []
        }
        this.addControlTest_ = this.addControlTest_.bind(this);
        this.editControlTest_ = this.editControlTest_.bind(this);
        this.deleteControlTest_ = this.deleteControlTest_.bind(this);
    }

    deleteControlTest_(id){
        ControlTest_Service.deleteControlTest_(id).then( res => {
            this.setState({controlTest_s: this.state.controlTest_s.filter(controlTest_ => controlTest_.controlTest_Id !== id)});
        });
    }
    viewControlTest_(id){
        this.props.history.push(`/view-controlTest_/${id}`);
    }
    editControlTest_(id){
        this.props.history.push(`/add-controlTest_/${id}`);
    }

    componentDidMount(){
        ControlTest_Service.getControlTest_s().then((res) => {
            this.setState({ controlTest_s: res.data});
        });
    }

    addControlTest_(){
        this.props.history.push('/add-controlTest_/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ControlTest_ List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addControlTest_}> Add ControlTest_</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> TestPeriodStart </th>
                                    <th> TestPeriodEnd </th>
                                    <th> SampleSize </th>
                                    <th> TestType </th>
                                    <th> Effectiveness </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.controlTest_s.map(
                                        controlTest_ => 
                                        <tr key = {controlTest_.controlTest_Id}>
                                             <td> { controlTest_.name } </td>
                                             <td> { controlTest_.testPeriodStart } </td>
                                             <td> { controlTest_.testPeriodEnd } </td>
                                             <td> { controlTest_.sampleSize } </td>
                                             <td> { controlTest_.testType } </td>
                                             <td> { controlTest_.effectiveness } </td>
                                             <td> { controlTest_.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editControlTest_(controlTest_.controlTest_Id)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteControlTest_(controlTest_.controlTest_Id)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewControlTest_(controlTest_.controlTest_Id)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListControlTest_Component
