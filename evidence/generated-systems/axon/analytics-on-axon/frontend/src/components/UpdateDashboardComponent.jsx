import React, { Component } from 'react'
import DashboardService from '../services/DashboardService';

class UpdateDashboardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                theme: '',
                status: ''
        }
        this.updateDashboard = this.updateDashboard.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changethemeHandler = this.changethemeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        DashboardService.getDashboardById(this.state.id).then( (res) =>{
            let dashboard = res.data;
            this.setState({
                title: dashboard.title,
                theme: dashboard.theme,
                status: dashboard.status
            });
        });
    }

    updateDashboard = (e) => {
        e.preventDefault();
        let dashboard = {
            dashboardId: this.state.id,
            title: this.state.title,
            theme: this.state.theme,
            status: this.state.status
        };
        console.log('dashboard => ' + JSON.stringify(dashboard));
        console.log('id => ' + JSON.stringify(this.state.id));
        DashboardService.updateDashboard(dashboard).then( res => {
            this.props.history.push('/dashboards');
        });
    }

    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changethemeHandler= (event) => {
        this.setState({theme: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/dashboards');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Dashboard</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> theme: </label>
                                                <input placeholder="theme" name="theme" className="form-control" value={this.state.theme} onChange={this.changethemeHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Live
                      </option>
                      <option name="Status" className="form-control" >
                          Archived
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDashboard}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateDashboardComponent
