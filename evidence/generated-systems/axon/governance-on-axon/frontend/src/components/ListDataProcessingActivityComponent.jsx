import React, { Component } from 'react'
import DataProcessingActivityService from '../services/DataProcessingActivityService'

class ListDataProcessingActivityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataProcessingActivitys: []
        }
        this.addDataProcessingActivity = this.addDataProcessingActivity.bind(this);
        this.editDataProcessingActivity = this.editDataProcessingActivity.bind(this);
        this.deleteDataProcessingActivity = this.deleteDataProcessingActivity.bind(this);
    }

    deleteDataProcessingActivity(id){
        DataProcessingActivityService.deleteDataProcessingActivity(id).then( res => {
            this.setState({dataProcessingActivitys: this.state.dataProcessingActivitys.filter(dataProcessingActivity => dataProcessingActivity.dataProcessingActivityId !== id)});
        });
    }
    viewDataProcessingActivity(id){
        this.props.history.push(`/view-dataProcessingActivity/${id}`);
    }
    editDataProcessingActivity(id){
        this.props.history.push(`/add-dataProcessingActivity/${id}`);
    }

    componentDidMount(){
        DataProcessingActivityService.getDataProcessingActivitys().then((res) => {
            this.setState({ dataProcessingActivitys: res.data});
        });
    }

    addDataProcessingActivity(){
        this.props.history.push('/add-dataProcessingActivity/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataProcessingActivity List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataProcessingActivity}> Add DataProcessingActivity</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Purpose </th>
                                    <th> StartDate </th>
                                    <th> LawfulBasis </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataProcessingActivitys.map(
                                        dataProcessingActivity => 
                                        <tr key = {dataProcessingActivity.dataProcessingActivityId}>
                                             <td> { dataProcessingActivity.name } </td>
                                             <td> { dataProcessingActivity.purpose } </td>
                                             <td> { dataProcessingActivity.startDate } </td>
                                             <td> { dataProcessingActivity.lawfulBasis } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataProcessingActivity(dataProcessingActivity.dataProcessingActivityId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataProcessingActivity(dataProcessingActivity.dataProcessingActivityId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataProcessingActivity(dataProcessingActivity.dataProcessingActivityId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDataProcessingActivityComponent
