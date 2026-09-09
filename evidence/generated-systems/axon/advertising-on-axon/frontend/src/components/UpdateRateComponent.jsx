import React, { Component } from 'react'
import RateService from '../services/RateService';

class UpdateRateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                unitPrice: '',
                adFormat: '',
                pricingModel: ''
        }
        this.updateRate = this.updateRate.bind(this);

        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changeAdFormatHandler = this.changeAdFormatHandler.bind(this);
        this.changePricingModelHandler = this.changePricingModelHandler.bind(this);
    }

    componentDidMount(){
        RateService.getRateById(this.state.id).then( (res) =>{
            let rate = res.data;
            this.setState({
                unitPrice: rate.unitPrice,
                adFormat: rate.adFormat,
                pricingModel: rate.pricingModel
            });
        });
    }

    updateRate = (e) => {
        e.preventDefault();
        let rate = {
            rateId: this.state.id,
            unitPrice: this.state.unitPrice,
            adFormat: this.state.adFormat,
            pricingModel: this.state.pricingModel
        };
        console.log('rate => ' + JSON.stringify(rate));
        console.log('id => ' + JSON.stringify(this.state.id));
        RateService.updateRate(rate).then( res => {
            this.props.history.push('/rates');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Rate</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> unitPrice: </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> AdFormat: </label>
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

                                            <label> PricingModel: </label>
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
                                        <button className="btn btn-success" onClick={this.updateRate}>Save</button>
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

export default UpdateRateComponent
