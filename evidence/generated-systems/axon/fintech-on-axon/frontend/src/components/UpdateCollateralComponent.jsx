import React, { Component } from 'react'
import CollateralService from '../services/CollateralService';

class UpdateCollateralComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                description: '',
                value: '',
                collateralType: ''
        }
        this.updateCollateral = this.updateCollateral.bind(this);

        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
        this.changeCollateralTypeHandler = this.changeCollateralTypeHandler.bind(this);
    }

    componentDidMount(){
        CollateralService.getCollateralById(this.state.id).then( (res) =>{
            let collateral = res.data;
            this.setState({
                description: collateral.description,
                value: collateral.value,
                collateralType: collateral.collateralType
            });
        });
    }

    updateCollateral = (e) => {
        e.preventDefault();
        let collateral = {
            collateralId: this.state.id,
            description: this.state.description,
            value: this.state.value,
            collateralType: this.state.collateralType
        };
        console.log('collateral => ' + JSON.stringify(collateral));
        console.log('id => ' + JSON.stringify(this.state.id));
        CollateralService.updateCollateral(collateral).then( res => {
            this.props.history.push('/collaterals');
        });
    }

    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changevalueHandler= (event) => {
        this.setState({value: event.target.value});
    }
    changeCollateralTypeHandler= (event) => {
        this.setState({collateralType: event.target.value});
    }

    cancel(){
        this.props.history.push('/collaterals');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Collateral</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> value: </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                            <label> CollateralType: </label>
                                                <select value={this.state.collateralType} onChange={this.changeCollateralTypeHandler}>
                      <option name="CollateralType" className="form-control" >
                          RealEstate
                      </option>
                      <option name="CollateralType" className="form-control" >
                          Deposit
                      </option>
                      <option name="CollateralType" className="form-control" >
                          PersonalGuarantee
                      </option>
                      <option name="CollateralType" className="form-control" >
                          Inventory
                      </option>
                      <option name="CollateralType" className="form-control" >
                          Equipment
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCollateral}>Save</button>
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

export default UpdateCollateralComponent
