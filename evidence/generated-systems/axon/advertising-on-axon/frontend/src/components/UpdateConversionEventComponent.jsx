import React, { Component } from 'react'
import ConversionEventService from '../services/ConversionEventService';

class UpdateConversionEventComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                timestamp: '',
                value: '',
                eventType: '',
                attributionModel: ''
        }
        this.updateConversionEvent = this.updateConversionEvent.bind(this);

        this.changetimestampHandler = this.changetimestampHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
        this.changeEventTypeHandler = this.changeEventTypeHandler.bind(this);
        this.changeAttributionModelHandler = this.changeAttributionModelHandler.bind(this);
    }

    componentDidMount(){
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

    updateConversionEvent = (e) => {
        e.preventDefault();
        let conversionEvent = {
            conversionEventId: this.state.id,
            timestamp: this.state.timestamp,
            value: this.state.value,
            eventType: this.state.eventType,
            attributionModel: this.state.attributionModel
        };
        console.log('conversionEvent => ' + JSON.stringify(conversionEvent));
        console.log('id => ' + JSON.stringify(this.state.id));
        ConversionEventService.updateConversionEvent(conversionEvent).then( res => {
            this.props.history.push('/conversionEvents');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ConversionEvent</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> timestamp: </label>
                                                <input type="time" placeholder="timestamp" name="timestamp" className="form-control" value={this.state.timestamp} onChange={this.changetimestampHandler}/>

                                            <label> value: </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                            <label> EventType: </label>
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

                                            <label> AttributionModel: </label>
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
                                        <button className="btn btn-success" onClick={this.updateConversionEvent}>Save</button>
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

export default UpdateConversionEventComponent
