import React, { Component } from 'react'
import InferenceEndpointService from '../services/InferenceEndpointService';

class CreateInferenceEndpointComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                endpointUrl: '',
                trafficShare: '',
                mode: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeendpointUrlHandler = this.changeendpointUrlHandler.bind(this);
        this.changetrafficShareHandler = this.changetrafficShareHandler.bind(this);
        this.changeModeHandler = this.changeModeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateInferenceEndpoint = (e) => {
        e.preventDefault();
        let inferenceEndpoint = {
                inferenceEndpointId: this.state.id,
                name: this.state.name,
                endpointUrl: this.state.endpointUrl,
                trafficShare: this.state.trafficShare,
                mode: this.state.mode
            };
        console.log('inferenceEndpoint => ' + JSON.stringify(inferenceEndpoint));

        // step 5
        if(this.state.id === '_add'){
            inferenceEndpoint.inferenceEndpointId=''
            InferenceEndpointService.createInferenceEndpoint(inferenceEndpoint).then(res =>{
                this.props.history.push('/inferenceEndpoints');
            });
        }else{
            InferenceEndpointService.updateInferenceEndpoint(inferenceEndpoint).then( res => {
                this.props.history.push('/inferenceEndpoints');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InferenceEndpoint</h3>
        }else{
            return <h3 className="text-center">Update InferenceEndpoint</h3>
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

                                            <label> endpointUrl:&emsp; </label>
                                                <input placeholder="endpointUrl" name="endpointUrl" className="form-control" value={this.state.endpointUrl} onChange={this.changeendpointUrlHandler}/>

                                            <label> trafficShare:&emsp; </label>
                                                <input placeholder="trafficShare" name="trafficShare" className="form-control" value={this.state.trafficShare} onChange={this.changetrafficShareHandler}/>

                                            <label> Mode:&emsp; </label>
                                                <select value={this.state.mode} onChange={this.changeModeHandler}>
                      <option name="Mode" className="form-control" >
                          Batch
                      </option>
                      <option name="Mode" className="form-control" >
                          RealTime
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInferenceEndpoint}>Save</button>
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

export default CreateInferenceEndpointComponent
