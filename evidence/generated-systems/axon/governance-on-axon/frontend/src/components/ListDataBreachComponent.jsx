import React, { Component } from 'react'
import DataBreachService from '../services/DataBreachService'

class ListDataBreachComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataBreachs: []
        }
        this.addDataBreach = this.addDataBreach.bind(this);
        this.editDataBreach = this.editDataBreach.bind(this);
        this.deleteDataBreach = this.deleteDataBreach.bind(this);
    }

    deleteDataBreach(id){
        DataBreachService.deleteDataBreach(id).then( res => {
            this.setState({dataBreachs: this.state.dataBreachs.filter(dataBreach => dataBreach.dataBreachId !== id)});
        });
    }
    viewDataBreach(id){
        this.props.history.push(`/view-dataBreach/${id}`);
    }
    editDataBreach(id){
        this.props.history.push(`/add-dataBreach/${id}`);
    }

    componentDidMount(){
        DataBreachService.getDataBreachs().then((res) => {
            this.setState({ dataBreachs: res.data});
        });
    }

    addDataBreach(){
        this.props.history.push('/add-dataBreach/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataBreach List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataBreach}> Add DataBreach</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> IncidentDate </th>
                                    <th> Description </th>
                                    <th> RecordsAffected </th>
                                    <th> NotificationRequired </th>
                                    <th> Severity </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataBreachs.map(
                                        dataBreach => 
                                        <tr key = {dataBreach.dataBreachId}>
                                             <td> { dataBreach.incidentDate } </td>
                                             <td> { dataBreach.description } </td>
                                             <td> { dataBreach.recordsAffected } </td>
                                             <td> { dataBreach.notificationRequired } </td>
                                             <td> { dataBreach.severity } </td>
                                             <td> { dataBreach.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataBreach(dataBreach.dataBreachId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataBreach(dataBreach.dataBreachId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataBreach(dataBreach.dataBreachId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDataBreachComponent
