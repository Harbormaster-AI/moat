import React, { Component } from 'react'
import RoutingService from '../services/RoutingService';

class CreateRoutingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                routingNumber: '',
                revision: '',
                effectivityStart: '',
                effectivityEnd: '',
                routingType: '',
                status: ''
        }
        this.changeroutingNumberHandler = this.changeroutingNumberHandler.bind(this);
        this.changerevisionHandler = this.changerevisionHandler.bind(this);
        this.changeeffectivityStartHandler = this.changeeffectivityStartHandler.bind(this);
        this.changeeffectivityEndHandler = this.changeeffectivityEndHandler.bind(this);
        this.changeRoutingTypeHandler = this.changeRoutingTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RoutingService.getRoutingById(this.state.id).then( (res) =>{
                let routing = res.data;
                this.setState({
                    routingNumber: routing.routingNumber,
                    revision: routing.revision,
                    effectivityStart: routing.effectivityStart,
                    effectivityEnd: routing.effectivityEnd,
                    routingType: routing.routingType,
                    status: routing.status
                });
            });
        }        
    }
    saveOrUpdateRouting = (e) => {
        e.preventDefault();
        let routing = {
                routingId: this.state.id,
                routingNumber: this.state.routingNumber,
                revision: this.state.revision,
                effectivityStart: this.state.effectivityStart,
                effectivityEnd: this.state.effectivityEnd,
                routingType: this.state.routingType,
                status: this.state.status
            };
        console.log('routing => ' + JSON.stringify(routing));

        // step 5
        if(this.state.id === '_add'){
            routing.routingId=''
            RoutingService.createRouting(routing).then(res =>{
                this.props.history.push('/routings');
            });
        }else{
            RoutingService.updateRouting(routing).then( res => {
                this.props.history.push('/routings');
            });
        }
    }
    
    changeroutingNumberHandler= (event) => {
        this.setState({routingNumber: event.target.value});
    }
    changerevisionHandler= (event) => {
        this.setState({revision: event.target.value});
    }
    changeeffectivityStartHandler= (event) => {
        this.setState({effectivityStart: event.target.value});
    }
    changeeffectivityEndHandler= (event) => {
        this.setState({effectivityEnd: event.target.value});
    }
    changeRoutingTypeHandler= (event) => {
        this.setState({routingType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/routings');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Routing</h3>
        }else{
            return <h3 className="text-center">Update Routing</h3>
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
                                            <label> routingNumber:&emsp; </label>
                                                <input placeholder="routingNumber" name="routingNumber" className="form-control" value={this.state.routingNumber} onChange={this.changeroutingNumberHandler}/>

                                            <label> revision:&emsp; </label>
                                                <input placeholder="revision" name="revision" className="form-control" value={this.state.revision} onChange={this.changerevisionHandler}/>

                                            <label> effectivityStart:&emsp; </label>
                                                <input type="date" placeholder="effectivityStart" name="effectivityStart" className="form-control" value={this.state.effectivityStart} onChange={this.changeeffectivityStartHandler}/>

                                            <label> effectivityEnd:&emsp; </label>
                                                <input type="date" placeholder="effectivityEnd" name="effectivityEnd" className="form-control" value={this.state.effectivityEnd} onChange={this.changeeffectivityEndHandler}/>

                                            <label> RoutingType:&emsp; </label>
                                                <select value={this.state.routingType} onChange={this.changeRoutingTypeHandler}>
                      <option name="RoutingType" className="form-control" >
                          Standard
                      </option>
                      <option name="RoutingType" className="form-control" >
                          Alternate
                      </option>
                      <option name="RoutingType" className="form-control" >
                          Rework
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          Obsolete
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRouting}>Save</button>
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

export default CreateRoutingComponent
