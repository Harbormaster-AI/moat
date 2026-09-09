import React, { Component } from 'react'
import MRPRunService from '../services/MRPRunService'

class ListMRPRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                mRPRuns: []
        }
        this.addMRPRun = this.addMRPRun.bind(this);
        this.editMRPRun = this.editMRPRun.bind(this);
        this.deleteMRPRun = this.deleteMRPRun.bind(this);
    }

    deleteMRPRun(id){
        MRPRunService.deleteMRPRun(id).then( res => {
            this.setState({mRPRuns: this.state.mRPRuns.filter(mRPRun => mRPRun.mRPRunId !== id)});
        });
    }
    viewMRPRun(id){
        this.props.history.push(`/view-mRPRun/${id}`);
    }
    editMRPRun(id){
        this.props.history.push(`/add-mRPRun/${id}`);
    }

    componentDidMount(){
        MRPRunService.getMRPRuns().then((res) => {
            this.setState({ mRPRuns: res.data});
        });
    }

    addMRPRun(){
        this.props.history.push('/add-mRPRun/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MRPRun List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMRPRun}> Add MRPRun</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RunNumber </th>
                                    <th> RunDateTime </th>
                                    <th> PlanningHorizonDays </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.mRPRuns.map(
                                        mRPRun => 
                                        <tr key = {mRPRun.mRPRunId}>
                                             <td> { mRPRun.runNumber } </td>
                                             <td> { mRPRun.runDateTime } </td>
                                             <td> { mRPRun.planningHorizonDays } </td>
                                             <td> { mRPRun.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editMRPRun(mRPRun.mRPRunId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMRPRun(mRPRun.mRPRunId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMRPRun(mRPRun.mRPRunId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMRPRunComponent
