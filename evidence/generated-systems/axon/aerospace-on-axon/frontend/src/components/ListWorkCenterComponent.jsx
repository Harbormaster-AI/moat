import React, { Component } from 'react'
import WorkCenterService from '../services/WorkCenterService'

class ListWorkCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                workCenters: []
        }
        this.addWorkCenter = this.addWorkCenter.bind(this);
        this.editWorkCenter = this.editWorkCenter.bind(this);
        this.deleteWorkCenter = this.deleteWorkCenter.bind(this);
    }

    deleteWorkCenter(id){
        WorkCenterService.deleteWorkCenter(id).then( res => {
            this.setState({workCenters: this.state.workCenters.filter(workCenter => workCenter.workCenterId !== id)});
        });
    }
    viewWorkCenter(id){
        this.props.history.push(`/view-workCenter/${id}`);
    }
    editWorkCenter(id){
        this.props.history.push(`/add-workCenter/${id}`);
    }

    componentDidMount(){
        WorkCenterService.getWorkCenters().then((res) => {
            this.setState({ workCenters: res.data});
        });
    }

    addWorkCenter(){
        this.props.history.push('/add-workCenter/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">WorkCenter List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWorkCenter}> Add WorkCenter</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Capability </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.workCenters.map(
                                        workCenter => 
                                        <tr key = {workCenter.workCenterId}>
                                             <td> { workCenter.name } </td>
                                             <td> { workCenter.capability } </td>
                                             <td>
                                                 <button onClick={ () => this.editWorkCenter(workCenter.workCenterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWorkCenter(workCenter.workCenterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWorkCenter(workCenter.workCenterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWorkCenterComponent
