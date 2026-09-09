import React, { Component } from 'react'
import DealService from '../services/DealService';

class CreateDealComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                floorPrice: '',
                dealType: ''
        }
        this.changefloorPriceHandler = this.changefloorPriceHandler.bind(this);
        this.changeDealTypeHandler = this.changeDealTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DealService.getDealById(this.state.id).then( (res) =>{
                let deal = res.data;
                this.setState({
                    floorPrice: deal.floorPrice,
                    dealType: deal.dealType
                });
            });
        }        
    }
    saveOrUpdateDeal = (e) => {
        e.preventDefault();
        let deal = {
                dealId: this.state.id,
                floorPrice: this.state.floorPrice,
                dealType: this.state.dealType
            };
        console.log('deal => ' + JSON.stringify(deal));

        // step 5
        if(this.state.id === '_add'){
            deal.dealId=''
            DealService.createDeal(deal).then(res =>{
                this.props.history.push('/deals');
            });
        }else{
            DealService.updateDeal(deal).then( res => {
                this.props.history.push('/deals');
            });
        }
    }
    
    changefloorPriceHandler= (event) => {
        this.setState({floorPrice: event.target.value});
    }
    changeDealTypeHandler= (event) => {
        this.setState({dealType: event.target.value});
    }

    cancel(){
        this.props.history.push('/deals');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Deal</h3>
        }else{
            return <h3 className="text-center">Update Deal</h3>
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
                                            <label> floorPrice:&emsp; </label>
                                                <input placeholder="floorPrice" name="floorPrice" className="form-control" value={this.state.floorPrice} onChange={this.changefloorPriceHandler}/>

                                            <label> DealType:&emsp; </label>
                                                <select value={this.state.dealType} onChange={this.changeDealTypeHandler}>
                      <option name="DealType" className="form-control" >
                          OpenAuction
                      </option>
                      <option name="DealType" className="form-control" >
                          PrivateAuction
                      </option>
                      <option name="DealType" className="form-control" >
                          PreferredDeal
                      </option>
                      <option name="DealType" className="form-control" >
                          ProgrammaticGuaranteed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDeal}>Save</button>
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

export default CreateDealComponent
