import React, { Component } from 'react'
import FulfillmentCenterService from '../services/FulfillmentCenterService';

class CreateFulfillmentCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                centerCode: '',
                address: '',
                timezone: '',
                asActive: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecenterCodeHandler = this.changecenterCodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changetimezoneHandler = this.changetimezoneHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            FulfillmentCenterService.getFulfillmentCenterById(this.state.id).then( (res) =>{
                let fulfillmentCenter = res.data;
                this.setState({
                    name: fulfillmentCenter.name,
                    centerCode: fulfillmentCenter.centerCode,
                    address: fulfillmentCenter.address,
                    timezone: fulfillmentCenter.timezone,
                    asActive: fulfillmentCenter.asActive
                });
            });
        }        
    }
    saveOrUpdateFulfillmentCenter = (e) => {
        e.preventDefault();
        let fulfillmentCenter = {
                fulfillmentCenterId: this.state.id,
                name: this.state.name,
                centerCode: this.state.centerCode,
                address: this.state.address,
                timezone: this.state.timezone,
                asActive: this.state.asActive
            };
        console.log('fulfillmentCenter => ' + JSON.stringify(fulfillmentCenter));

        // step 5
        if(this.state.id === '_add'){
            fulfillmentCenter.fulfillmentCenterId=''
            FulfillmentCenterService.createFulfillmentCenter(fulfillmentCenter).then(res =>{
                this.props.history.push('/fulfillmentCenters');
            });
        }else{
            FulfillmentCenterService.updateFulfillmentCenter(fulfillmentCenter).then( res => {
                this.props.history.push('/fulfillmentCenters');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecenterCodeHandler= (event) => {
        this.setState({centerCode: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changetimezoneHandler= (event) => {
        this.setState({timezone: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }

    cancel(){
        this.props.history.push('/fulfillmentCenters');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add FulfillmentCenter</h3>
        }else{
            return <h3 className="text-center">Update FulfillmentCenter</h3>
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

                                            <label> centerCode:&emsp; </label>
                                                <input placeholder="centerCode" name="centerCode" className="form-control" value={this.state.centerCode} onChange={this.changecenterCodeHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> timezone:&emsp; </label>
                                                <input placeholder="timezone" name="timezone" className="form-control" value={this.state.timezone} onChange={this.changetimezoneHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateFulfillmentCenter}>Save</button>
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

export default CreateFulfillmentCenterComponent
