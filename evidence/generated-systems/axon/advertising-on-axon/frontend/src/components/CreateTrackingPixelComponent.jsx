import React, { Component } from 'react'
import TrackingPixelService from '../services/TrackingPixelService';

class CreateTrackingPixelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                url: '',
                eventType: '',
                pixelType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeurlHandler = this.changeurlHandler.bind(this);
        this.changeEventTypeHandler = this.changeEventTypeHandler.bind(this);
        this.changePixelTypeHandler = this.changePixelTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TrackingPixelService.getTrackingPixelById(this.state.id).then( (res) =>{
                let trackingPixel = res.data;
                this.setState({
                    name: trackingPixel.name,
                    url: trackingPixel.url,
                    eventType: trackingPixel.eventType,
                    pixelType: trackingPixel.pixelType
                });
            });
        }        
    }
    saveOrUpdateTrackingPixel = (e) => {
        e.preventDefault();
        let trackingPixel = {
                trackingPixelId: this.state.id,
                name: this.state.name,
                url: this.state.url,
                eventType: this.state.eventType,
                pixelType: this.state.pixelType
            };
        console.log('trackingPixel => ' + JSON.stringify(trackingPixel));

        // step 5
        if(this.state.id === '_add'){
            trackingPixel.trackingPixelId=''
            TrackingPixelService.createTrackingPixel(trackingPixel).then(res =>{
                this.props.history.push('/trackingPixels');
            });
        }else{
            TrackingPixelService.updateTrackingPixel(trackingPixel).then( res => {
                this.props.history.push('/trackingPixels');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeurlHandler= (event) => {
        this.setState({url: event.target.value});
    }
    changeEventTypeHandler= (event) => {
        this.setState({eventType: event.target.value});
    }
    changePixelTypeHandler= (event) => {
        this.setState({pixelType: event.target.value});
    }

    cancel(){
        this.props.history.push('/trackingPixels');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TrackingPixel</h3>
        }else{
            return <h3 className="text-center">Update TrackingPixel</h3>
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

                                            <label> url:&emsp; </label>
                                                <input placeholder="url" name="url" className="form-control" value={this.state.url} onChange={this.changeurlHandler}/>

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

                                            <label> PixelType:&emsp; </label>
                                                <select value={this.state.pixelType} onChange={this.changePixelTypeHandler}>
                      <option name="PixelType" className="form-control" >
                          Image
                      </option>
                      <option name="PixelType" className="form-control" >
                          JavaScript
                      </option>
                      <option name="PixelType" className="form-control" >
                          ServerSide
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTrackingPixel}>Save</button>
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

export default CreateTrackingPixelComponent
