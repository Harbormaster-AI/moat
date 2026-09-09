import React, { Component } from 'react'
import ObligationService from '../services/ObligationService'

class ListObligationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                obligations: []
        }
        this.addObligation = this.addObligation.bind(this);
        this.editObligation = this.editObligation.bind(this);
        this.deleteObligation = this.deleteObligation.bind(this);
    }

    deleteObligation(id){
        ObligationService.deleteObligation(id).then( res => {
            this.setState({obligations: this.state.obligations.filter(obligation => obligation.obligationId !== id)});
        });
    }
    viewObligation(id){
        this.props.history.push(`/view-obligation/${id}`);
    }
    editObligation(id){
        this.props.history.push(`/add-obligation/${id}`);
    }

    componentDidMount(){
        ObligationService.getObligations().then((res) => {
            this.setState({ obligations: res.data});
        });
    }

    addObligation(){
        this.props.history.push('/add-obligation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Obligation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addObligation}> Add Obligation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReferenceNumber </th>
                                    <th> DescriptionText </th>
                                    <th> ObligationType </th>
                                    <th> ReviewFrequency </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.obligations.map(
                                        obligation => 
                                        <tr key = {obligation.obligationId}>
                                             <td> { obligation.referenceNumber } </td>
                                             <td> { obligation.descriptionText } </td>
                                             <td> { obligation.obligationType } </td>
                                             <td> { obligation.reviewFrequency } </td>
                                             <td>
                                                 <button onClick={ () => this.editObligation(obligation.obligationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteObligation(obligation.obligationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewObligation(obligation.obligationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListObligationComponent
