import React, { Component } from 'react'
import SubscriberService from '../services/SubscriberService';

class CreateSubscriberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                address: '',
                channel: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeChannelHandler = this.changeChannelHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SubscriberService.getSubscriberById(this.state.id).then( (res) =>{
                let subscriber = res.data;
                this.setState({
                    name: subscriber.name,
                    address: subscriber.address,
                    channel: subscriber.channel
                });
            });
        }        
    }
    saveOrUpdateSubscriber = (e) => {
        e.preventDefault();
        let subscriber = {
                subscriberId: this.state.id,
                name: this.state.name,
                address: this.state.address,
                channel: this.state.channel
            };
        console.log('subscriber => ' + JSON.stringify(subscriber));

        // step 5
        if(this.state.id === '_add'){
            subscriber.subscriberId=''
            SubscriberService.createSubscriber(subscriber).then(res =>{
                this.props.history.push('/subscribers');
            });
        }else{
            SubscriberService.updateSubscriber(subscriber).then( res => {
                this.props.history.push('/subscribers');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeChannelHandler= (event) => {
        this.setState({channel: event.target.value});
    }

    cancel(){
        this.props.history.push('/subscribers');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Subscriber</h3>
        }else{
            return <h3 className="text-center">Update Subscriber</h3>
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

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> Channel:&emsp; </label>
                                                <select value={this.state.channel} onChange={this.changeChannelHandler}>
                      <option name="Channel" className="form-control" >
                          Email
                      </option>
                      <option name="Channel" className="form-control" >
                          SMS
                      </option>
                      <option name="Channel" className="form-control" >
                          Webhook
                      </option>
                      <option name="Channel" className="form-control" >
                          Chat
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSubscriber}>Save</button>
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

export default CreateSubscriberComponent
