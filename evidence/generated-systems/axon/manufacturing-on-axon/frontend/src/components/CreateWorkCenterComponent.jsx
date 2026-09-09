import React, { Component } from 'react'
import WorkCenterService from '../services/WorkCenterService';

class CreateWorkCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                code: '',
                capacityPerHour: '',
                oeeTarget: '',
                workCenterType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changecapacityPerHourHandler = this.changecapacityPerHourHandler.bind(this);
        this.changeoeeTargetHandler = this.changeoeeTargetHandler.bind(this);
        this.changeWorkCenterTypeHandler = this.changeWorkCenterTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            WorkCenterService.getWorkCenterById(this.state.id).then( (res) =>{
                let workCenter = res.data;
                this.setState({
                    name: workCenter.name,
                    code: workCenter.code,
                    capacityPerHour: workCenter.capacityPerHour,
                    oeeTarget: workCenter.oeeTarget,
                    workCenterType: workCenter.workCenterType
                });
            });
        }        
    }
    saveOrUpdateWorkCenter = (e) => {
        e.preventDefault();
        let workCenter = {
                workCenterId: this.state.id,
                name: this.state.name,
                code: this.state.code,
                capacityPerHour: this.state.capacityPerHour,
                oeeTarget: this.state.oeeTarget,
                workCenterType: this.state.workCenterType
            };
        console.log('workCenter => ' + JSON.stringify(workCenter));

        // step 5
        if(this.state.id === '_add'){
            workCenter.workCenterId=''
            WorkCenterService.createWorkCenter(workCenter).then(res =>{
                this.props.history.push('/workCenters');
            });
        }else{
            WorkCenterService.updateWorkCenter(workCenter).then( res => {
                this.props.history.push('/workCenters');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changecapacityPerHourHandler= (event) => {
        this.setState({capacityPerHour: event.target.value});
    }
    changeoeeTargetHandler= (event) => {
        this.setState({oeeTarget: event.target.value});
    }
    changeWorkCenterTypeHandler= (event) => {
        this.setState({workCenterType: event.target.value});
    }

    cancel(){
        this.props.history.push('/workCenters');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add WorkCenter</h3>
        }else{
            return <h3 className="text-center">Update WorkCenter</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> capacityPerHour:&emsp; </label>
                                                <input type="number" placeholder="capacityPerHour" name="capacityPerHour" className="form-control" value={this.state.capacityPerHour} onChange={this.changecapacityPerHourHandler}/>

                                            <label> oeeTarget:&emsp; </label>
                                                <input placeholder="oeeTarget" name="oeeTarget" className="form-control" value={this.state.oeeTarget} onChange={this.changeoeeTargetHandler}/>

                                            <label> WorkCenterType:&emsp; </label>
                                                <select value={this.state.workCenterType} onChange={this.changeWorkCenterTypeHandler}>
                      <option name="WorkCenterType" className="form-control" >
                          Machining
                      </option>
                      <option name="WorkCenterType" className="form-control" >
                          Assembly
                      </option>
                      <option name="WorkCenterType" className="form-control" >
                          Painting
                      </option>
                      <option name="WorkCenterType" className="form-control" >
                          Packaging
                      </option>
                      <option name="WorkCenterType" className="form-control" >
                          Test
                      </option>
                      <option name="WorkCenterType" className="form-control" >
                          Warehouse
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWorkCenter}>Save</button>
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

export default CreateWorkCenterComponent
