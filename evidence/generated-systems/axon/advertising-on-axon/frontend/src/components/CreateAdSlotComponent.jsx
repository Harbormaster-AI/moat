import React, { Component } from 'react'
import AdSlotService from '../services/AdSlotService';

class CreateAdSlotComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                slotCode: '',
                width: '',
                height: '',
                floorPrice: '',
                format: ''
        }
        this.changeslotCodeHandler = this.changeslotCodeHandler.bind(this);
        this.changewidthHandler = this.changewidthHandler.bind(this);
        this.changeheightHandler = this.changeheightHandler.bind(this);
        this.changefloorPriceHandler = this.changefloorPriceHandler.bind(this);
        this.changeFormatHandler = this.changeFormatHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AdSlotService.getAdSlotById(this.state.id).then( (res) =>{
                let adSlot = res.data;
                this.setState({
                    slotCode: adSlot.slotCode,
                    width: adSlot.width,
                    height: adSlot.height,
                    floorPrice: adSlot.floorPrice,
                    format: adSlot.format
                });
            });
        }        
    }
    saveOrUpdateAdSlot = (e) => {
        e.preventDefault();
        let adSlot = {
                adSlotId: this.state.id,
                slotCode: this.state.slotCode,
                width: this.state.width,
                height: this.state.height,
                floorPrice: this.state.floorPrice,
                format: this.state.format
            };
        console.log('adSlot => ' + JSON.stringify(adSlot));

        // step 5
        if(this.state.id === '_add'){
            adSlot.adSlotId=''
            AdSlotService.createAdSlot(adSlot).then(res =>{
                this.props.history.push('/adSlots');
            });
        }else{
            AdSlotService.updateAdSlot(adSlot).then( res => {
                this.props.history.push('/adSlots');
            });
        }
    }
    
    changeslotCodeHandler= (event) => {
        this.setState({slotCode: event.target.value});
    }
    changewidthHandler= (event) => {
        this.setState({width: event.target.value});
    }
    changeheightHandler= (event) => {
        this.setState({height: event.target.value});
    }
    changefloorPriceHandler= (event) => {
        this.setState({floorPrice: event.target.value});
    }
    changeFormatHandler= (event) => {
        this.setState({format: event.target.value});
    }

    cancel(){
        this.props.history.push('/adSlots');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AdSlot</h3>
        }else{
            return <h3 className="text-center">Update AdSlot</h3>
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
                                            <label> slotCode:&emsp; </label>
                                                <input placeholder="slotCode" name="slotCode" className="form-control" value={this.state.slotCode} onChange={this.changeslotCodeHandler}/>

                                            <label> width:&emsp; </label>
                                                <input type="number" placeholder="width" name="width" className="form-control" value={this.state.width} onChange={this.changewidthHandler}/>

                                            <label> height:&emsp; </label>
                                                <input type="number" placeholder="height" name="height" className="form-control" value={this.state.height} onChange={this.changeheightHandler}/>

                                            <label> floorPrice:&emsp; </label>
                                                <input placeholder="floorPrice" name="floorPrice" className="form-control" value={this.state.floorPrice} onChange={this.changefloorPriceHandler}/>

                                            <label> Format:&emsp; </label>
                                                <select value={this.state.format} onChange={this.changeFormatHandler}>
                      <option name="Format" className="form-control" >
                          Banner
                      </option>
                      <option name="Format" className="form-control" >
                          Video
                      </option>
                      <option name="Format" className="form-control" >
                          Native
                      </option>
                      <option name="Format" className="form-control" >
                          Audio
                      </option>
                      <option name="Format" className="form-control" >
                          Interstitial
                      </option>
                      <option name="Format" className="form-control" >
                          RichMedia
                      </option>
                      <option name="Format" className="form-control" >
                          SearchText
                      </option>
                      <option name="Format" className="form-control" >
                          SocialPost
                      </option>
                      <option name="Format" className="form-control" >
                          CTVVideo
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAdSlot}>Save</button>
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

export default CreateAdSlotComponent
