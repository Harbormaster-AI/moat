import React, { Component } from 'react'
import AnalyticsWorkspaceService from '../services/AnalyticsWorkspaceService'

class ListAnalyticsWorkspaceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                analyticsWorkspaces: []
        }
        this.addAnalyticsWorkspace = this.addAnalyticsWorkspace.bind(this);
        this.editAnalyticsWorkspace = this.editAnalyticsWorkspace.bind(this);
        this.deleteAnalyticsWorkspace = this.deleteAnalyticsWorkspace.bind(this);
    }

    deleteAnalyticsWorkspace(id){
        AnalyticsWorkspaceService.deleteAnalyticsWorkspace(id).then( res => {
            this.setState({analyticsWorkspaces: this.state.analyticsWorkspaces.filter(analyticsWorkspace => analyticsWorkspace.analyticsWorkspaceId !== id)});
        });
    }
    viewAnalyticsWorkspace(id){
        this.props.history.push(`/view-analyticsWorkspace/${id}`);
    }
    editAnalyticsWorkspace(id){
        this.props.history.push(`/add-analyticsWorkspace/${id}`);
    }

    componentDidMount(){
        AnalyticsWorkspaceService.getAnalyticsWorkspaces().then((res) => {
            this.setState({ analyticsWorkspaces: res.data});
        });
    }

    addAnalyticsWorkspace(){
        this.props.history.push('/add-analyticsWorkspace/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AnalyticsWorkspace List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAnalyticsWorkspace}> Add AnalyticsWorkspace</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> BusinessDomain </th>
                                    <th> OwnerTeam </th>
                                    <th> GovernanceTier </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.analyticsWorkspaces.map(
                                        analyticsWorkspace => 
                                        <tr key = {analyticsWorkspace.analyticsWorkspaceId}>
                                             <td> { analyticsWorkspace.name } </td>
                                             <td> { analyticsWorkspace.businessDomain } </td>
                                             <td> { analyticsWorkspace.ownerTeam } </td>
                                             <td> { analyticsWorkspace.governanceTier } </td>
                                             <td>
                                                 <button onClick={ () => this.editAnalyticsWorkspace(analyticsWorkspace.analyticsWorkspaceId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAnalyticsWorkspace(analyticsWorkspace.analyticsWorkspaceId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAnalyticsWorkspace(analyticsWorkspace.analyticsWorkspaceId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAnalyticsWorkspaceComponent
