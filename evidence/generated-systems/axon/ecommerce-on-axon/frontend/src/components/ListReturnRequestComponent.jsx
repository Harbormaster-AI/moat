import React, { Component } from 'react'
import ReturnRequestService from '../services/ReturnRequestService'

class ListReturnRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                returnRequests: []
        }
        this.addReturnRequest = this.addReturnRequest.bind(this);
        this.editReturnRequest = this.editReturnRequest.bind(this);
        this.deleteReturnRequest = this.deleteReturnRequest.bind(this);
    }

    deleteReturnRequest(id){
        ReturnRequestService.deleteReturnRequest(id).then( res => {
            this.setState({returnRequests: this.state.returnRequests.filter(returnRequest => returnRequest.returnRequestId !== id)});
        });
    }
    viewReturnRequest(id){
        this.props.history.push(`/view-returnRequest/${id}`);
    }
    editReturnRequest(id){
        this.props.history.push(`/add-returnRequest/${id}`);
    }

    componentDidMount(){
        ReturnRequestService.getReturnRequests().then((res) => {
            this.setState({ returnRequests: res.data});
        });
    }

    addReturnRequest(){
        this.props.history.push('/add-returnRequest/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ReturnRequest List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addReturnRequest}> Add ReturnRequest</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReturnNumber </th>
                                    <th> CreatedAt </th>
                                    <th> RefundAmount </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.returnRequests.map(
                                        returnRequest => 
                                        <tr key = {returnRequest.returnRequestId}>
                                             <td> { returnRequest.returnNumber } </td>
                                             <td> { returnRequest.createdAt } </td>
                                             <td> { returnRequest.refundAmount } </td>
                                             <td> { returnRequest.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editReturnRequest(returnRequest.returnRequestId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteReturnRequest(returnRequest.returnRequestId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewReturnRequest(returnRequest.returnRequestId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListReturnRequestComponent
