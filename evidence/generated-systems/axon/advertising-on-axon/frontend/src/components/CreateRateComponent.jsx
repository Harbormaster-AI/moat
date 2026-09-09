import React, { Component } from 'react'
import RateService from '../services/RateService';

class CreateRateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                unitPrice: '',
                adFormat: '',
                pricingModel: ''
        }
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changeAdFormatHandler = this.changeAdFormatHandler.bind(this);
        this.changePricingModelHandler = this.changePricingModelHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RateService.getRateById(this.state.id).then( (res) =>{
                let rate = res.data;
                this.setState({
                    unitPrice: rate.unitPrice,
                    adFormat: rate.adFormat,
                    pricingModel: rate.pricingModel
                });
            });
        }        
    }
    saveOrUpdateRate = (e) => {
        e.preventDefault();
        let rate = {
                rateId: this.state.id,
                unitPrice: this.state.unitPrice,
                adFormat: this.state.adFormat,
                pricingModel: this.state.pricingModel
            };
        console.log('rate => ' + JSON.stringify(rate));

        // step 5
        if(this.state.id === '_add'){
            rate.rateId=''
            RateService.createRate(rate).then(res =>{
                this.props.history.push('/rates');
            });
        }else{
            RateService.updateRate(rate).then( res => {
                this.props.history.push('/rates');
            });
        }
    }
    
    changeunitPriceHandler= (event) => {
        this.setState({unitPrice: event.target.value});
    }
    changeAdFormatHandler= (event) => {
        this.setState({adFormat: event.target.value});
    }
    changePricingModelHandler= (event) => {
        this.setState({pricingModel: event.target.value});
    }

    cancel(){
        this.props.history.push('/rates');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Rate</h3>
        }else{
            return <h3 className="text-center">Update Rate</h3>
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
                                            <label> unitPrice:&emsp; </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> AdFormat:&emsp; </label>
                                                <select value={this.state.adFormat} onChange={this.changeAdFormatHandler}>
                      <option name="AdFormat" className="form-control" >
                          Banner
                      </option>
                      <option name="AdFormat" className="form-control" >
                          Video
                      </option>
                      <option name="AdFormat" className="form-control" >
                          Native
                      </option>
                      <option name="AdFormat" className="form-control" >
                          Audio
                      </option>
                      <option name="AdFormat" className="form-control" >
                          Interstitial
                      </option>
                      <option name="AdFormat" className="form-control" >
                          RichMedia
                      </option>
                      <option name="AdFormat" className="form-control" >
                          SearchText
                      </option>
                      <option name="AdFormat" className="form-control" >
                          SocialPost
                      </option>
                      <option name="AdFormat" className="form-control" >
                          CTVVideo
                      </option>
                    </select>

                                            <label> PricingModel:&emsp; </label>
                                                <select value={this.state.pricingModel} onChange={this.changePricingModelHandler}>
                      <option name="PricingModel" className="form-control" >
                          CPM
                      </option>
                      <option name="PricingModel" className="form-control" >
                          CPC
                      </option>
                      <option name="PricingModel" className="form-control" >
                          CPA
                      </option>
                      <option name="PricingModel" className="form-control" >
                          CPL
                      </option>
                      <option name="PricingModel" className="form-control" >
                          CPV
                      </option>
                      <option name="PricingModel" className="form-control" >
                          FlatFee
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRate}>Save</button>
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

export default CreateRateComponent
