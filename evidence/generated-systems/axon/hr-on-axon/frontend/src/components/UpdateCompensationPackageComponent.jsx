import React, { Component } from 'react'
import CompensationPackageService from '../services/CompensationPackageService';

class UpdateCompensationPackageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                effectiveFrom: '',
                effectiveTo: '',
                currency: ''
        }
        this.updateCompensationPackage = this.updateCompensationPackage.bind(this);

        this.changeeffectiveFromHandler = this.changeeffectiveFromHandler.bind(this);
        this.changeeffectiveToHandler = this.changeeffectiveToHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
    }

    componentDidMount(){
        CompensationPackageService.getCompensationPackageById(this.state.id).then( (res) =>{
            let compensationPackage = res.data;
            this.setState({
                effectiveFrom: compensationPackage.effectiveFrom,
                effectiveTo: compensationPackage.effectiveTo,
                currency: compensationPackage.currency
            });
        });
    }

    updateCompensationPackage = (e) => {
        e.preventDefault();
        let compensationPackage = {
            compensationPackageId: this.state.id,
            effectiveFrom: this.state.effectiveFrom,
            effectiveTo: this.state.effectiveTo,
            currency: this.state.currency
        };
        console.log('compensationPackage => ' + JSON.stringify(compensationPackage));
        console.log('id => ' + JSON.stringify(this.state.id));
        CompensationPackageService.updateCompensationPackage(compensationPackage).then( res => {
            this.props.history.push('/compensationPackages');
        });
    }

    changeeffectiveFromHandler= (event) => {
        this.setState({effectiveFrom: event.target.value});
    }
    changeeffectiveToHandler= (event) => {
        this.setState({effectiveTo: event.target.value});
    }
    changecurrencyHandler= (event) => {
        this.setState({currency: event.target.value});
    }

    cancel(){
        this.props.history.push('/compensationPackages');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CompensationPackage</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> effectiveFrom: </label>
                                                <input type="date" placeholder="effectiveFrom" name="effectiveFrom" className="form-control" value={this.state.effectiveFrom} onChange={this.changeeffectiveFromHandler}/>

                                            <label> effectiveTo: </label>
                                                <input type="date" placeholder="effectiveTo" name="effectiveTo" className="form-control" value={this.state.effectiveTo} onChange={this.changeeffectiveToHandler}/>

                                            <label> currency: </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCompensationPackage}>Save</button>
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

export default UpdateCompensationPackageComponent
