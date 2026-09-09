import React, { Component } from 'react'
import RoutingService from '../services/RoutingService';

class UpdateRoutingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                routingNumber: '',
                revision: '',
                effectivityStart: '',
                effectivityEnd: '',
                routingType: '',
                status: ''
        }
        this.updateRouting = this.updateRouting.bind(this);

        this.changeroutingNumberHandler = this.changeroutingNumberHandler.bind(this);
        this.changerevisionHandler = this.changerevisionHandler.bind(this);
        this.changeeffectivityStartHandler = this.changeeffectivityStartHandler.bind(this);
        this.changeeffectivityEndHandler = this.changeeffectivityEndHandler.bind(this);
        this.changeRoutingTypeHandler = this.changeRoutingTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateRouting = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        RoutingService.updateRouting(routing).then( res => {
            this.props.history.push('/routings');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Routing</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> routingNumber: </label>
                                                <input placeholder="routingNumber" name="routingNumber" className="form-control" value={this.state.routingNumber} onChange={this.changeroutingNumberHandler}/>

                                            <label> revision: </label>
                                                <input placeholder="revision" name="revision" className="form-control" value={this.state.revision} onChange={this.changerevisionHandler}/>

                                            <label> effectivityStart: </label>
                                                <input type="date" placeholder="effectivityStart" name="effectivityStart" className="form-control" value={this.state.effectivityStart} onChange={this.changeeffectivityStartHandler}/>

                                            <label> effectivityEnd: </label>
                                                <input type="date" placeholder="effectivityEnd" name="effectivityEnd" className="form-control" value={this.state.effectivityEnd} onChange={this.changeeffectivityEndHandler}/>

                                            <label> RoutingType: </label>
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

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateRouting}>Save</button>
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

export default UpdateRoutingComponent
