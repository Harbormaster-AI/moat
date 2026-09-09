import React, { Component } from 'react'
import EndorsementService from '../services/EndorsementService'

class ListEndorsementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                endorsements: []
        }
        this.addEndorsement = this.addEndorsement.bind(this);
        this.editEndorsement = this.editEndorsement.bind(this);
        this.deleteEndorsement = this.deleteEndorsement.bind(this);
    }

    deleteEndorsement(id){
        EndorsementService.deleteEndorsement(id).then( res => {
            this.setState({endorsements: this.state.endorsements.filter(endorsement => endorsement.endorsementId !== id)});
        });
    }
    viewEndorsement(id){
        this.props.history.push(`/view-endorsement/${id}`);
    }
    editEndorsement(id){
        this.props.history.push(`/add-endorsement/${id}`);
    }

    componentDidMount(){
        EndorsementService.getEndorsements().then((res) => {
            this.setState({ endorsements: res.data});
        });
    }

    addEndorsement(){
        this.props.history.push('/add-endorsement/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Endorsement List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEndorsement}> Add Endorsement</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EndorsementNumber </th>
                                    <th> EffectiveDate </th>
                                    <th> Description </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.endorsements.map(
                                        endorsement => 
                                        <tr key = {endorsement.endorsementId}>
                                             <td> { endorsement.endorsementNumber } </td>
                                             <td> { endorsement.effectiveDate } </td>
                                             <td> { endorsement.description } </td>
                                             <td>
                                                 <button onClick={ () => this.editEndorsement(endorsement.endorsementId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEndorsement(endorsement.endorsementId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEndorsement(endorsement.endorsementId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEndorsementComponent
