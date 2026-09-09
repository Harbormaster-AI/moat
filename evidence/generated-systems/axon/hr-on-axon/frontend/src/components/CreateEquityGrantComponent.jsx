import React, { Component } from 'react'
import EquityGrantService from '../services/EquityGrantService';

class CreateEquityGrantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                grantId: '',
                grantedUnits: '',
                vestingStart: '',
                grantType: ''
        }
        this.changegrantIdHandler = this.changegrantIdHandler.bind(this);
        this.changegrantedUnitsHandler = this.changegrantedUnitsHandler.bind(this);
        this.changevestingStartHandler = this.changevestingStartHandler.bind(this);
        this.changeGrantTypeHandler = this.changeGrantTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateEquityGrant = (e) => {
        e.preventDefault();
        let equityGrant = {
                equityGrantId: this.state.id,
                grantId: this.state.grantId,
                grantedUnits: this.state.grantedUnits,
                vestingStart: this.state.vestingStart,
                grantType: this.state.grantType
            };
        console.log('equityGrant => ' + JSON.stringify(equityGrant));

        // step 5
        if(this.state.id === '_add'){
            equityGrant.equityGrantId=''
            EquityGrantService.createEquityGrant(equityGrant).then(res =>{
                this.props.history.push('/equityGrants');
            });
        }else{
            EquityGrantService.updateEquityGrant(equityGrant).then( res => {
                this.props.history.push('/equityGrants');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add EquityGrant</h3>
        }else{
            return <h3 className="text-center">Update EquityGrant</h3>
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
                                            <label> grantId:&emsp; </label>
                                                <input placeholder="grantId" name="grantId" className="form-control" value={this.state.grantId} onChange={this.changegrantIdHandler}/>

                                            <label> grantedUnits:&emsp; </label>
                                                <input type="number" placeholder="grantedUnits" name="grantedUnits" className="form-control" value={this.state.grantedUnits} onChange={this.changegrantedUnitsHandler}/>

                                            <label> vestingStart:&emsp; </label>
                                                <input type="date" placeholder="vestingStart" name="vestingStart" className="form-control" value={this.state.vestingStart} onChange={this.changevestingStartHandler}/>

                                            <label> GrantType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateEquityGrant}>Save</button>
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

export default CreateEquityGrantComponent
