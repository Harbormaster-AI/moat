import React, { Component } from 'react'
import PerformanceCycleService from '../services/PerformanceCycleService'

class ListPerformanceCycleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                performanceCycles: []
        }
        this.addPerformanceCycle = this.addPerformanceCycle.bind(this);
        this.editPerformanceCycle = this.editPerformanceCycle.bind(this);
        this.deletePerformanceCycle = this.deletePerformanceCycle.bind(this);
    }

    deletePerformanceCycle(id){
        PerformanceCycleService.deletePerformanceCycle(id).then( res => {
            this.setState({performanceCycles: this.state.performanceCycles.filter(performanceCycle => performanceCycle.performanceCycleId !== id)});
        });
    }
    viewPerformanceCycle(id){
        this.props.history.push(`/view-performanceCycle/${id}`);
    }
    editPerformanceCycle(id){
        this.props.history.push(`/add-performanceCycle/${id}`);
    }

    componentDidMount(){
        PerformanceCycleService.getPerformanceCycles().then((res) => {
            this.setState({ performanceCycles: res.data});
        });
    }

    addPerformanceCycle(){
        this.props.history.push('/add-performanceCycle/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PerformanceCycle List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPerformanceCycle}> Add PerformanceCycle</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.performanceCycles.map(
                                        performanceCycle => 
                                        <tr key = {performanceCycle.performanceCycleId}>
                                             <td> { performanceCycle.name } </td>
                                             <td> { performanceCycle.startDate } </td>
                                             <td> { performanceCycle.endDate } </td>
                                             <td> { performanceCycle.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPerformanceCycle(performanceCycle.performanceCycleId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePerformanceCycle(performanceCycle.performanceCycleId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPerformanceCycle(performanceCycle.performanceCycleId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPerformanceCycleComponent
