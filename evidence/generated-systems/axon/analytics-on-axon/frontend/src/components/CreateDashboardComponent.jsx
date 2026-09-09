import React, { Component } from 'react'
import DashboardService from '../services/DashboardService';

class CreateDashboardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                theme: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changethemeHandler = this.changethemeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DashboardService.getDashboardById(this.state.id).then( (res) =>{
                let dashboard = res.data;
                this.setState({
                    title: dashboard.title,
                    theme: dashboard.theme,
                    status: dashboard.status
                });
            });
        }        
    }
    saveOrUpdateDashboard = (e) => {
        e.preventDefault();
        let dashboard = {
                dashboardId: this.state.id,
                title: this.state.title,
                theme: this.state.theme,
                status: this.state.status
            };
        console.log('dashboard => ' + JSON.stringify(dashboard));

        // step 5
        if(this.state.id === '_add'){
            dashboard.dashboardId=''
            DashboardService.createDashboard(dashboard).then(res =>{
                this.props.history.push('/dashboards');
            });
        }else{
            DashboardService.updateDashboard(dashboard).then( res => {
                this.props.history.push('/dashboards');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Dashboard</h3>
        }else{
            return <h3 className="text-center">Update Dashboard</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> theme:&emsp; </label>
                                                <input placeholder="theme" name="theme" className="form-control" value={this.state.theme} onChange={this.changethemeHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDashboard}>Save</button>
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

export default CreateDashboardComponent
