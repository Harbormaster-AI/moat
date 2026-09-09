import React, { Component } from 'react'
import OpportunityService from '../services/OpportunityService'

class ListOpportunityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                opportunitys: []
        }
        this.addOpportunity = this.addOpportunity.bind(this);
        this.editOpportunity = this.editOpportunity.bind(this);
        this.deleteOpportunity = this.deleteOpportunity.bind(this);
    }

    deleteOpportunity(id){
        OpportunityService.deleteOpportunity(id).then( res => {
            this.setState({opportunitys: this.state.opportunitys.filter(opportunity => opportunity.opportunityId !== id)});
        });
    }
    viewOpportunity(id){
        this.props.history.push(`/view-opportunity/${id}`);
    }
    editOpportunity(id){
        this.props.history.push(`/add-opportunity/${id}`);
    }

    componentDidMount(){
        OpportunityService.getOpportunitys().then((res) => {
            this.setState({ opportunitys: res.data});
        });
    }

    addOpportunity(){
        this.props.history.push('/add-opportunity/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Opportunity List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOpportunity}> Add Opportunity</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Amount </th>
                                    <th> CloseDate </th>
                                    <th> Probability </th>
                                    <th> Description </th>
                                    <th> Stage </th>
                                    <th> Type </th>
                                    <th> ForecastCategory </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.opportunitys.map(
                                        opportunity => 
                                        <tr key = {opportunity.opportunityId}>
                                             <td> { opportunity.name } </td>
                                             <td> { opportunity.amount } </td>
                                             <td> { opportunity.closeDate } </td>
                                             <td> { opportunity.probability } </td>
                                             <td> { opportunity.description } </td>
                                             <td> { opportunity.stage } </td>
                                             <td> { opportunity.type } </td>
                                             <td> { opportunity.forecastCategory } </td>
                                             <td>
                                                 <button onClick={ () => this.editOpportunity(opportunity.opportunityId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOpportunity(opportunity.opportunityId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOpportunity(opportunity.opportunityId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOpportunityComponent
