import React, { Component } from 'react'
import GiftCardService from '../services/GiftCardService';

class CreateGiftCardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                balance: '',
                expirationDate: '',
                status: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changebalanceHandler = this.changebalanceHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            GiftCardService.getGiftCardById(this.state.id).then( (res) =>{
                let giftCard = res.data;
                this.setState({
                    code: giftCard.code,
                    balance: giftCard.balance,
                    expirationDate: giftCard.expirationDate,
                    status: giftCard.status
                });
            });
        }        
    }
    saveOrUpdateGiftCard = (e) => {
        e.preventDefault();
        let giftCard = {
                giftCardId: this.state.id,
                code: this.state.code,
                balance: this.state.balance,
                expirationDate: this.state.expirationDate,
                status: this.state.status
            };
        console.log('giftCard => ' + JSON.stringify(giftCard));

        // step 5
        if(this.state.id === '_add'){
            giftCard.giftCardId=''
            GiftCardService.createGiftCard(giftCard).then(res =>{
                this.props.history.push('/giftCards');
            });
        }else{
            GiftCardService.updateGiftCard(giftCard).then( res => {
                this.props.history.push('/giftCards');
            });
        }
    }
    
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changebalanceHandler= (event) => {
        this.setState({balance: event.target.value});
    }
    changeexpirationDateHandler= (event) => {
        this.setState({expirationDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/giftCards');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add GiftCard</h3>
        }else{
            return <h3 className="text-center">Update GiftCard</h3>
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
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> balance:&emsp; </label>
                                                <input placeholder="balance" name="balance" className="form-control" value={this.state.balance} onChange={this.changebalanceHandler}/>

                                            <label> expirationDate:&emsp; </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Redeemed
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                      <option name="Status" className="form-control" >
                          Disabled
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateGiftCard}>Save</button>
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

export default CreateGiftCardComponent
