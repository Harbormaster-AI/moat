import React, { Component } from 'react'
import PurchaseAgreementService from '../services/PurchaseAgreementService';

class UpdatePurchaseAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                agreementNumber: '',
                effectiveDate: ''
        }
        this.updatePurchaseAgreement = this.updatePurchaseAgreement.bind(this);

        this.changeagreementNumberHandler = this.changeagreementNumberHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
    }

    componentDidMount(){
        PurchaseAgreementService.getPurchaseAgreementById(this.state.id).then( (res) =>{
            let purchaseAgreement = res.data;
            this.setState({
                agreementNumber: purchaseAgreement.agreementNumber,
                effectiveDate: purchaseAgreement.effectiveDate
            });
        });
    }

    updatePurchaseAgreement = (e) => {
        e.preventDefault();
        let purchaseAgreement = {
            purchaseAgreementId: this.state.id,
            agreementNumber: this.state.agreementNumber,
            effectiveDate: this.state.effectiveDate
        };
        console.log('purchaseAgreement => ' + JSON.stringify(purchaseAgreement));
        console.log('id => ' + JSON.stringify(this.state.id));
        PurchaseAgreementService.updatePurchaseAgreement(purchaseAgreement).then( res => {
            this.props.history.push('/purchaseAgreements');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PurchaseAgreement</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> agreementNumber: </label>
                                                <input placeholder="agreementNumber" name="agreementNumber" className="form-control" value={this.state.agreementNumber} onChange={this.changeagreementNumberHandler}/>

                                            <label> effectiveDate: </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePurchaseAgreement}>Save</button>
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

export default UpdatePurchaseAgreementComponent
