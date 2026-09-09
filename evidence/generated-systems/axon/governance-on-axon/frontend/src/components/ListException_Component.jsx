import React, { Component } from 'react'
import Exception_Service from '../services/Exception_Service'

class ListException_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
                exception_s: []
        }
        this.addException_ = this.addException_.bind(this);
        this.editException_ = this.editException_.bind(this);
        this.deleteException_ = this.deleteException_.bind(this);
    }

    deleteException_(id){
        Exception_Service.deleteException_(id).then( res => {
            this.setState({exception_s: this.state.exception_s.filter(exception_ => exception_.exception_Id !== id)});
        });
    }
    viewException_(id){
        this.props.history.push(`/view-exception_/${id}`);
    }
    editException_(id){
        this.props.history.push(`/add-exception_/${id}`);
    }

    componentDidMount(){
        Exception_Service.getException_s().then((res) => {
            this.setState({ exception_s: res.data});
        });
    }

    addException_(){
        this.props.history.push('/add-exception_/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Exception_ List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addException_}> Add Exception_</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> Justification </th>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> ExceptionType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.exception_s.map(
                                        exception_ => 
                                        <tr key = {exception_.exception_Id}>
                                             <td> { exception_.title } </td>
                                             <td> { exception_.justification } </td>
                                             <td> { exception_.startDate } </td>
                                             <td> { exception_.endDate } </td>
                                             <td> { exception_.exceptionType } </td>
                                             <td> { exception_.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editException_(exception_.exception_Id)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteException_(exception_.exception_Id)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewException_(exception_.exception_Id)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListException_Component
