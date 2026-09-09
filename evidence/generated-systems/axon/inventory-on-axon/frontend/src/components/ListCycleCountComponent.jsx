import React, { Component } from 'react'
import CycleCountService from '../services/CycleCountService'

class ListCycleCountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                cycleCounts: []
        }
        this.addCycleCount = this.addCycleCount.bind(this);
        this.editCycleCount = this.editCycleCount.bind(this);
        this.deleteCycleCount = this.deleteCycleCount.bind(this);
    }

    deleteCycleCount(id){
        CycleCountService.deleteCycleCount(id).then( res => {
            this.setState({cycleCounts: this.state.cycleCounts.filter(cycleCount => cycleCount.cycleCountId !== id)});
        });
    }
    viewCycleCount(id){
        this.props.history.push(`/view-cycleCount/${id}`);
    }
    editCycleCount(id){
        this.props.history.push(`/add-cycleCount/${id}`);
    }

    componentDidMount(){
        CycleCountService.getCycleCounts().then((res) => {
            this.setState({ cycleCounts: res.data});
        });
    }

    addCycleCount(){
        this.props.history.push('/add-cycleCount/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CycleCount List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCycleCount}> Add CycleCount</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CountNumber </th>
                                    <th> ScheduledDate </th>
                                    <th> PerformedDate </th>
                                    <th> ApprovedBy </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.cycleCounts.map(
                                        cycleCount => 
                                        <tr key = {cycleCount.cycleCountId}>
                                             <td> { cycleCount.countNumber } </td>
                                             <td> { cycleCount.scheduledDate } </td>
                                             <td> { cycleCount.performedDate } </td>
                                             <td> { cycleCount.approvedBy } </td>
                                             <td> { cycleCount.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCycleCount(cycleCount.cycleCountId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCycleCount(cycleCount.cycleCountId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCycleCount(cycleCount.cycleCountId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCycleCountComponent
