import React, { Component } from 'react'
import EquityGrantService from '../services/EquityGrantService';

class UpdateEquityGrantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                grantId: '',
                grantedUnits: '',
                vestingStart: '',
                grantType: ''
        }
        this.updateEquityGrant = this.updateEquityGrant.bind(this);

        this.changegrantIdHandler = this.changegrantIdHandler.bind(this);
        this.changegrantedUnitsHandler = this.changegrantedUnitsHandler.bind(this);
        this.changevestingStartHandler = this.changevestingStartHandler.bind(this);
        this.changeGrantTypeHandler = this.changeGrantTypeHandler.bind(this);
    }

    componentDidMount(){
        EquityGrantService.getEquityGrantById(this.state.id).then( (res) =>{
            let equityGrant = res.data;
            this.setState({
                grantId: equityGrant.grantId,
                grantedUnits: equityGrant.grantedUnits,
                vestingStart: equityGrant.vestingStart,
                grantType: equityGrant.grantType
            });
        });
    }

    updateEquityGrant = (e) => {
        e.preventDefault();
        let equityGrant = {
            equityGrantId: this.state.id,
            grantId: this.state.grantId,
            grantedUnits: this.state.grantedUnits,
            vestingStart: this.state.vestingStart,
            grantType: this.state.grantType
        };
        console.log('equityGrant => ' + JSON.stringify(equityGrant));
        console.log('id => ' + JSON.stringify(this.state.id));
        EquityGrantService.updateEquityGrant(equityGrant).then( res => {
            this.props.history.push('/equityGrants');
        });
    }

    changegrantIdHandler= (event) => {
        this.setState({grantId: event.target.value});
    }
    changegrantedUnitsHandler= (event) => {
        this.setState({grantedUnits: event.target.value});
    }
    changevestingStartHandler= (event) => {
        this.setState({vestingStart: event.target.value});
    }
    changeGrantTypeHandler= (event) => {
        this.setState({grantType: event.target.value});
    }

    cancel(){
        this.props.history.push('/equityGrants');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update EquityGrant</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> grantId: </label>
                                                <input placeholder="grantId" name="grantId" className="form-control" value={this.state.grantId} onChange={this.changegrantIdHandler}/>

                                            <label> grantedUnits: </label>
                                                <input type="number" placeholder="grantedUnits" name="grantedUnits" className="form-control" value={this.state.grantedUnits} onChange={this.changegrantedUnitsHandler}/>

                                            <label> vestingStart: </label>
                                                <input type="date" placeholder="vestingStart" name="vestingStart" className="form-control" value={this.state.vestingStart} onChange={this.changevestingStartHandler}/>

                                            <label> GrantType: </label>
                                                <select value={this.state.grantType} onChange={this.changeGrantTypeHandler}>
                      <option name="GrantType" className="form-control" >
                          RSU
                      </option>
                      <option name="GrantType" className="form-control" >
                          StockOption
                      </option>
                      <option name="GrantType" className="form-control" >
                          ESPP
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateEquityGrant}>Save</button>
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

export default UpdateEquityGrantComponent
