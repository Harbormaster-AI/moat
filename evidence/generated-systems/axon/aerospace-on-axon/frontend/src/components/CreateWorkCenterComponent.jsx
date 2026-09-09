import React, { Component } from 'react'
import WorkCenterService from '../services/WorkCenterService';

class CreateWorkCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                capability: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecapabilityHandler = this.changecapabilityHandler.bind(this);
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
                    capability: workCenter.capability
                });
            });
        }        
    }
    saveOrUpdateWorkCenter = (e) => {
        e.preventDefault();
        let workCenter = {
                workCenterId: this.state.id,
                name: this.state.name,
                capability: this.state.capability
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
    changecapabilityHandler= (event) => {
        this.setState({capability: event.target.value});
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

                                            <label> capability:&emsp; </label>
                                                <input placeholder="capability" name="capability" className="form-control" value={this.state.capability} onChange={this.changecapabilityHandler}/>

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
