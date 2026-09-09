import React, { Component } from 'react'
import ConversionEventService from '../services/ConversionEventService';

class CreateConversionEventComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                timestamp: '',
                value: '',
                eventType: '',
                attributionModel: ''
        }
        this.changetimestampHandler = this.changetimestampHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
        this.changeEventTypeHandler = this.changeEventTypeHandler.bind(this);
        this.changeAttributionModelHandler = this.changeAttributionModelHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ConversionEventService.getConversionEventById(this.state.id).then( (res) =>{
                let conversionEvent = res.data;
                this.setState({
                    timestamp: conversionEvent.timestamp,
                    value: conversionEvent.value,
                    eventType: conversionEvent.eventType,
                    attributionModel: conversionEvent.attributionModel
                });
            });
        }        
    }
    saveOrUpdateConversionEvent = (e) => {
        e.preventDefault();
        let conversionEvent = {
                conversionEventId: this.state.id,
                timestamp: this.state.timestamp,
                value: this.state.value,
                eventType: this.state.eventType,
                attributionModel: this.state.attributionModel
            };
        console.log('conversionEvent => ' + JSON.stringify(conversionEvent));

        // step 5
        if(this.state.id === '_add'){
            conversionEvent.conversionEventId=''
            ConversionEventService.createConversionEvent(conversionEvent).then(res =>{
                this.props.history.push('/conversionEvents');
            });
        }else{
            ConversionEventService.updateConversionEvent(conversionEvent).then( res => {
                this.props.history.push('/conversionEvents');
            });
        }
    }
    
    changetimestampHandler= (event) => {
        this.setState({timestamp: event.target.value});
    }
    changevalueHandler= (event) => {
        this.setState({value: event.target.value});
    }
    changeEventTypeHandler= (event) => {
        this.setState({eventType: event.target.value});
    }
    changeAttributionModelHandler= (event) => {
        this.setState({attributionModel: event.target.value});
    }

    cancel(){
        this.props.history.push('/conversionEvents');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ConversionEvent</h3>
        }else{
            return <h3 className="text-center">Update ConversionEvent</h3>
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
                                            <label> timestamp:&emsp; </label>
                                                <input type="time" placeholder="timestamp" name="timestamp" className="form-control" value={this.state.timestamp} onChange={this.changetimestampHandler}/>

                                            <label> value:&emsp; </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                            <label> EventType:&emsp; </label>
                                                <select value={this.state.eventType} onChange={this.changeEventTypeHandler}>
                      <option name="EventType" className="form-control" >
                          Lead
                      </option>
                      <option name="EventType" className="form-control" >
                          Purchase
                      </option>
                      <option name="EventType" className="form-control" >
                          Signup
                      </option>
                      <option name="EventType" className="form-control" >
                          AddToCart
                      </option>
                      <option name="EventType" className="form-control" >
                          ViewContent
                      </option>
                      <option name="EventType" className="form-control" >
                          AppInstall
                      </option>
                    </select>

                                            <label> AttributionModel:&emsp; </label>
                                                <select value={this.state.attributionModel} onChange={this.changeAttributionModelHandler}>
                      <option name="AttributionModel" className="form-control" >
                          LastClick
                      </option>
                      <option name="AttributionModel" className="form-control" >
                          FirstTouch
                      </option>
                      <option name="AttributionModel" className="form-control" >
                          Linear
                      </option>
                      <option name="AttributionModel" className="form-control" >
                          TimeDecay
                      </option>
                      <option name="AttributionModel" className="form-control" >
                          PositionBased
                      </option>
                      <option name="AttributionModel" className="form-control" >
                          DataDriven
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateConversionEvent}>Save</button>
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

export default CreateConversionEventComponent
