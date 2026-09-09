import React, { Component } from 'react'
import OutboundAllocationService from '../services/OutboundAllocationService'

class ListOutboundAllocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                outboundAllocations: []
        }
        this.addOutboundAllocation = this.addOutboundAllocation.bind(this);
        this.editOutboundAllocation = this.editOutboundAllocation.bind(this);
        this.deleteOutboundAllocation = this.deleteOutboundAllocation.bind(this);
    }

    deleteOutboundAllocation(id){
        OutboundAllocationService.deleteOutboundAllocation(id).then( res => {
            this.setState({outboundAllocations: this.state.outboundAllocations.filter(outboundAllocation => outboundAllocation.outboundAllocationId !== id)});
        });
    }
    viewOutboundAllocation(id){
        this.props.history.push(`/view-outboundAllocation/${id}`);
    }
    editOutboundAllocation(id){
        this.props.history.push(`/add-outboundAllocation/${id}`);
    }

    componentDidMount(){
        OutboundAllocationService.getOutboundAllocations().then((res) => {
            this.setState({ outboundAllocations: res.data});
        });
    }

    addOutboundAllocation(){
        this.props.history.push('/add-outboundAllocation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">OutboundAllocation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOutboundAllocation}> Add OutboundAllocation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AllocationNumber </th>
                                    <th> AllocatedQuantity </th>
                                    <th> AllocationDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.outboundAllocations.map(
                                        outboundAllocation => 
                                        <tr key = {outboundAllocation.outboundAllocationId}>
                                             <td> { outboundAllocation.allocationNumber } </td>
                                             <td> { outboundAllocation.allocatedQuantity } </td>
                                             <td> { outboundAllocation.allocationDate } </td>
                                             <td> { outboundAllocation.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editOutboundAllocation(outboundAllocation.outboundAllocationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOutboundAllocation(outboundAllocation.outboundAllocationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOutboundAllocation(outboundAllocation.outboundAllocationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOutboundAllocationComponent
