import React, { Component } from 'react'
import BusinessUnitService from '../services/BusinessUnitService';

class UpdateBusinessUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                code: '',
                category: ''
        }
        this.updateBusinessUnit = this.updateBusinessUnit.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    componentDidMount(){
        BusinessUnitService.getBusinessUnitById(this.state.id).then( (res) =>{
            let businessUnit = res.data;
            this.setState({
                name: businessUnit.name,
                code: businessUnit.code,
                category: businessUnit.category
            });
        });
    }

    updateBusinessUnit = (e) => {
        e.preventDefault();
        let businessUnit = {
            businessUnitId: this.state.id,
            name: this.state.name,
            code: this.state.code,
            category: this.state.category
        };
        console.log('businessUnit => ' + JSON.stringify(businessUnit));
        console.log('id => ' + JSON.stringify(this.state.id));
        BusinessUnitService.updateBusinessUnit(businessUnit).then( res => {
            this.props.history.push('/businessUnits');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changeCategoryHandler= (event) => {
        this.setState({category: event.target.value});
    }

    cancel(){
        this.props.history.push('/businessUnits');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BusinessUnit</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> Category: </label>
                                                <select value={this.state.category} onChange={this.changeCategoryHandler}>
                      <option name="Category" className="form-control" >
                          ConsumerGoods
                      </option>
                      <option name="Category" className="form-control" >
                          IndustrialEquipment
                      </option>
                      <option name="Category" className="form-control" >
                          Electronics
                      </option>
                      <option name="Category" className="form-control" >
                          Pharmaceuticals
                      </option>
                      <option name="Category" className="form-control" >
                          FoodBeverage
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBusinessUnit}>Save</button>
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

export default UpdateBusinessUnitComponent
