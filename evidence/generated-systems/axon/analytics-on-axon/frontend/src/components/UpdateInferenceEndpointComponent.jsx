import React, { Component } from 'react'
import InferenceEndpointService from '../services/InferenceEndpointService';

class UpdateInferenceEndpointComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                endpointUrl: '',
                trafficShare: '',
                mode: ''
        }
        this.updateInferenceEndpoint = this.updateInferenceEndpoint.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeendpointUrlHandler = this.changeendpointUrlHandler.bind(this);
        this.changetrafficShareHandler = this.changetrafficShareHandler.bind(this);
        this.changeModeHandler = this.changeModeHandler.bind(this);
    }

    componentDidMount(){
        InferenceEndpointService.getInferenceEndpointById(this.state.id).then( (res) =>{
            let inferenceEndpoint = res.data;
            this.setState({
                name: inferenceEndpoint.name,
                endpointUrl: inferenceEndpoint.endpointUrl,
                trafficShare: inferenceEndpoint.trafficShare,
                mode: inferenceEndpoint.mode
            });
        });
    }

    updateInferenceEndpoint = (e) => {
        e.preventDefault();
        let inferenceEndpoint = {
            inferenceEndpointId: this.state.id,
            name: this.state.name,
            endpointUrl: this.state.endpointUrl,
            trafficShare: this.state.trafficShare,
            mode: this.state.mode
        };
        console.log('inferenceEndpoint => ' + JSON.stringify(inferenceEndpoint));
        console.log('id => ' + JSON.stringify(this.state.id));
        InferenceEndpointService.updateInferenceEndpoint(inferenceEndpoint).then( res => {
            this.props.history.push('/inferenceEndpoints');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeendpointUrlHandler= (event) => {
        this.setState({endpointUrl: event.target.value});
    }
    changetrafficShareHandler= (event) => {
        this.setState({trafficShare: event.target.value});
    }
    changeModeHandler= (event) => {
        this.setState({mode: event.target.value});
    }

    cancel(){
        this.props.history.push('/inferenceEndpoints');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InferenceEndpoint</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> endpointUrl: </label>
                                                <input placeholder="endpointUrl" name="endpointUrl" className="form-control" value={this.state.endpointUrl} onChange={this.changeendpointUrlHandler}/>

                                            <label> trafficShare: </label>
                                                <input placeholder="trafficShare" name="trafficShare" className="form-control" value={this.state.trafficShare} onChange={this.changetrafficShareHandler}/>

                                            <label> Mode: </label>
                                                <select value={this.state.mode} onChange={this.changeModeHandler}>
                      <option name="Mode" className="form-control" >
                          Batch
                      </option>
                      <option name="Mode" className="form-control" >
                          RealTime
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInferenceEndpoint}>Save</button>
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

export default UpdateInferenceEndpointComponent
