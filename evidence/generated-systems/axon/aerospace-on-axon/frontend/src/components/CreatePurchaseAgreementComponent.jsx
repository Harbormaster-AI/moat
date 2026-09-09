import React, { Component } from 'react'
import PurchaseAgreementService from '../services/PurchaseAgreementService';

class CreatePurchaseAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                agreementNumber: '',
                effectiveDate: ''
        }
        this.changeagreementNumberHandler = this.changeagreementNumberHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PurchaseAgreementService.getPurchaseAgreementById(this.state.id).then( (res) =>{
                let purchaseAgreement = res.data;
                this.setState({
                    agreementNumber: purchaseAgreement.agreementNumber,
                    effectiveDate: purchaseAgreement.effectiveDate
                });
            });
        }        
    }
    saveOrUpdatePurchaseAgreement = (e) => {
        e.preventDefault();
        let purchaseAgreement = {
                purchaseAgreementId: this.state.id,
                agreementNumber: this.state.agreementNumber,
                effectiveDate: this.state.effectiveDate
            };
        console.log('purchaseAgreement => ' + JSON.stringify(purchaseAgreement));

        // step 5
        if(this.state.id === '_add'){
            purchaseAgreement.purchaseAgreementId=''
            PurchaseAgreementService.createPurchaseAgreement(purchaseAgreement).then(res =>{
                this.props.history.push('/purchaseAgreements');
            });
        }else{
            PurchaseAgreementService.updatePurchaseAgreement(purchaseAgreement).then( res => {
                this.props.history.push('/purchaseAgreements');
            });
        }
    }
    
    changeagreementNumberHandler= (event) => {
        this.setState({agreementNumber: event.target.value});
    }
    changeeffectiveDateHandler= (event) => {
        this.setState({effectiveDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/purchaseAgreements');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PurchaseAgreement</h3>
        }else{
            return <h3 className="text-center">Update PurchaseAgreement</h3>
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
                                            <label> agreementNumber:&emsp; </label>
                                                <input placeholder="agreementNumber" name="agreementNumber" className="form-control" value={this.state.agreementNumber} onChange={this.changeagreementNumberHandler}/>

                                            <label> effectiveDate:&emsp; </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePurchaseAgreement}>Save</button>
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

export default CreatePurchaseAgreementComponent
