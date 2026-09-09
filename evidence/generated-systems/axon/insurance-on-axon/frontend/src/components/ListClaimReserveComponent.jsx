import React, { Component } from 'react'
import ClaimReserveService from '../services/ClaimReserveService'

class ListClaimReserveComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                claimReserves: []
        }
        this.addClaimReserve = this.addClaimReserve.bind(this);
        this.editClaimReserve = this.editClaimReserve.bind(this);
        this.deleteClaimReserve = this.deleteClaimReserve.bind(this);
    }

    deleteClaimReserve(id){
        ClaimReserveService.deleteClaimReserve(id).then( res => {
            this.setState({claimReserves: this.state.claimReserves.filter(claimReserve => claimReserve.claimReserveId !== id)});
        });
    }
    viewClaimReserve(id){
        this.props.history.push(`/view-claimReserve/${id}`);
    }
    editClaimReserve(id){
        this.props.history.push(`/add-claimReserve/${id}`);
    }

    componentDidMount(){
        ClaimReserveService.getClaimReserves().then((res) => {
            this.setState({ claimReserves: res.data});
        });
    }

    addClaimReserve(){
        this.props.history.push('/add-claimReserve/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ClaimReserve List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addClaimReserve}> Add ClaimReserve</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Amount </th>
                                    <th> SetDate </th>
                                    <th> ReserveType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.claimReserves.map(
                                        claimReserve => 
                                        <tr key = {claimReserve.claimReserveId}>
                                             <td> { claimReserve.amount } </td>
                                             <td> { claimReserve.setDate } </td>
                                             <td> { claimReserve.reserveType } </td>
                                             <td> { claimReserve.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editClaimReserve(claimReserve.claimReserveId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteClaimReserve(claimReserve.claimReserveId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewClaimReserve(claimReserve.claimReserveId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListClaimReserveComponent
