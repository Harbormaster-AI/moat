import React, { Component } from 'react'
import OpportunityStageHistoryService from '../services/OpportunityStageHistoryService'

class ListOpportunityStageHistoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                opportunityStageHistorys: []
        }
        this.addOpportunityStageHistory = this.addOpportunityStageHistory.bind(this);
        this.editOpportunityStageHistory = this.editOpportunityStageHistory.bind(this);
        this.deleteOpportunityStageHistory = this.deleteOpportunityStageHistory.bind(this);
    }

    deleteOpportunityStageHistory(id){
        OpportunityStageHistoryService.deleteOpportunityStageHistory(id).then( res => {
            this.setState({opportunityStageHistorys: this.state.opportunityStageHistorys.filter(opportunityStageHistory => opportunityStageHistory.opportunityStageHistoryId !== id)});
        });
    }
    viewOpportunityStageHistory(id){
        this.props.history.push(`/view-opportunityStageHistory/${id}`);
    }
    editOpportunityStageHistory(id){
        this.props.history.push(`/add-opportunityStageHistory/${id}`);
    }

    componentDidMount(){
        OpportunityStageHistoryService.getOpportunityStageHistorys().then((res) => {
            this.setState({ opportunityStageHistorys: res.data});
        });
    }

    addOpportunityStageHistory(){
        this.props.history.push('/add-opportunityStageHistory/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">OpportunityStageHistory List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOpportunityStageHistory}> Add OpportunityStageHistory</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ChangedAt </th>
                                    <th> Comment </th>
                                    <th> FromStage </th>
                                    <th> ToStage </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.opportunityStageHistorys.map(
                                        opportunityStageHistory => 
                                        <tr key = {opportunityStageHistory.opportunityStageHistoryId}>
                                             <td> { opportunityStageHistory.changedAt } </td>
                                             <td> { opportunityStageHistory.comment } </td>
                                             <td> { opportunityStageHistory.fromStage } </td>
                                             <td> { opportunityStageHistory.toStage } </td>
                                             <td>
                                                 <button onClick={ () => this.editOpportunityStageHistory(opportunityStageHistory.opportunityStageHistoryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOpportunityStageHistory(opportunityStageHistory.opportunityStageHistoryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOpportunityStageHistory(opportunityStageHistory.opportunityStageHistoryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOpportunityStageHistoryComponent
