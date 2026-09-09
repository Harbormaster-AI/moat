import React, { Component } from 'react'
import ClaimService from '../services/ClaimService'

class ListClaimComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                claims: []
        }
        this.addClaim = this.addClaim.bind(this);
        this.editClaim = this.editClaim.bind(this);
        this.deleteClaim = this.deleteClaim.bind(this);
    }

    deleteClaim(id){
        ClaimService.deleteClaim(id).then( res => {
            this.setState({claims: this.state.claims.filter(claim => claim.claimId !== id)});
        });
    }
    viewClaim(id){
        this.props.history.push(`/view-claim/${id}`);
    }
    editClaim(id){
        this.props.history.push(`/add-claim/${id}`);
    }

    componentDidMount(){
        ClaimService.getClaims().then((res) => {
            this.setState({ claims: res.data});
        });
    }

    addClaim(){
        this.props.history.push('/add-claim/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Claim List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addClaim}> Add Claim</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ClaimNumber </th>
                                    <th> NoticeDate </th>
                                    <th> LossDate </th>
                                    <th> ReportedBy </th>
                                    <th> Status </th>
                                    <th> LossCause </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.claims.map(
                                        claim => 
                                        <tr key = {claim.claimId}>
                                             <td> { claim.claimNumber } </td>
                                             <td> { claim.noticeDate } </td>
                                             <td> { claim.lossDate } </td>
                                             <td> { claim.reportedBy } </td>
                                             <td> { claim.status } </td>
                                             <td> { claim.lossCause } </td>
                                             <td>
                                                 <button onClick={ () => this.editClaim(claim.claimId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteClaim(claim.claimId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewClaim(claim.claimId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListClaimComponent
