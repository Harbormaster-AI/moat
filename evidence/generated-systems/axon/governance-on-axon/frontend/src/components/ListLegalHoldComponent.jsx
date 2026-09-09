import React, { Component } from 'react'
import LegalHoldService from '../services/LegalHoldService'

class ListLegalHoldComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                legalHolds: []
        }
        this.addLegalHold = this.addLegalHold.bind(this);
        this.editLegalHold = this.editLegalHold.bind(this);
        this.deleteLegalHold = this.deleteLegalHold.bind(this);
    }

    deleteLegalHold(id){
        LegalHoldService.deleteLegalHold(id).then( res => {
            this.setState({legalHolds: this.state.legalHolds.filter(legalHold => legalHold.legalHoldId !== id)});
        });
    }
    viewLegalHold(id){
        this.props.history.push(`/view-legalHold/${id}`);
    }
    editLegalHold(id){
        this.props.history.push(`/add-legalHold/${id}`);
    }

    componentDidMount(){
        LegalHoldService.getLegalHolds().then((res) => {
            this.setState({ legalHolds: res.data});
        });
    }

    addLegalHold(){
        this.props.history.push('/add-legalHold/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LegalHold List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLegalHold}> Add LegalHold</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Reason </th>
                                    <th> IssuedDate </th>
                                    <th> ReleaseDate </th>
                                    <th> HoldStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.legalHolds.map(
                                        legalHold => 
                                        <tr key = {legalHold.legalHoldId}>
                                             <td> { legalHold.name } </td>
                                             <td> { legalHold.reason } </td>
                                             <td> { legalHold.issuedDate } </td>
                                             <td> { legalHold.releaseDate } </td>
                                             <td> { legalHold.holdStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editLegalHold(legalHold.legalHoldId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLegalHold(legalHold.legalHoldId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLegalHold(legalHold.legalHoldId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLegalHoldComponent
