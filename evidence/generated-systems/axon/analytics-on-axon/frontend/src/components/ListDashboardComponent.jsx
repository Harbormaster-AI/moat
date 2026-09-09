import React, { Component } from 'react'
import DashboardService from '../services/DashboardService'

class ListDashboardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dashboards: []
        }
        this.addDashboard = this.addDashboard.bind(this);
        this.editDashboard = this.editDashboard.bind(this);
        this.deleteDashboard = this.deleteDashboard.bind(this);
    }

    deleteDashboard(id){
        DashboardService.deleteDashboard(id).then( res => {
            this.setState({dashboards: this.state.dashboards.filter(dashboard => dashboard.dashboardId !== id)});
        });
    }
    viewDashboard(id){
        this.props.history.push(`/view-dashboard/${id}`);
    }
    editDashboard(id){
        this.props.history.push(`/add-dashboard/${id}`);
    }

    componentDidMount(){
        DashboardService.getDashboards().then((res) => {
            this.setState({ dashboards: res.data});
        });
    }

    addDashboard(){
        this.props.history.push('/add-dashboard/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Dashboard List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDashboard}> Add Dashboard</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> Theme </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dashboards.map(
                                        dashboard => 
                                        <tr key = {dashboard.dashboardId}>
                                             <td> { dashboard.title } </td>
                                             <td> { dashboard.theme } </td>
                                             <td> { dashboard.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editDashboard(dashboard.dashboardId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDashboard(dashboard.dashboardId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDashboard(dashboard.dashboardId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDashboardComponent
