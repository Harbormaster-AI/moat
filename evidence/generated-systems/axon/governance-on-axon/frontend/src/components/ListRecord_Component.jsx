import React, { Component } from 'react'
import Record_Service from '../services/Record_Service'

class ListRecord_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
                record_s: []
        }
        this.addRecord_ = this.addRecord_.bind(this);
        this.editRecord_ = this.editRecord_.bind(this);
        this.deleteRecord_ = this.deleteRecord_.bind(this);
    }

    deleteRecord_(id){
        Record_Service.deleteRecord_(id).then( res => {
            this.setState({record_s: this.state.record_s.filter(record_ => record_.record_Id !== id)});
        });
    }
    viewRecord_(id){
        this.props.history.push(`/view-record_/${id}`);
    }
    editRecord_(id){
        this.props.history.push(`/add-record_/${id}`);
    }

    componentDidMount(){
        Record_Service.getRecord_s().then((res) => {
            this.setState({ record_s: res.data});
        });
    }

    addRecord_(){
        this.props.history.push('/add-record_/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Record_ List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRecord_}> Add Record_</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> CreationDate </th>
                                    <th> RecordType </th>
                                    <th> Classification </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.record_s.map(
                                        record_ => 
                                        <tr key = {record_.record_Id}>
                                             <td> { record_.title } </td>
                                             <td> { record_.creationDate } </td>
                                             <td> { record_.recordType } </td>
                                             <td> { record_.classification } </td>
                                             <td> { record_.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editRecord_(record_.record_Id)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRecord_(record_.record_Id)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRecord_(record_.record_Id)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRecord_Component
